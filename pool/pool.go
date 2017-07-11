package pool

import (
	"github.com/mediocregopher/radix.v2/redis"
	"time"
)

// Pool is a simple connection pool for redis Clients. It will create a small
// pool of initial connections, and if more connections are needed they will be
// created on demand. If a connection is Put back and the pool is full it will
// be closed.
type Pool struct {
	pool chan *redis.Client
	df   DialFunc
	MaxSize,ActivitySize,InSize int
	// The network/address that the pool is connecting to. These are going to be
	// whatever was passed into the New function. These should not be
	// changed after the pool is initialized
	Network, Addr,Passwrod string
}

// DialFunc is a function which can be passed into NewCustom
type DialFunc func(network, addr  string) (*redis.Client, error)

// NewCustom is like New except you can specify a DialFunc which will be
// used when creating new connections for the pool. The common use-case is to do
// authentication for new connections.
func NewCustom(network, addr,password string, size int, df DialFunc) (*Pool, error) {
	var client *redis.Client
	var err error
	pool := make([]*redis.Client, 0, size)
	for i := 0; i < size; i++ {
		client, err = df(network, addr)
		if err != nil {
			for _, client = range pool {
				client.Close()
			}
			pool = pool[:0]
			break
		}
		if len(password)>0{
			_,err=client.Auth(password)
			if err != nil {
				client.Close()
				for _, client = range pool {
					client.Close()
				}
				pool = pool[:0]
				break
			}
		}
		_,err=client.Ping()
		if err!=nil {
			client.Close()
			for _, client = range pool {
				client.Close()
			}
			pool = pool[:0]
			break
		}
		pool = append(pool, client)
	}
	p := Pool{
		Network: network,
		Addr:    addr,
		Passwrod: password,
		MaxSize: size,
		ActivitySize: 0,
		InSize: size,
		pool:    make(chan *redis.Client, len(pool)),
		df:      df,
	}
	for i := range pool {
		p.pool <- pool[i]
	}
	return &p, err
}

// New creates a new Pool whose connections are all created using
// redis.Dial(network, addr). The size indicates the maximum number of idle
// connections to have waiting to be used at any given moment. If an error is
// encountered an empty (but still usable) pool is returned alongside that error
func New(network, addr,password string, size int) (*Pool, error) {
	return NewCustom(network, addr,password, size, redis.Dial)
}

// Get retrieves an available redis client. If there are none available it will
// create a new one on the fly
func (p *Pool) Get() (*redis.Client, error) {
	select {
	case conn := <-p.pool:
		p.InSize-=1
		_,err:=conn.Ping()
		if err!=nil{
			conn.Close()
			return p.Get()
		}
		p.ActivitySize+=1
		return conn, nil
	default:
		if p.ActivitySize+p.InSize<p.MaxSize{
			p.ActivitySize+=1
			client, err := p.df(p.Network, p.Addr)
			if err != nil {
				return client,err
			}
			if len(p.Passwrod)>0{
				_,err=client.Auth(p.Passwrod)
				if err != nil {
					client.Close()
					return client,err
				}
			}
			_,err=client.Ping()
			if err!=nil {
				client.Close()
				return client,err
			}
			return client,err
		}else {
			time.Sleep(time.Millisecond*50)
			return p.Get()
		}
	}
}

// Put returns a client back to the pool. If the pool is full the client is
// closed instead. If the client is already closed (due to connection failure or
// what-have-you) it will not be put back in the pool
func (p *Pool) Put(conn *redis.Client) {
	p.ActivitySize-=1
	if conn.LastCritical == nil {
		select {
		case p.pool <- conn:
			p.InSize+=1
		default:
			conn.Close()
		}
	}
}

// Cmd automatically gets one client from the pool, executes the given command
// (returning its result), and puts the client back in the pool
func (p *Pool) Cmd(cmd string, args ...interface{}) *redis.Resp {
	c, err := p.Get()
	if err != nil {
		return redis.NewResp(err)
	}
	defer p.Put(c)

	return c.Cmd(cmd, args...)
}

// Empty removes and calls Close() on all the connections currently in the pool.
// Assuming there are no other connections waiting to be Put back this method
// effectively closes and cleans up the pool.
func (p *Pool) Empty() {
	p.InSize=0
	var conn *redis.Client
	for {
		select {
		case conn = <-p.pool:
			conn.Close()
		default:
			return
		}
	}
}

// Avail returns the number of connections currently available to be gotten from
// the Pool using Get. If the number is zero then subsequent calls to Get will
// be creating new connections on the fly
func (p *Pool) Avail() int {
	return len(p.pool)
}

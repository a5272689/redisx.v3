package pool

import (
	"strconv"
	"time"
)



func (p *Pool) Lpush(key string,vals ...interface{}) (int, error) {
	newargs:=[]interface{}{key}
	for _,tmpval:=range vals{
		newargs=append(newargs,tmpval)
	}
	rs:=p.Cmd("lpush",newargs...)
	result,err:=rs.Int()
	return result,err
}

func (p *Pool) Llen(key string) (int, error) {
	rs:=p.Cmd("llen",key)
	result,err:=rs.Int()
	return result,err
}

func (p *Pool) Lindex(key string ,index int) (string, error)  {
	rs:=p.Cmd("lindex",key,index)
	result,err:=rs.Str()
	return result,err
}

func (p *Pool) Ldel(key string ,index int) (error)  {
	tmpval:=strconv.Itoa(int(time.Now().UnixNano()))
	_,err:=c.Lset(key,index,tmpval)
	if err==nil{
		_,err=c.Lrem(key,1,tmpval)
	}
	return err
}

func (p *Pool) Lset(key string ,index int,val interface{}) (string, error)  {
	rs:=p.Cmd("lset",key,index,val)
	result,err:=rs.Str()
	return result,err
}

func (p *Pool) Lrem(key string ,count int,val interface{}) (int, error)  {
	rs:=p.Cmd("lrem",key,count,val)
	result,err:=rs.Int()
	return result,err
}

func (p *Pool) Lrange(key string ,start,stop int) ([]string, error)  {
	rs:=p.Cmd("lrange",key,start,stop)
	result,err:=rs.List()
	return result,err
}

func (p *Pool) Ltrim(key string ,start,stop int) (string, error)  {
	rs:=p.Cmd("ltrim",key,start,stop)
	result,err:=rs.Str()
	return result,err
}

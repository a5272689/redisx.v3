package redis

import (
	"time"
	"strconv"
)


func (c *Client) Lpush(key string,vals ...interface{}) (int, error) {
	newargs:=[]interface{}{key}
	for _,tmpval:=range vals{
		newargs=append(newargs,tmpval)
	}
	rs:=c.Cmd("lpush",newargs...)
	result,err:=rs.Int()
	return result,err
}

func (c *Client) Llen(key string) (int, error) {
	rs:=c.Cmd("llen",key)
	result,err:=rs.Int()
	return result,err
}

func (c *Client) Lindex(key string ,index int) (string, error)  {
	rs:=c.Cmd("lindex",key,index)
	result,err:=rs.Str()
	return result,err
}

func (c *Client) Ldel(key string ,index int) (error)  {
	tmpval:=strconv.Itoa(int(time.Now().UnixNano()))
	resultLset,err:=c.Lset(key,index,tmpval)
	if resultLset=="ok"{
		_,err=c.Lrem(key,1,tmpval)
	}
	return err
}

func (c *Client) Lset(key string ,index int,val interface{}) (string, error)  {
	rs:=c.Cmd("lset",key,index,val)
	result,err:=rs.Str()
	return result,err
}

func (c *Client) Lrem(key string ,count int,val interface{}) (int, error)  {
	rs:=c.Cmd("lrem",key,count,val)
	result,err:=rs.Int()
	return result,err
}

func (c *Client) Lrange(key string ,start,stop int) ([]string, error)  {
	rs:=c.Cmd("lrange",key,start,stop)
	result,err:=rs.List()
	return result,err
}
package redis

func (c *Client) Hset(key,field string,value interface{}) (int, error) {
	rs:=c.Cmd("hset",key,field,value)
	result,err:=rs.Int()
	return result,err
}


func (c *Client) Hmset(key,field string,value interface{},args ...interface{}) (string, error) {
	newargs:=[]interface{}{key,field,value}
	for _,tmparg:=range args{
		newargs=append(newargs,tmparg)
	}
	rs:=c.Cmd("hmset",newargs...)
	result,err:=rs.Str()
	return result,err
}

func (c *Client) Hexists(key,field string) (bool, error) {
	rs:=c.Cmd("hexists",key,field)
	result,err:=rs.Int()
	if result==0{
		return false,err
	}
	return true,err
}


func (c *Client) Hmget(key,field string,fields ...interface{}) ([]string, error) {
	newargs:=[]interface{}{key,field}
	for _,tmpfield:=range fields{
		newargs=append(newargs,tmpfield)
	}
	rs:=c.Cmd("hmget",newargs...)
	result,err:=rs.List()
	return result,err
}

func (c *Client) Hget(key,field string) (string, error) {
	rs:=c.Cmd("hget",key,field)
	result,err:=rs.Str()
	return result,err
}


func (c *Client) Hdel(key string,field string,fields ...string) (int, error) {
	newargs:=[]interface{}{key,field}
	for _,tmpfield:=range fields{
		newargs=append(newargs,tmpfield)
	}
	rs:=c.Cmd("hdel",newargs...)
	result,err:=rs.Int()
	return result,err
}

func (c *Client) Hkeys(key string) ([]string, error) {
	rs:=c.Cmd("hkeys",key)
	result,err:=rs.List()
	return result,err
}

func (c *Client) Hgetall(key string) (map[string]string, error) {
	rs:=c.Cmd("hgetall",key)
	result,err:=rs.Map()
	return result,err
}

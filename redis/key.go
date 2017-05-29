package redis

func (c *Client) Keys(pattern string) ([]string, error) {
	rs:=c.Cmd("keys",pattern)
	result,err:=rs.List()
	return result,err
}

func (c *Client) Type(key string) (string, error) {
	rs:=c.Cmd("type",key)
	result,err:=rs.Str()
	return result,err
}

func (c *Client) Ttl(key string) (int, error) {
	rs:=c.Cmd("ttl",key)
	result,err:=rs.Int()
	return result,err
}

func (c *Client) Del(key string) (int, error) {
	rs:=c.Cmd("del",key)
	result,err:=rs.Int()
	return result,err
}


func (c *Client) Expire(key string,seconds int) (int, error) {
	rs:=c.Cmd("expire",key,seconds)
	result,err:=rs.Int()
	return result,err
}

func (c *Client) Persist(key string) (int, error) {
	rs:=c.Cmd("persist",key)
	result,err:=rs.Int()
	return result,err
}


func (c *Client) Renamenx(key,newkey string) (int, error) {
	rs:=c.Cmd("renamenx",key,newkey)
	result,err:=rs.Int()
	return result,err
}


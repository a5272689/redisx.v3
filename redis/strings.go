package redis

func (c *Client) Get(key string) (string, error) {
	rs:=c.Cmd("get",key)
	result,err:=rs.Str()
	return result,err
}

func (c *Client) Set(key,val interface{}) (string, error) {
	rs:=c.Cmd("set",key,val)
	result,err:=rs.Str()
	return result,err
}


package redis


func (c *Client) Sadd(key string,members ...interface{}) (int, error) {
	newargs:=[]interface{}{key}
	for _,tmpval:=range members{
		newargs=append(newargs,tmpval)
	}
	rs:=c.Cmd("sadd",newargs...)
	result,err:=rs.Int()
	return result,err
}


func (c *Client) Smembers(key string,members ...interface{}) ([]string, error) {
	newargs:=[]interface{}{key}
	for _,tmpval:=range members{
		newargs=append(newargs,tmpval)
	}
	rs:=c.Cmd("SMEMBERS",newargs...)
	result,err:=rs.List()
	return result,err
}

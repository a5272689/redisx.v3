package redis

func (c *Client) Auth(password string) (string, error) {
	rs:=c.Cmd("auth",password)
	result,err:=rs.Str()
	return result,err
}

func (c *Client) Ping() (string, error) {
	rs:=c.Cmd("ping")
	result,err:=rs.Str()
	return result,err
}

func (c *Client) Echo(message string) (string, error) {
	rs:=c.Cmd("echo",message)
	result,err:=rs.Str()
	return result,err
}

func (c *Client) Select(index int) (string, error) {
	rs:=c.Cmd("select",index)
	result,err:=rs.Str()
	return result,err
}

func (c *Client) Quit() (string, error) {
	rs:=c.Cmd("quit")
	result,err:=rs.Str()
	return result,err
}
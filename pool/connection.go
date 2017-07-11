package pool

func (p *Pool) Auth(password string) (string, error) {
	rs:=p.Cmd("auth",password)
	result,err:=rs.Str()
	return result,err
}

func (p *Pool) Ping() (string, error) {
	rs:=p.Cmd("ping")
	result,err:=rs.Str()
	return result,err
}

func (p *Pool) Echo(message string) (string, error) {
	rs:=p.Cmd("echo",message)
	result,err:=rs.Str()
	return result,err
}

func (p *Pool) Select(index int) (string, error) {
	rs:=p.Cmd("select",index)
	result,err:=rs.Str()
	return result,err
}

func (p *Pool) Quit() (string, error) {
	rs:=p.Cmd("quit")
	result,err:=rs.Str()
	return result,err
}

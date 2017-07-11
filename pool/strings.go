package pool

func (p *Pool) GetStr(key string) (string, error) {
	rs:=p.Cmd("get",key)
	result,err:=rs.Str()
	return result,err
}

func (p *Pool) Set(key,val interface{}) (string, error) {
	rs:=p.Cmd("set",key,val)
	result,err:=rs.Str()
	return result,err
}

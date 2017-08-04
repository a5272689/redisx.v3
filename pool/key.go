package pool

func (p *Pool) Keys(pattern string) ([]string, error) {
	rs:=p.Cmd("keys",pattern)
	result,err:=rs.List()
	return result,err
}

func (p *Pool) Type(key string) (string, error) {
	rs:=p.Cmd("type",key)
	result,err:=rs.Str()
	return result,err
}

func (p *Pool) Ttl(key string) (int, error) {
	rs:=p.Cmd("ttl",key)
	result,err:=rs.Int()
	return result,err
}

func (p *Pool) Del(key string) (int, error) {
	rs:=p.Cmd("del",key)
	result,err:=rs.Int()
	return result,err
}


func (p *Pool) Expire(key string,seconds int) (int, error) {
	rs:=p.Cmd("expire",key,seconds)
	result,err:=rs.Int()
	return result,err
}

func (p *Pool) Persist(key string) (int, error) {
	rs:=p.Cmd("persist",key)
	result,err:=rs.Int()
	return result,err
}


func (p *Pool) Renamenx(key,newkey string) (int, error) {
	rs:=p.Cmd("renamenx",key,newkey)
	result,err:=rs.Int()
	return result,err
}

func (p *Pool) Exists(key string) (int, error) {
	rs:=p.Cmd("exists",key)
	result,err:=rs.Int()
	return result,err
}


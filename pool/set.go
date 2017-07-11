package pool

func (p *Pool) Sadd(key string,members ...interface{}) (int, error) {
	newargs:=[]interface{}{key}
	for _,tmpval:=range members{
		newargs=append(newargs,tmpval)
	}
	rs:=p.Cmd("sadd",newargs...)
	result,err:=rs.Int()
	return result,err
}


func (p *Pool) Smembers(key string) ([]string, error) {
	rs:=p.Cmd("SMEMBERS",key)
	result,err:=rs.List()
	return result,err
}

func (p *Pool) Srem(key string,members ...interface{}) (int,error)  {
	newargs:=[]interface{}{key}
	for _,tmpval:=range members{
		newargs=append(newargs,tmpval)
	}
	rs:=p.Cmd("srem",newargs...)
	result,err:=rs.Int()
	return result,err
}

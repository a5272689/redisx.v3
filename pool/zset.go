package pool

func (p *Pool) Zadd(key string,score int,member interface{},args ...interface{}) (int, error) {
	newargs:=[]interface{}{key,score,member}
	for _,tmparg:=range args{
		newargs=append(newargs,tmparg)
	}
	rs:=p.Cmd("zadd",newargs...)
	result,err:=rs.Int()
	return result,err
}

func (p *Pool) Zrange(key string,start,stop int,withscore bool) ([]string, error) {
	newargs:=[]interface{}{key,start,stop}
	if withscore{
		newargs=append(newargs,"WITHSCORES")
	}
	rs:=p.Cmd("zrange",newargs...)
	result,err:=rs.List()
	return result,err
}

func (p *Pool) Zrem(key string,members ...interface{}) (int,error)  {
	newargs:=[]interface{}{key}
	for _,tmpval:=range members{
		newargs=append(newargs,tmpval)
	}
	rs:=p.Cmd("zrem",newargs...)
	result,err:=rs.Int()
	return result,err
}

package pool

func (p *Pool) Hset(key,field string,value interface{}) (int, error) {
	rs:=p.Cmd("hset",key,field,value)
	result,err:=rs.Int()
	return result,err
}


func (p *Pool) Hmset(key,field string,value interface{},args ...interface{}) (string, error) {
	newargs:=[]interface{}{key,field,value}
	for _,tmparg:=range args{
		newargs=append(newargs,tmparg)
	}
	rs:=p.Cmd("hmset",newargs...)
	result,err:=rs.Str()
	return result,err
}

func (p *Pool) Hexists(key,field string) (bool, error) {
	rs:=p.Cmd("hexists",key,field)
	result,err:=rs.Int()
	if result==0{
		return false,err
	}
	return true,err
}


func (p *Pool) Hmget(key,field string,fields ...interface{}) ([]string, error) {
	newargs:=[]interface{}{key,field}
	for _,tmpfield:=range fields{
		newargs=append(newargs,tmpfield)
	}
	rs:=p.Cmd("hmget",newargs...)
	result,err:=rs.List()
	return result,err
}

func (p *Pool) Hget(key,field string) (string, error) {
	rs:=p.Cmd("hget",key,field)
	result,err:=rs.Str()
	return result,err
}


func (p *Pool) Hdel(key string,field string,fields ...string) (int, error) {
	newargs:=[]interface{}{key,field}
	for _,tmpfield:=range fields{
		newargs=append(newargs,tmpfield)
	}
	rs:=p.Cmd("hdel",newargs...)
	result,err:=rs.Int()
	return result,err
}

func (p *Pool) Hkeys(key string) ([]string, error) {
	rs:=p.Cmd("hkeys",key)
	result,err:=rs.List()
	return result,err
}

func (p *Pool) Hgetall(key string) (map[string]string, error) {
	rs:=p.Cmd("hgetall",key)
	result,err:=rs.Map()
	return result,err
}


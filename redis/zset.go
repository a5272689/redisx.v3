package redis


func (c *Client) Zadd(key string,score int,member interface{},args ...interface{}) (int, error) {
	newargs:=[]interface{}{key,score,member}
	for _,tmparg:=range args{
		newargs=append(newargs,tmparg)
	}
	rs:=c.Cmd("zadd",newargs...)
	result,err:=rs.Int()
	return result,err
}

func (c *Client) Zrange(key string,start,stop int,withscore bool) ([]string, error) {
	newargs:=[]interface{}{key,start,stop}
	if withscore{
		newargs=append(newargs,"WITHSCORES")
	}
	rs:=c.Cmd("zrange",newargs...)
	result,err:=rs.List()
	return result,err
}

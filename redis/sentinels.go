package redis



func (c *Client) Masters() ([]map[string]string, error) {
	rs:=c.Cmd("SENTINEL","masters")
	result:=[]map[string]string{}
	masters,err:=rs.Array()
	if err!=nil{
		return result,err
	}
	for _,master:=range masters{
		masterinfo,_:=master.Map()
		result=append(result,masterinfo)
	}
	return result,err
}


func (c *Client) Slaves(mastername string) ([]map[string]string, error) {
	rs:=c.Cmd("SENTINEL","slaves",mastername)
	result:=[]map[string]string{}
	slaves,err:=rs.Array()
	if err!=nil{
		return result,err
	}
	for _,slave:=range slaves{
		slaveinfo,_:=slave.Map()
		result=append(result,slaveinfo)
	}
	return result,err
}

package redis

import (
	"strings"
)

func (c *Client) Role() ([]string, error) {
	rs:=c.Cmd("role")
	result,err:=rs.List()
	return result,err
}


func (c *Client) Info(sections ...string) (map[string]string, error) {
	newargs:=[]interface{}{}
	for _,tmpsection:=range sections{
		newargs=append(newargs,tmpsection)
	}
	rs:=c.Cmd("info",newargs...)
	resultmap:=make(map[string]string)
	result,err:=rs.Str()
	infos:=strings.Split(result,"\n")
	for _,infostr:=range infos{
		infolist:=strings.Split(infostr,":")
		if len(infolist)==2{
			resultmap[infolist[0]]=strings.Trim(infolist[1],"\r")
		}
	}
	return resultmap,err
}

func (c *Client) ConfigGet(parameter string) (map[string]string, error) {
	rs:=c.Cmd("config","get",parameter)
	result,err:=rs.List()
	mapresult:=make(map[string]string)
	for i,config:=range result{
		if i==0||i%2!=1{
			mapresult[config]=result[i+1]
		}
	}
	return mapresult,err
}
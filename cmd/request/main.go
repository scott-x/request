package main

import (
	"flag"
	"fmt"
	"github.com/scott-x/request"
	"log"
	"regexp"
	"strings"
)

func main()  {
	var url string
	flag.StringVar(&url,"u","http://www.baidu.com","the url is used for request")
	flag.Parse()
	re := regexp.MustCompile(`http[s]?://.*`)
	if len(re.FindString(url))>0 {
		content,err:=request.Fetch(url)
		if err!=nil{
			if strings.Contains(fmt.Sprintf("%v",err),"net/http: TLS handshake timeout"){
				log.Println("该网站不支持golang爬虫，请放弃治疗吧")
				return
			}
		}
		fmt.Println(string(content))
	}
}
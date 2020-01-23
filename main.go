package main

import (
	"fmt"
	"net/http"
	"webStressTesting/route"
)

func main(){
	fmt.Println("浏览器打开网址:http://localhost")
	http.HandleFunc("/",route.Index)
	http.HandleFunc("/work",route.Work)
	//exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost").Start()
	_ = http.ListenAndServe("localhost:80", nil)


}

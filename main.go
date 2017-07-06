package main

import (
	"fmt"
	"strings"
	"io"
	"log"
	"net/http"
	//_ "bingyan/routers"
	//"github.com/astaxie/beego"
)

type data struct()
	headline string
	content string
	like bool
	
}

func init(){
	//导入mysq数据库用，不过我现在不会、、
}

func (ex data)like() /*接收者*/{
	//点赞
}

func (ex data)comment(){
	//评论
}////////////////////////////////////要用chinnel吗？？




func pushup(this *data)
func main() {
	http.HandleFunc("/",sayhello)

	err := http.ListenAndServe(":9090",nil)
	if err !=nil{
		log.Fatal("htp.ListenAndServe:",err)
	}
	//beego.Run()
}

func sayhello(w http.ResponseWriter , r *http.Request){
	io.WriteString(w, "hello world")
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["uel_long"])
	for k, v := range r.Form{
		fmt.Println("keys:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w, "hello world!")
}

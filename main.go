package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)
func GetImage(w http.ResponseWriter, r *http.Request) {

	//开启跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	// 解析url传递的参数
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		// join() 方法用于把数组中的所有元素放入一个字符串。
		// 元素是通过指定的分隔符进行分隔的
		fmt.Println("val:", strings.Join(v, ""))
	}
	// 输出到客户端
	imageData :=r.Form["imageData"]
	ImageUrl :=r.Form["ImageUrl"]
	fmt.Println("imageData = %s",imageData)
	fmt.Println("ImageUrl = %s",ImageUrl)
}

func ShoeImage(w http.ResponseWriter, r *http.Request) {
	//开启跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		// func (t *Template) Execute(wr io.Writer, data interface{}) error {
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		imageData :=r.Form["imageData"]
		ImageUrl :=r.Form["ImageUrl"]
		fmt.Println("imageData = %s",imageData)
		fmt.Println("ImageUrl = %s",ImageUrl)
	}
}

func main() {

	http.HandleFunc("/", GetImage)
	http.HandleFunc("/login", ShoeImage)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndserve:", err)
	}
}

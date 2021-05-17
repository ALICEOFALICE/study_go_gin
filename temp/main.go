package main

import (
	"fmt"
	"net/http"

	"html/template"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"Name":   "小王子",
		"Age":    18,
		"Gender": "男",
	}
	t, err := template.ParseFiles("./html.tmpl")
	if err != nil {
		fmt.Printf("报错")
	}
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
	if err != nil {
		fmt.Printf("模板渲染失败")
	}

}
func main() {
	http.HandleFunc("/", sayhello)
	fmt.Printf("已经运行于9000端口，注意安全 请注意安全性\n")
	fmt.Printf("http://127.0.0.1:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("错误")
	}

}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"io"
	"os"
	"crypto/md5"
	"strconv"
)

var counter int = 0
func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload",upload)
	http.HandleFunc("/count",Counter)
	err := http.ListenAndServe(":8081", nil) //监听端口，并且接收请求
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func sayHello(p http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认不会解析的
	fmt.Printf("%[1]T,%[1]v\n", r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val:", v)
	}
	fmt.Fprintf(p, "Hello") //写入客户端
}

//登录
func login(p http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认不会解析的
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(p, nil))
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("pwd:", r.Form["password"])
		fmt.Fprintf(p, "hello "+r.FormValue("username"))
	}
}

//文件上传
func upload(p http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(currentTime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(p,token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file,handler,err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return 
		}
		defer file.Close()
		fmt.Fprintf(p,"%v",handler.Header)
		f,err := os.OpenFile("./test/"+handler.Filename,os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
	}
}

func Counter(p http.ResponseWriter,r *http.Request) {
	counter++
	fmt.Fprintf(p,strconv.Itoa(counter))
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type PlayerModule interface {
	DecreaseCoins(num int)
}

func DecreaseCoins(num int) {
	fmt.Println("I eat bone")
}

//定义一个接口，接口内有eat和run两个方法
type Animal interface {
	Eat()
	Run()
}

//定义一个dog实体类，实现Animal接口
type Dog struct {
}

func (dog *Dog) Eat() {
	fmt.Println("I eat bone")
}
func (dog *Dog) Run() {
	fmt.Println("I run very fast")
}

func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, getHelloMsg())

}

func getHelloMsg() string {
	log.Println("getHelloMsg ", "aa")
	return "Let's Shopping!"

}

func testGetHelloMsg(t *testing.T) string {
	msg := getHelloMsg()
	if msg != "Let's Shopping!" {
		t.Error("Get hello message error")
	}
	return "OK"
}

func main() {
	//	var (
	//		Player PlayerModule
	//	)
	var a Animal
	a = &Dog{}
	a.Eat()
	a.Run()
	fmt.Println("hello world!!!")

	http.HandleFunc("/", Hello) //注册URI路径与相应的处理函数
	fmt.Println("Start listening...")

	err := http.ListenAndServe(":4040", nil) // 监听9090端口，就跟javaweb中tomcat用的8080差不多一个意思吧
	if err != nil {
		panic(err)
		log.Fatal("ListenAndServe: ", err)
	}

	// http://localhost:4040/

}

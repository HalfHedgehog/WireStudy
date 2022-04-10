/**
    @Author: qiyou_wu
    @CreateDate: 2022/4/10
    @Description:
**/
package main

import "fmt"

type Message struct {
	msg string
}

type Message1 struct {
	msg int
}

type Greeter struct {
	Message Message
}
type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// NewMessage1 Message的构造函数
func NewMessage1(msg int) Message1 {
	return Message1{
		msg: msg,
	}
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
func (g Greeter) Greet() Message {
	return g.Message
}

// 使用wire前
//func main() {
//	message := NewMessage("hello world")
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	event.Start()
//}

// 使用wire后
func main() {
	event := InitializeEvent("hello_world")

	event.Start()
}

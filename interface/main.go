package main

import "fmt"

// Phone 接口，有 1 个 call 方法
type Phone interface {
	call()
}

type NokiaPhone struct {}

// call 方法实现 Phone 接口
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I'm Nokia, i can call you!")
}

type ApplePhone struct {}

// call 方法实现 Phone 接口
func (applePhone ApplePhone) call() {
	fmt.Println("I'm Apple, i can call you!")
}

func main() {
	// Phone 接口变量
	var phone Phone
	// new() 和 NokiaPhone{} 都可以初始化对象，然后 Phone 指向具体的对象
	// phone = new(NokiaPhone)
	phone = NokiaPhone{}
	// 调用方法
	phone.call()

	phone = new(ApplePhone)
	// phone = ApplePhone{}
	phone.call()
}

// 上述体现了 interface 接口的语法，在main函数中，也体现了多态的特性。
// 同样一个 phone 抽象接口，分别指向不同的实体对象，调用 call() 方法，打印的效果不同，体现了多态的特性。
package main

import (
	"fmt"
)

// 面向抽象层的依赖倒转

/*
在设计一个系统的时候，将模块分为3个层次，抽象层、实现层、业务逻辑层。
  - 首先将抽象层的模块和接口定义出来，这里就需要了interface接口的设计
  - 然后我们依照抽象层，依次实现每个实现层的模块，在我们写实现层代码的时候，只需要参考对应的抽象层实现就好了，
    实现每个模块，也和其他的实现的模块没有关系，这样也符合了上面介绍的开闭原则。这样每个模块只依赖对象的接口，
    而和其他模块没关系，依赖关系单一。系统容易扩展和维护。
  - 业务逻辑只需要参考抽象层的接口，抽象层暴露出来的接口就是我们业务层可以使用的方法，然后可以通过多态，
    接口指针指向哪个实现模块，调用了就是具体的实现方法，这样我们业务逻辑层也是依赖抽象层编程。

上述这种设计原则叫做依赖倒转原则。业务逻辑层、实现层都依赖抽象层
*/

// 1、抽象层

// driver 接口
type driver interface {
	drive(car Car)  // drive 方法
}

// Car 接口
type Car interface {
	run()  // run 方法
}

// 2、实现层

// 奔驰
type Benz struct {}

func (benz *Benz) run() {
	fmt.Println("Benz is running ...!")
}

// 宝马
type BMW struct {}

func (bmw *BMW) run() {
	fmt.Println("BMW is running ...!")
}

// zhangsan
type ZhangSan struct {}

func (zs *ZhangSan) drive(car Car) {
	fmt.Println("ZhangSan drive car!")
	car.run()
}

// lisi
type LiSi struct {}

func (ls *LiSi) drive(car Car) {
	fmt.Println("LiSi drive car!")
	car.run()
}

// 3、业务逻辑层
func main() {
	// 1、张三开宝马
	var bmw Car
	bmw = new(BMW)

	var zs ZhangSan
	zs.drive(bmw)

	// 2、李四开奔驰
	var benz Car
	benz = new(Benz)

	var ls LiSi
	ls.drive(benz)
}


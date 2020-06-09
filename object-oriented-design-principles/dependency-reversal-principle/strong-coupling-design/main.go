package main

import "fmt"

// 强耦合设计，没有用到 interface

// 奔驰汽车
type Benz struct {}

func (benz *Benz) Run() {
	fmt.Println("Benz is running...")
}

// 宝马汽车
type BMW struct {}

func (bmw *BMW) Run() {
	fmt.Println("BMW is running ...")
}

// 司机张三
type ZhangSan struct {}

func (zs *ZhangSan) DriveBenZ(benz *Benz) {
	fmt.Println("zhangsan Drive Benz")
	benz.Run()
}

func (zs *ZhangSan) DriveBMW(bmw *BMW) {
	fmt.Println("zhangsan drive BMW")
	bmw.Run()
}

// 司机李四
type LiSi struct {}

func (ls *LiSi) DriveBenZ(benz *Benz) {
	fmt.Println("lisi Drive Benz")
	benz.Run()
}

func (ls *LiSi) DriveBMW(bmw *BMW) {
	fmt.Println("lisi drive BMW")
	bmw.Run()
}

func main() {
	// 业务1 张三开奔驰
	benz := &Benz{}
	zs := &ZhangSan{}
	zs.DriveBenZ(benz)

	// 业务2 李四开宝马
	bmw := &BMW{}
	ls := &LiSi{}
	ls.DriveBMW(bmw)
}
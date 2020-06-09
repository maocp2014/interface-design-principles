package main

import "fmt"

// 平铺式的模块设计

/*
银行业务员有很多的业务，比如Save()存款、Transfer()转账、Pay()支付等。如果这个业务员模块只有这几个方法还好，但是随着
业务的增加，业务员模块可能就要增加方法，导致业务员模块越来越臃肿。

存在的问题：
这样的设计会导致，当我们去给Banker添加新的业务的时候，会直接修改原有的Banker代码，那么Banker模块的功能会越来越多，
出现问题的几率也就越来越大，假如此时Banker已经有99个业务了，现在我们要添加第100个业务，可能由于一次的不小心，
导致之前99个业务也一起崩溃，因为所有的业务都在一个Banker类里，他们的耦合度太高，Banker的职责也不够单一，
代码的维护成本随着业务的复杂正比成倍增大。
*/

type Banker struct {}

func (b *Banker) Save() {
	fmt.Println("存款业务！")
}

func (b *Banker) Transfer() {
	fmt.Println("转账业务！")
}

func (b *Banker) Pay() {
	fmt.Println("支付业务！")
}

func main() {
	// banker := Banker{}
    banker := new(Banker)
    banker.Save()
    banker.Transfer()
    banker.Pay()
}
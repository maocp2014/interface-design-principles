package main

import "fmt"

//  开闭原则设计

// 抽象的业务员接口
type Banker interface {
	DoBusi()   // 处理业务的方法
}

// 存款业务员
type SaveBanker struct {}

// 实现 Banker 接口
func (save *SaveBanker) DoBusi() {
	fmt.Println("存款业务！")
}

// 转账业务员
type TransferBanker struct {}

// 实现 Banker 接口
func (transfer *TransferBanker) DoBusi() {
	fmt.Println("转账业务！")
}

// 支付业务员
type PayBanker struct {}

// 实现 Banker 接口
func (pay *PayBanker) DoBusi() {
	fmt.Println("支付业务！")
}

// 实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
// 函数形参是接口
func BankerBusiness(banker Banker) {
	// 通过接口来向下调用(多态现象)
	banker.DoBusi()
}

func main() {
	// var banker Banker
	// 存款
	// banker = new(SaveBanker)  // new 会返回指针
	// banker.DoBusi()

	// 转账
	// banker = new(TransferBanker)
	// banker.DoBusi()

	// 支付
	// banker = new(PayBanker)
	// banker.DoBusi()

    BankerBusiness(new(SaveBanker))
    BankerBusiness(new(TransferBanker))
    BankerBusiness(new(PayBanker))
}

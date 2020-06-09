package open_close_principle

import "fmt"

// 开闭原则设计

// 开闭原则: 一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。
// 简单的说就是在修改需求的时候，应该尽量通过扩展来实现变化，而不是通过修改已有代码来实现变化。

/*
利用 interface 进行抽象，抽象 1 个 Banker 模块，提供一个抽象的方法。根据这个抽象模块，分别去实现支付
Banker（实现支付方法）, 转账Banker（实现转账方法）


当想要给 Banker 添加额外功能时，平铺式模块设计是直接修改 Banker 内容，现在可以单独定义一个股票
Banker(实现股票方法)到这个系统中。而且股票Banker的实现成功或者失败都不会影响之前的稳定系统，他很单一，而且独立。

当我们给一个系统添加一个功能的时候，不是通过修改代码，而是通过增添代码来完成，这就是开闭原则的核心思想了。
所以要想满足上面的要求，是一定需要 interface 来提供一层抽象的接口的。
*/

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

func main() {
	var banker Banker
	// 存款
	banker = new(SaveBanker)  // new 会返回指针
	banker.DoBusi()

	// 转账
	banker = new(TransferBanker)
	banker.DoBusi()

	// 支付
	banker = new(PayBanker)
	banker.DoBusi()
}

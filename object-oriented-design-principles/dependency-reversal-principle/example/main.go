package main

import "fmt"

// 使用Golang的interface接口设计原则
// https://www.jianshu.com/p/c03d2a807e94

// 1、抽象层

// 显卡
type Card interface {
	display()
}

// 内存
type Memory interface {
	storage()
}

// cpu
type CPU interface {
	calculate()
}

// computer
type Computer struct {
	cpu CPU
	card Card
	memory Memory
}

// 构造computer的方法
func NewComputer(cpu CPU, card Card, memory Memory) *Computer {
	return &Computer{
		cpu:    cpu,
		card:   card,
		memory: memory,
	}
}

func (c *Computer) DoWork() {
	c.memory.storage()
	c.card.display()
	c.cpu.calculate()
}

// 2、实现层

// Intel 公司
type IntelCPU struct {
	CPU
}

func (intelCPU *IntelCPU) calculate() {
	fmt.Println("Intel CPU 开始计算了...")
}

type IntelMemory struct {
	Memory
}

func (intelMemory *IntelMemory) storage() {
	fmt.Println("Intel Memory 开始存储了...")
}

type IntelCard struct {
	Card
}

func (intelCard *IntelCard) display() {
	fmt.Println("Intel Card 开始显示了...")
}

// kingston
type KingstonMemory struct {
	Memory
}

func (kinMemory *KingstonMemory) storage() {
	fmt.Println("Kingston memory storage...")
}

// nvidia
type NvidiaCard struct {
	Card
}

func (nvidiaCard *NvidiaCard) display() {
	fmt.Println("Nvidia card display...")
}

// 3、业务逻辑层
func main() {
	// intel系列的电脑
	com1 := NewComputer(&IntelCPU{}, &IntelCard{}, &IntelMemory{})
	com1.DoWork()

	// 杂牌子
	com2 := NewComputer(&IntelCPU{}, &NvidiaCard{}, &KingstonMemory{})
	com2.DoWork()
}
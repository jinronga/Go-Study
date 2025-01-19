package adapter

import "fmt"

type (
	PhoneCharger interface {
		Output5V()
	}

	HuaWeiCharger struct {
	}

	XiaoMiCharger struct {
	}

	MacBookCharger struct {
	}

	MacBookChargerAdapter struct {
		core *MacBookCharger
	}
	Phone interface {
		Charge(phoneCharger PhoneCharger)
	}

	HuaWeiPhone struct {
	}
)

func NewHuaWeiCharger() *HuaWeiCharger {
	return &HuaWeiCharger{}
}

func NewXiaoMiCharger() *XiaoMiCharger {
	return &XiaoMiCharger{}
}
func NewMacBookCharger() *MacBookCharger {
	return &MacBookCharger{}
}

func NewMacBookChargerAdapter(m *MacBookCharger) *MacBookChargerAdapter {
	return &MacBookChargerAdapter{
		core: m,
	}
}

func NewHuaWeiPhone() Phone {
	return &HuaWeiPhone{}
}

func (m *MacBookChargerAdapter) Output5V() {
	m.core.Output28V()
	fmt.Println("适配器将输出电压调整为 5V...")
}

func (m *MacBookCharger) Output28V() {
	fmt.Println("苹果笔记本充电器输出 28V 电压...")
}

func (h *HuaWeiCharger) Output5V() {
	fmt.Println("华为手机充电器输出 5V 电压...")
}
func (x *XiaoMiCharger) Output5V() {
	fmt.Println("小米手机充电器输出 5V 电压...")
}

func (h *HuaWeiPhone) Charge(phoneCharger PhoneCharger) {
	fmt.Println("华为手机准备开始充电...")
	phoneCharger.Output5V()
}

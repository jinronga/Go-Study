package responsibility

import (
	"fmt"
	"testing"
)

func Test_ChainOfResponsibility(t *testing.T) {
	// 创建请求对象
	request := &PurchaseRequest{amount: 6000}

	// 创建责任链
	manager := &Manager{}
	director := &Director{}
	ceo := &CEO{}

	// 设置责任链
	manager.setNext(director)
	director.setNext(ceo)

	// 处理请求
	if !manager.approve(request) {
		fmt.Println("Request was not approved.")
	}
}

package responsibility

import "fmt"

type (

	// Approver 处理者接口
	Approver interface {
		approve(request *PurchaseRequest) bool
		setNext(approver Approver)
	}

	// Manager 具体处理者：经理
	Manager struct {
		next Approver
	}
	PurchaseRequest struct {
		amount int
	}

	// Director 具体处理者：总监
	Director struct {
		next Approver
	}
	// CEO 具体处理者：CEO
	CEO struct {
		next Approver
	}
)

func (m *Manager) approve(request *PurchaseRequest) bool {
	if request.amount <= 1000 {
		fmt.Println("Manager approved the request.")
		return true
	}
	if m.next != nil {
		return m.next.approve(request)
	}
	return false
}

func (m *Manager) setNext(approver Approver) {
	m.next = approver
}
func (d *Director) approve(request *PurchaseRequest) bool {
	if request.amount <= 5000 {
		fmt.Println("Director approved the request.")
		return true
	}
	if d.next != nil {
		return d.next.approve(request)
	}
	return false
}

func (d *Director) setNext(approver Approver) {
	d.next = approver
}

func (c *CEO) approve(request *PurchaseRequest) bool {
	if request.amount <= 10000 {
		fmt.Println("CEO approved the request.")
		return true
	}
	if c.next != nil {
		return c.next.approve(request)
	}
	return false
}

func (c *CEO) setNext(approver Approver) {
	c.next = approver
}

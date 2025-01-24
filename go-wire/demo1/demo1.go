package demo1

import (
	"fmt"
)

type (
	Dao struct {
	}
	Service struct {
		dao *Dao
	}
)

func NewDao() *Dao {
	return &Dao{}
}

func (d *Dao) Do() {
	fmt.Println("do dao")
}

func NewService(dao *Dao) *Service {
	return &Service{
		dao: dao,
	}
}

type Controller struct {
	service *Service
}

func NewController(s *Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) Run() {
	c.service.Execute()
}
func (s *Service) Execute() {
	fmt.Println("Executing service...")
	s.dao.Do()
}

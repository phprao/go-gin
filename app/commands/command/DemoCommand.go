package command

import (
	"fmt"
)

var DemoCommandObj DemoCommand

type DemoCommand struct {
	Signature   string
	Description string
}

func init() {
	DemoCommandObj = DemoCommand{
		"DemoCommand",
		"测试命令",
	}
}

func (DemoCommand) Handle() {
	fmt.Println("DemoCommand")
}

func (DemoCommand) GetSignature() string {
	return DemoCommandObj.Signature
}

func (DemoCommand) GetDescription() string {
	return DemoCommandObj.Description
}

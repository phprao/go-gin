package command

import (
	"fmt"

	"votePlatfom/app/commands/interfaces"
)

var TestCommandObj TestCommand

type TestCommand struct {
	interfaces.CommandInfo
}

func init() {
	TestCommandObj = TestCommand{}
	TestCommandObj.Signature = "TestCommand"
	TestCommandObj.Description = "测试命令"
}

func (TestCommand) Handle() {
	fmt.Println("TestCommand 测试命令")
}

func (TestCommand) GetSignature() string {
	return TestCommandObj.Signature
}

func (TestCommand) GetDescription() string {
	return TestCommandObj.Description
}

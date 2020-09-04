package interfaces

import "fmt"

var ListCommandObj ListCommand

type ListCommand struct {
	CommandInfo
}

func init() {
	ListCommandObj = ListCommand{}
	ListCommandObj.Signature = "list"
	ListCommandObj.Description = "输出命令列表"
}

func (ListCommand) GetSignature() string {
	return ListCommandObj.Signature
}

func (ListCommand) GetDescription() string {
	return ListCommandObj.Description
}

func (ListCommand) Handle() {
	fmt.Println()
	for k, v := range CommandMap {
		fmt.Printf("%s\t\t\t%s\n", k, v.GetDescription())
	}
	fmt.Println()
}


package commands

import (
	"fmt"

	"votePlatfom/app/commands/command"
	"votePlatfom/app/commands/interfaces"
)

func init() {
	interfaces.Commands = append(interfaces.Commands,
		command.TestCommandObj,
		command.DemoCommandObj,
	)

	register(interfaces.Commands)
}

func Handle(cmdStr string) {
	cmd, ok := interfaces.CommandMap[cmdStr]
	if !ok {
		panic(fmt.Sprintf("command %s was not registered\n", cmdStr))
	}

	cmd.Handle()
}

func register(C []interfaces.CommandHandle)  {
	for _, v := range C {
		interfaces.CommandMap[v.GetSignature()] = v
	}
}

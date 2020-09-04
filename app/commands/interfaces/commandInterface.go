package interfaces

var Commands []CommandHandle
var CommandMap map[string]CommandHandle

type CommandHandle interface {
	Handle()
	GetSignature() string
	GetDescription() string
}

type CommandInfo struct {
	Signature   string
	Description string
}

func init()  {
	CommandMap = make(map[string]CommandHandle)
	Commands = append(Commands, ListCommandObj)
}
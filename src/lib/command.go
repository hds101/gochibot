package lib

type Command func(Context)

type CommandStruct struct {
	command Command
	help    string
}

type CommandHandler struct {
	commands map[string]CommandStruct
}

func InitCommandHandler() *CommandHandler {
	return &CommandHandler{make(map[string]CommandStruct)}
}

func (handler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := handler.commands[name]
	return &cmd.command, found
}

func (handler CommandHandler) Register(name string, command Command, help string) {
	newCommand := CommandStruct{command: command, help: help}
	handler.commands[name] = newCommand
}

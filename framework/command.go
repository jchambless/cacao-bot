package framework

type (
	Command func(Context)

	CommandStruct struct {
		command     Command
		help        string
		description string
	}

	CmdMap map[string]CommandStruct

	CommandHandler struct {
		cmds CmdMap
	}
)

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

func (handler CommandHandler) GetCmds() CmdMap {
	return handler.cmds
}

func (handler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := handler.cmds[name]
	return &cmd.command, found
}

func (command CommandStruct) GetHelp() string {
	return command.help
}

func (command CommandStruct) GetDescription() string {
	return command.description
}

func (handler CommandHandler) Register(name string, command Command, helpmsg string, description string) {
	cmdstruct := CommandStruct{command: command, help: helpmsg, description: description}
	handler.cmds[name] = cmdstruct
	if len(name) > 1 {
		handler.cmds[name[:1]] = cmdstruct
	}
}

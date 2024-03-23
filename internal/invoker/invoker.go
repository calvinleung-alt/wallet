package invoker

import "errors"

type Command interface {
	Execute(params interface{}) (result interface{}, err error)
}

type Invoker struct {
	cmdMapping map[string]Command
}

func NewInvoker(cmdMapping map[string]Command) *Invoker {
	return &Invoker{
		cmdMapping: cmdMapping,
	}
}

func (i *Invoker) Get(opID string) (Command, error) {
	if cmd, ok := i.cmdMapping[opID]; ok {
		return cmd, nil
	}
	return nil, errors.New("command not found")
}

package cmd

import (
	"fmt"
)

type Command struct {
	options *Options
	class   string
	args    []string
}

func (cd *Command) Class() string {
	return cd.class
}

func (cd *Command) Args() []string {
	return cd.args
}

func (cd *Command) Options() *Options {
	return cd.options
}

func ParseCommand(osArgs []string) (cmd *Command, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	argReader := &ArgReader{osArgs[1:]}
	cmd = &Command{
		options: parseOptions(argReader),
		class:   argReader.removeFirse(),
		args:    argReader.args,
	}

	return
}

func PrintUsage() {
	fmt.Println("usage: goVM [-option] class [args...]")
}

package foo

import (
	"github.com/dev-vinicius-andrade/go-cobra-example/cmd/foo/bar"
	"github.com/dev-vinicius-andrade/go-cobra-example/types"
	"github.com/spf13/cobra"
)

type commandDefinition struct {
	Context       *types.CliContext
	ParentCommand *cobra.Command
}

var context = types.CliContext{}

func CreateCommand() *cobra.Command {
	c := commandDefinition{}
	return c.CreateCommandDefinition()
}
func (c *commandDefinition) CreateCommandDefinition() *cobra.Command {
	command := cobra.Command{
		Use:              "foo",
		Short:            "Foo is the main command of this tool",
		Long:             `This is a long description of Foo is the main command of this tool`,
		Aliases:          []string{},
		TraverseChildren: true, // This is important to always run the Run function of the parent command, good for global flags, context, etc
	}
	defineFlags(&command)
	defineSubCommands(&command)
	setRunCommand(&command)

	return &command
}

func defineSubCommands(cmd *cobra.Command) {
	bar.CreateCommand(&context, cmd)
}

func runCommand(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		cmd.Help()
		return
	}
}

func setRunCommand(command *cobra.Command) {
	command.Run = runCommand
}
func defineFlags(command *cobra.Command) {
	command.PersistentFlags().StringVarP(&context.GlobalMessage, "message", "m", "Default Global Message", "Global message to be used in all commands")
}

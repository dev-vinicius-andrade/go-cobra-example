package bar

import (
	"fmt"

	"github.com/dev-vinicius-andrade/go-cobra-example/helpers"
	"github.com/dev-vinicius-andrade/go-cobra-example/types"
	"github.com/spf13/cobra"
)

type commandDefinition struct {
	Context       *types.CliContext
	ParentCommand *cobra.Command
}

func CreateCommand(context *types.CliContext, parentCommand *cobra.Command) (*cobra.Command, error) {
	command := commandDefinition{
		Context:       context,
		ParentCommand: parentCommand,
	}
	return command.createCommandDefinition(context, parentCommand), nil
}
func (c *commandDefinition) createCommandDefinition(context *types.CliContext, parentCommand *cobra.Command) *cobra.Command {
	command := cobra.Command{
		Use:              "bar",
		Short:            "Bar is subcommand of foo",
		Long:             `This is a long description of Bar is a subcommand of foo`,
		Aliases:          []string{},
		TraverseChildren: true,
	}
	defineFlags(&command)
	defineSubCommands(&command)
	setRunCommand(context, &command)
	helpers.CobraHelper.AddCommandToParent(&command, parentCommand)
	return &command
}

func defineSubCommands(cmd *cobra.Command) {
	//define your subcommands here
}

func runCommand(context *types.CliContext, cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", context.GlobalMessage)
}

func setRunCommand(context *types.CliContext, command *cobra.Command) {
	command.Run = func(cmd *cobra.Command, args []string) {
		runCommand(context, cmd, args)
	}
}
func defineFlags(command *cobra.Command) {
	//define your flags of your subcommand here
}

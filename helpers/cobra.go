package helpers

import "github.com/spf13/cobra"

type cobraHelper struct{}

func (c *cobraHelper) AddCommandToParent(command *cobra.Command, parentCommand *cobra.Command) {
	if parentCommand == nil {
		return
	} else {
		parentCommand.AddCommand(command)
	}
}

var CobraHelper = cobraHelper{}

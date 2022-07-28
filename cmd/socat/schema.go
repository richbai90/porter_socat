package main

import (
	"github.com/getporter/socat/pkg/socat"
	"github.com/spf13/cobra"
)

func buildSchemaCommand(m *socat.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Print the json schema for the mixin",
		Run: func(cmd *cobra.Command, args []string) {
			m.PrintSchema()
		},
	}
	return cmd
}

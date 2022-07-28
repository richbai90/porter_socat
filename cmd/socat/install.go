package main

import (
	"os"

	"github.com/getporter/socat/pkg/socat"
	"github.com/spf13/cobra"
)

var (
	commandFile string
)

func buildInstallCommand(m *socat.Mixin) *cobra.Command {
	r, w, _ := os.Pipe()
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Execute the install functionality of this mixin",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			//Do something here if needed
			if _, dbg := os.LookupEnv("debugger"); dbg {
				w.WriteString("install:\n  - socat:\n      name: setup socat\n      description: Create a socat socket that can communicate cross platform. Required for macos to communicate via sockets.\n      leftAddress: UNIX-LISTEN:/,fork\n      options:\n        - -d\n        - -d\n      rightAddress: UNIX-CONNECT:/libvirt/\n")
				m.Context.In = r
			}
			defer w.Close()
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			defer r.Close()
			return m.Execute()
		},
	}
	return cmd
}

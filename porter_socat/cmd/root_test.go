/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestRunE(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Extra Args Test",
			args: args {
				cmd: rootCmd,
				args: []string {
					"UNIX-LISTEN:/tmp/test", "-", "--", "-d", "-d",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunE(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("RunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

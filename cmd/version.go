package main

import (
	"encoding/json"
	"fmt"
	versionpkg "github.com/nduyphuong/gorya/internal/version"
	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use: "version",
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := json.Marshal(versionpkg.GetVersion())
			if err != nil {
				return err
			}
			fmt.Println(string(data))
			return nil
		},
	}
}

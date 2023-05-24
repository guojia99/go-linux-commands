/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/5/24 下午9:26.
 * Author:  guojia(https://github.com/guojia99)
 */

package cmd

import (
	"github.com/guojia99/go-linux-commands/v1/ls"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "gl",
		Short: "Linux command line tool library implemented by golang",
		Long:  "Linux command line tool library implemented by golang",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	root.AddCommand(ls.NewLsCmd())
	return root
}

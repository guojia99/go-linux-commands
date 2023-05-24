/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/5/24 下午9:57.
 * Author:  guojia(https://github.com/guojia99)
 */

package ls

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewLsCmd() *cobra.Command {
	opt := Option{}
	cmd := &cobra.Command{
		Use:               "ls",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%+v", opt)
		},
	}
	opt.SetFlags(cmd, map[string]string{})
	return cmd
}

/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/5/24 下午9:25.
 * Author:  guojia(https://github.com/guojia99)
 */

package utils

import (
	"github.com/spf13/cobra"
)

func InitHelpFlag(cmd *cobra.Command) {
	if cmd.Flags().Lookup("help") == nil {
		usage := "help for "
		if cmd.Name() == "" {
			usage += "this command"
		} else {
			usage += cmd.Name()
		}
		cmd.Flags().Bool("help", false, usage)
		_ = cmd.Flags().SetAnnotation("help", cobra.FlagSetByCobraAnnotation, []string{"true"})
	}
}

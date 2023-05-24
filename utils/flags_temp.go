/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/5/25 上午12:07.
 * Author:  guojia(https://github.com/guojia99)
 */

package utils

import (
	"fmt"
	"github.com/spf13/cobra"
)

func SetCmdTemp(cmd *cobra.Command) {
	cmd.SetUsageFunc(func(command *cobra.Command) error {
		fmt.Println("-------------------")
		// todo 这里输出是是help格式
		fmt.Println(command.LocalFlags().FlagUsages())
		fmt.Println("-------------------")
		return nil
	})
}

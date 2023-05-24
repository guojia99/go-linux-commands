/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/5/24 下午11:30.
 * Author:  guojia(https://github.com/guojia99)
 */

package ls

import (
	"fmt"
	"github.com/guojia99/go-linux-commands/utils"
	"testing"
)

func TestOption_SetFlags(t *testing.T) {
	fmt.Printf(utils.GetFlagsPrinter(Option{}))
}

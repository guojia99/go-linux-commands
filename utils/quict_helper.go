/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/5/26 下午6:50.
 *  * Author: guojia(https://github.com/guojia99)
 */

package utils

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	helperMapTemp = `
var ZN = map[string]string{
	%s
}
`
)

// GetFlagsPrinter 用于快速输出一个结构体的flags绑定的代码段, 请勿直接用于代码中, 仅仅为了方便开发
// GetFlagsPrinter is used to quickly generate a code snippet that binds the flags of a struct for easy output.
// Please do not use it directly in your code. It is solely meant for developer convenience.
func GetFlagsPrinter(in interface{}) string {
	const (
		flagTagLong  = "long"
		flagTagShort = "short"
	)

	val, typ := reflect.ValueOf(in), reflect.TypeOf(in)
	out := "\tf := cmd.Flags()\n"

	helper := ""

	for i := 0; i < typ.NumField(); i++ {
		filedVal := val.Field(i)
		filedTyp := typ.Field(i)
		tag := filedTyp.Tag

		longTag, shortTag := tag.Get(flagTagLong), tag.Get(flagTagShort)

		if longTag == "" && shortTag == "" {
			continue
		}

		helper += fmt.Sprintf("\"%s\": ``,\n", filedTyp.Name)

		switch {
		case longTag != "" && shortTag != "":
			switch filedVal.Kind() {
			case reflect.Bool:
				out += fmt.Sprintf("\tf.BoolVarP(&o.%s, \"%s\", \"%s\", false,usageDict[\"%s\"] )\n", filedTyp.Name, longTag, shortTag, filedTyp.Name)
			case reflect.String:
				out += fmt.Sprintf("\tf.StringVarP(&o.%s, \"%s\", \"%s\", \"\",usageDict[\"%s\"] )\n", filedTyp.Name, longTag, shortTag, filedTyp.Name)
			case reflect.Int:
				out += fmt.Sprintf("\tf.IntVarP(&o.%s, \"%s\", \"%s\", 0,usageDict[\"%s\"] )\n", filedTyp.Name, longTag, shortTag, filedTyp.Name)
			}
		case longTag != "" && shortTag == "":
			switch filedVal.Kind() {
			case reflect.Bool:
				out += fmt.Sprintf("\tf.BoolVar(&o.%s, \"%s\", false,usageDict[\"%s\"] )\n", filedTyp.Name, longTag, filedTyp.Name)
			case reflect.String:
				out += fmt.Sprintf("\tf.StringVar(&o.%s, \"%s\", \"\",usageDict[\"%s\"] )\n", filedTyp.Name, longTag, filedTyp.Name)
			case reflect.Int:
				out += fmt.Sprintf("\tf.IntVar(&o.%s, \"%s\", 0 ,usageDict[\"%s\"] )\n", filedTyp.Name, longTag, filedTyp.Name)
			}
		case longTag == "" && shortTag != "":
			hideFiledName := fmt.Sprintf("%s_%s_XXX_hide", strings.ToLower(shortTag), filedTyp.Name)
			switch filedVal.Kind() {
			case reflect.Bool:
				out += fmt.Sprintf("\tf.BoolVarP(&o.%s, \"%s\", \"%s\", false,usageDict[\"%s\"] )\n", filedTyp.Name, hideFiledName, shortTag, filedTyp.Name)
			case reflect.String:
				out += fmt.Sprintf("\tf.StringVarP(&o.%s, \"%s\", \"%s\", \"\",usageDict[\"%s\"] )\n", filedTyp.Name, hideFiledName, shortTag, filedTyp.Name)
			case reflect.Int:
				out += fmt.Sprintf("\tf.IntVarP(&o.%s, \"%s\", \"%s\", 0,usageDict[\"%s\"] )\n", filedTyp.Name, hideFiledName, shortTag, filedTyp.Name)
			}
		}
	}
	out += fmt.Sprintf(helperMapTemp, helper)
	return out
}

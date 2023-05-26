/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/5/25 上午12:07.
 * Author:  guojia(https://github.com/guojia99)
 */

package utils

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strings"
)

func SetCmdTemp(cmd *cobra.Command) {
	cmd.SetUsageFunc(func(command *cobra.Command) error {
		buf := new(bytes.Buffer)
		lines := make([]string, 0)

		maxLen := 0
		f := command.Flags()

		f.VisitAll(func(flag *pflag.Flag) {
			if flag.Hidden {
				return
			}

			line := ""
			if strings.Contains(flag.Name, "XXX_hide") {
				flag.Name = ""
			}

			if flag.Shorthand != "" && flag.ShorthandDeprecated == "" && flag.Name != "" {
				line = fmt.Sprintf("  -%s, --%s", flag.Shorthand, flag.Name)
			} else if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
				line = fmt.Sprintf("  -%s, ", flag.Shorthand)
			} else {
				line = fmt.Sprintf("      --%s", flag.Name)
			}

			varname, usage := pflag.UnquoteUsage(flag)
			if varname != "" {
				line += " " + varname
			}
			if flag.NoOptDefVal != "" {
				switch flag.Value.Type() {
				case "string":
					line += fmt.Sprintf("[=\"%s\"]", flag.NoOptDefVal)
				case "bool":
					if flag.NoOptDefVal != "true" {
						line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
					}
				case "count":
					if flag.NoOptDefVal != "+1" {
						line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
					}
				default:
					line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
				}
			}

			// This special character will be replaced with spacing once the
			// correct alignment is calculated
			line += "\x00"
			if len(line) > maxLen {
				maxLen = len(line)
			}

			line += usage
			if flag.DefValue != "" {
				if flag.Value.Type() == "string" {
					line += fmt.Sprintf(" (default %q)", flag.DefValue)
				} else {
					line += fmt.Sprintf(" (default %s)", flag.DefValue)
				}
			}
			if len(flag.Deprecated) != 0 {
				line += fmt.Sprintf(" (DEPRECATED: %s)", flag.Deprecated)
			}
			lines = append(lines, line)
		})

		for _, line := range lines {
			sidx := strings.Index(line, "\x00")
			spacing := strings.Repeat(" ", maxLen-sidx)
			// maxlen + 2 comes from + 1 for the \x00 and + 1 for the (deliberate) off-by-one in maxlen-sidx
			fmt.Println(buf, line[:sidx], spacing, wrap(maxLen+2, 0, line[sidx+1:]))
		}

		return nil
	})
}

// Splits the string `s` on whitespace into an initial substring up to
// `i` runes in length and the remainder. Will go `slop` over `i` if
// that encompasses the entire string (which allows the caller to
// avoid short orphan words on the final line).
func wrapN(i, slop int, s string) (string, string) {
	if i+slop > len(s) {
		return s, ""
	}

	w := strings.LastIndexAny(s[:i], " \t\n")
	if w <= 0 {
		return s, ""
	}
	nlPos := strings.LastIndex(s[:i], "\n")
	if nlPos > 0 && nlPos < w {
		return s[:nlPos], s[nlPos+1:]
	}
	return s[:w], s[w+1:]
}

// Wraps the string `s` to a maximum width `w` with leading indent
// `i`. The first line is not indented (this is assumed to be done by
// caller). Pass `w` == 0 to do no wrapping
func wrap(i, w int, s string) string {
	if w == 0 {
		return strings.Replace(s, "\n", "\n"+strings.Repeat(" ", i), -1)
	}

	// space between indent i and end of line width w into which
	// we should wrap the text.
	wrap := w - i

	var r, l string

	// Not enough space for sensible wrapping. Wrap as a block on
	// the next line instead.
	if wrap < 24 {
		i = 16
		wrap = w - i
		r += "\n" + strings.Repeat(" ", i)
	}
	// If still not enough space then don't even try to wrap.
	if wrap < 24 {
		return strings.Replace(s, "\n", r, -1)
	}

	// Try to avoid short orphan words on the final line, by
	// allowing wrapN to go a bit over if that would fit in the
	// remainder of the line.
	slop := 5
	wrap = wrap - slop

	// Handle first line, which is indented by the caller (or the
	// special case above)
	l, s = wrapN(wrap, slop, s)
	r = r + strings.Replace(l, "\n", "\n"+strings.Repeat(" ", i), -1)

	// Now wrap the rest
	for s != "" {
		var t string

		t, s = wrapN(wrap, slop, s)
		r = r + "\n" + strings.Repeat(" ", i) + strings.Replace(t, "\n", "\n"+strings.Repeat(" ", i), -1)
	}

	return r

}

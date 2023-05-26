/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/5/26 下午6:48.
 *  * Author: guojia(https://github.com/guojia99)
 */

package helper

var ZN = map[string]string{
	"All":              `列出所有文件, 不隐藏任何以.开始的项目`,
	"AlmostAll":        `同 -a, 但不列出“.”(表示当前目录)和“..”(表示当前目录的父目录)。`,
	"Author":           `与-l 同时使用时列出每个文件的作者`,
	"Escape":           `以八进制溢出序列表示不可打印的字符, 即为非图形字符打印 C 风格转义符`,
	"BlockSize":        `使用 -l时，打印时按 SIZE 缩放尺寸 例如，'--block-size=M'； 请参阅下面的 SIZE 格式`,
	"IgnoreBackups":    `不打印 '~' 结尾的数据`,
	"CTime":            `使用 '-lt' 时可依据ctime去做排序`,
	"Columns":          `每栏由上至下列出项目, 将文件和目录按照列的方式排列，尽可能地利用终端的宽度来显示，而不会自动换行`,
	"Color":            `color=[WHEN] 控制是否使用色彩分辨文件。WHEN 可以是’never’、’always’或’auto’其中之一`,
	"Directory":        `如果指定的目标是个目录, 则只显示该目录, 而不会递归显示出来, 且优先级最高, 不受递归影响`,
	"ForwardSort":      `不排序, 并且启用 -aU, 禁用 -ls`,
	"Classify":         `显示文件类型, 用'*/=>@|' 等套接到文件后面 /：表示目录。 *：表示可执行文件。 @：表示符号链接。 |：表示命名管道。 =：表示套接字。`,
	"ClassifyFileType": `与 --classify 类似, 但是不会显示出 *`,
	"Format": `verbose：以详细格式显示文件和目录，类似于 long，但会包括更多信息。
 commas：以逗号分隔的方式显示文件和目录。
 horizontal：以水平格式显示文件和目录。
 across：以横向格式显示文件和目录。
 vertical：以纵向格式显示文件和目录。
 single-column：以单列方式显示文件和目录。
 with-escape：以转义字符的形式显示文件和目录。
 slash：在目录名后添加斜杠 /。
 none：不使用任何格式化选项，使用默认的输出格式。`,
	"FullTime":               ``,
	"GrepNotOwnerNameOutput": ``,
	"GrepNotGroupNameOutput": ``,
	"GroupDirectoriesFirst":  ``,
	"HumanReadable":          ``,
	"HumanReadableSize1000":  ``,
	"Hide":                   ``,
	"IndicatorStyle":         ``,
	"Inode":                  ``,
	"Ignore":                 ``,
	"KibIBytes":              ``,
	"List":                   ``,
	"LinkDereference":        ``,
	"MCommaFull":             ``,
	"NumericUidGid":          ``,
	"NameOrgLiteral":         ``,
	"NotGroup":               ``,
	"HideControlChars":       ``,
	"ShowControlChars":       ``,
	"QuoteName":              ``,
	"QuotingStyle":           ``,
	"Reverse":                ``,
	"Recursive":              ``,
	"Size":                   ``,
	"SizeSort":               ``,
	"Sort":                   ``,
	"Time":                   ``,
	"TimeStyle":              ``,
	"TimeSort":               ``,
	"TimeSortUse":            ``,
	"SortNotUse":             ``,
	"VersionSort":            ``,
	"ListOutput":             ``,
	"ExtensionSort":          ``,
	"Context":                ``,
	"One":                    ``,
}

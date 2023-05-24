/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/5/24 下午4:04.
 *  * Author: guojia(https://github.com/guojia99)
 */

package ls

// Usage: ls [OPTION]... [FILE]...
//Full documentation at: <http://www.gnu.org/software/coreutils/ls>
//or available locally via: info '(coreutils) ls invocation'

type Option struct {
	LanguageHelp string `json:"language_help"` // help 按语言输出

	All              bool   `json:"all" short:"a"`                // 所有文件
	AlmostAll        bool   `json:"almost-all" short:"A"`         // 所有文件, 但是不列出 `.` 和 `..`
	Author           bool   `json:"author"`                       // -l 时生效, 额外打印出作者
	Escape           bool   `json:"escape" short:"b"`             // 为非图形字符打印 C 风格转义符
	BlockSize        int    `json:"block-size" default:"1048567"` // 在输出一个文件大小前, 按
	IgnoreBackups    int    `json:"ignore-backups" short:"B"`     // 不打印 `~` 结尾的数据
	CTime            bool   `short:"c"`                           // 使用 `-lt` 时可依据ctime去做排序
	Columns          bool   `short:"C"`                           // 将文件和目录按照列的方式排列，尽可能地利用终端的宽度来显示，而不会自动换行
	Color            string `json:"color" default:"always"`       // 颜色输出, 依据 `echo $LS_COLORS` 可获取默认的参数, 可选`auto`, `never`, `always`
	Directory        bool   `json:"directory" short:"d"`          // 如果指定的目标是个目录, 则只显示该目录, 而不会递归显示出来, 且优先级最高, 不受递归影响
	ForwardSort      bool   `short:"f"`                           // 不排序, 并且启用 -aU, 禁用 -ls
	Classify         bool   `json:"classify" short:"F"`           // 显示文件类型, 用`*/=>@|` 等套接到文件后面 /：表示目录。 *：表示可执行文件。 @：表示符号链接。 |：表示命名管道。 =：表示套接字。
	ClassifyFileType bool   `json:"file-type"`                    // Classify 类似, 但是不会显示出 *

	//Format
	//verbose：以详细格式显示文件和目录，类似于 long，但会包括更多信息。
	//commas：以逗号分隔的方式显示文件和目录。
	//horizontal：以水平格式显示文件和目录。
	//across：以横向格式显示文件和目录。
	//vertical：以纵向格式显示文件和目录。
	//single-column：以单列方式显示文件和目录。
	//with-escape：以转义字符的形式显示文件和目录。
	//slash：在目录名后添加斜杠 /。
	//none：不使用任何格式化选项，使用默认的输出格式。
	Format                 string `json:"format"`                       // long：以长格式（详细信息）显示文件和目录。
	FullTime               string `json:"full-time"`                    // 格式化时间输出
	GrepNotOwnerNameOutput string `short:"g"`                           // 使能 -l, 并且不输出文件拥有者
	GrepNotGroupNameOutput string `short:"G"`                           // 不输出文件分组信息
	GroupDirectoriesFirst  bool   `json:"group-directories-first"`      // 输出前对不同类型的文件进行分组
	HumanReadable          bool   `json:"human-readable" short:"h"`     // 人类更易读的文件大小, 在 -l 和 -s时生效
	HumanReadableSize1000  bool   `json:"si"`                           // 使能 HumanReadable, 同时输出时的幂值为 1000, 默认是1024
	Hide                   string `json:"hide"`                         // 隐藏的文件
	IndicatorStyle         string `json:"indicator-style" short:"p"`    // 将样式为 WORD 的指标附加到条目名称, 使能以下内容:  none (default), slash (-p),  file-type (--file-type), classify (-F)
	Inode                  bool   `json:"inode" short:"i"`              // 打印文件的序列号
	Ignore                 string `json:"ignore" short:"I"`             // 和 Hide 一样
	KibIBytes              bool   `json:"kibibytes" short:"k"`          // 使用1024为一个byte, 和si冲突
	List                   bool   `short:"l"`                           // 使用长输出
	LinkDereference        bool   `json:"dereference" short:"L"`        // 输出时解除引用, 只输出原始文件, 而不输出引用的文件
	MCommaFull             bool   `short:"m"`                           // 用填充短输出, 解除长输出使能
	NumericUidGid          bool   `json:"numeric-uid-gid" short:"n"`    // 输出文件归属时不输出拥有者的数据, 而是序号
	NameOrgLiteral         bool   `json:"literal" short:"N"`            // 按原始的数据输出名字, 确保文件名以原始形式显示而不是转义空格或星号。
	NotGroup               bool   `short:"o"`                           // 不输出分组信息 GrepNotGroupNameOutput
	HideControlChars       bool   `json:"hide-control-chars" short:"q"` // 隐藏图标字符为 ?
	ShowControlChars       bool   `json:"show-control-chars"`           // 默认, 展示ls的图标数据
	QuoteName              bool   `json:"quote-name" short:"Q"`         // 输出名字时带上引号
	// literal：使用不带引号的原始文件名显示。这是默认值。
	//shell：在需要时使用shell引号（例如单引号或双引号）来引用文件名。
	//shell-always：始终使用shell引号来引用文件名，即使文件名中没有特殊字符。
	//c：使用C语言风格的引号（即双引号）来引用文件名。
	//escape：使用转义字符（反斜杠）来转义文件名中的特殊字符。
	QuotingStyle string `json:"quoting-style"`       // 引号的样式
	Reverse      bool   `json:"reverse" short:"r"`   // 倒序排序
	Recursive    bool   `json:"recursive" short:"R"` // 递归输出
	//The SIZE argument is an integer and optional unit (example: 10K is 10*1024).
	//Units are K,M,G,T,P,E,Z,Y (powers of 1024) or KB,MB,... (powers of 1000).
	Size     int    `json:"size" short:"s"` // 打印每个文件的分配大小，以块为单位
	SizeSort bool   `short:"S"`             // 按文件大小排序输出
	Sort     string `json:"sort"`           // 排序的依据 none (-U), size (-S), time (-t), version (-v), extension (-X)
	// 时间的输出信息来源
	// atime：显示访问时间（access time），即文件最后一次被访问的时间。
	// ctime：显示更改时间（change time），即文件元数据最后一次更改的时间。
	// mtime：显示修改时间（modification time），即文件内容最后一次被修改的时间。
	Time string `json:"time"`
	// 时间输出格式
	// full-iso：使用完整的 ISO 8601 格式显示时间（例如："2021-09-01 12:34:56"）。
	// long-iso：使用长格式的 ISO 8601 格式显示时间（例如："2021-09-01 12:34"）。
	// iso：使用简化的 ISO 8601 格式显示时间（例如："2021-09-01"）。
	// locale：根据当前语言环境使用本地化的时间格式。
	// +"格式"：使用自定义的时间格式。您可以使用各种格式化选项，例如 %Y（年份），%m（月份），%d（日期），%H（小时），%M（分钟），%S（秒）等。
	TimeStyle     string `json:"time-style"`
	TimeSort      bool   `short:"t"`                // 按时间来排序
	TimeSortUse   bool   `short:"u"`                // 使能 TimeSort 使用-lt：按访问时间排序并显示； -l: 显示访问时间并按名称排序； 否则：按访问时间排序，最新的在前
	SortNotUse    bool   `short:"U"`                // 不使用排序, 按文件顺序输出
	VersionSort   bool   `short:"v"`                // 按自然语言(版本号)的数字做排序
	ListOutput    bool   `short:"x"`                // 按行而不是按列列出条目
	ExtensionSort bool   `short:"X"`                //按条目扩展名字母顺序排序
	Context       bool   `json:"context" short:"Z"` // 打印文件的安全上下文
	One           bool   `short:"1"`                // 按一行打印一个文件的方式输出
}

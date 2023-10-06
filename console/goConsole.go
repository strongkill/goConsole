package console

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var color = map[string][]string{
	"bold":          {"\x1B[1m", "\x1B[22m"},
	"italic":        {"\x1B[1m", "\x1B[22m"},
	"underline":     {"\x1B[4m", "\x1B[24m"},
	"inverse":       {"\x1B[7m", "\x1B[27m"},
	"strikethrough": {"\x1B[9m", "\x1B[29m"},
	"white":         {"\x1B[37m", "\x1B[39m"},
	"grey":          {"\x1B[90m", "\x1B[39m"},
	"black":         {"\x1B[30m", "\x1B[39m"},
	"blue":          {"\x1B[34m", "\x1B[39m"},
	"cyan":          {"\x1B[36m", "\x1B[39m"},
	"green":         {"\x1B[32m", "\x1B[39m"},
	"magenta":       {"\x1B[35m", "\x1B[39m"},
	"red":           {"\x1B[31m", "\x1B[39m"},
	"yellow":        {"\x1B[33m", "\x1B[39m"},

	"whiteBG":   {"\x1B[47m", "\x1B[49m"},
	"greyBG":    {"\x1B[49;5;8m", "\x1B[49m"},
	"blackBG":   {"\x1B[40m", "\x1B[49m"},
	"blueBG":    {"\x1B[44m", "\x1B[49m"},
	"cyanBG":    {"\x1B[46m", "\x1B[49m"},
	"greenBG":   {"\x1B[42m", "\x1B[49m"},
	"magentaBG": {"\x1B[45m", "\x1B[49m"},
	"redBG":     {"\x1B[41m", "\x1B[49m"},
	"yellowBG":  {"\x1B[43m", "\x1B[49m"},
}

func stringAdd(sources string, dest ...string) string {
	var pack strings.Builder
	n := 0
	for i := 0; i < len(dest); i++ {
		n += len(dest[i])
	}
	pack.Grow(n)
	pack.WriteString(sources)
	for i := 0; i < len(dest); i++ {
		pack.WriteString(dest[i])
	}
	return pack.String()
}

func getFormatTimeStr() string {
	t := time.Now()
	nanoT := strconv.FormatInt(t.UnixNano(), 10)
	return t.Format("2006-01-02 15:04:05") + "." + nanoT[10:13]
}

func getStack() string {
	var src string
	pc, _, fileLine, ok := runtime.Caller(2)
	if ok {
		src = fmt.Sprintf("%s:%d", strings.Split(runtime.FuncForPC(pc).Name(), ".")[0], fileLine)
	}
	return src
}

func Log(args ...interface{}) {
	stacker := getStack()
	colorFmt := stringAdd(strings.Join(color["bold"], getFormatTimeStr()), "  ", strings.Join(color["red"], "|"), "   LOG ", strings.Join(color["red"], "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF135    ")
	logger := make([]interface{}, len(args)+1)
	logger[0] = colorFmt
	for i, arg := range args {
		logger[i+1] = arg
	}
	fmt.Println(logger...)
}

func Debug(args ...interface{}) {
	stacker := getStack()
	colorFmt := stringAdd(strings.Join(color["bold"], getFormatTimeStr()), "  ", strings.Join(color["red"], "|"), strings.Join(color["blue"], " DEBUG "), strings.Join(color["red"], "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF188    ")
	logger := make([]interface{}, len(args)+1)
	logger[0] = colorFmt
	for i, arg := range args {
		logger[i+1] = arg
	}
	fmt.Println(logger...)
}

func Info(args ...interface{}) {
	stacker := getStack()
	colorFmt := stringAdd(strings.Join(color["bold"], getFormatTimeStr()), "  ", strings.Join(color["red"], "|"), strings.Join(color["cyan"], "  INFO "), strings.Join(color["red"], "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF1D9    ")
	logger := make([]interface{}, len(args)+1)
	logger[0] = colorFmt
	for i, arg := range args {
		logger[i+1] = arg
	}
	fmt.Println(logger...)
}

func Warn(args ...interface{}) {
	stacker := getStack()
	colorFmt := stringAdd(strings.Join(color["bold"], getFormatTimeStr()), "  ", strings.Join(color["red"], "|"), strings.Join(color["yellow"], "  WARN "), strings.Join(color["red"], "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF21E    ")
	logger := make([]interface{}, len(args)+1)
	logger[0] = colorFmt
	for i, arg := range args {
		logger[i+1] = arg
	}
	fmt.Println(logger...)
}

func Error(args ...interface{}) {
	stacker := getStack()
	colorFmt := stringAdd(strings.Join(color["bold"], getFormatTimeStr()), "  ", strings.Join(color["red"], "|"), strings.Join(color["red"], " ERROR "), strings.Join(color["red"], "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF127    ")
	logger := make([]interface{}, len(args)+1)
	logger[0] = colorFmt
	for i, arg := range args {
		logger[i+1] = arg
	}
	fmt.Println(logger...)
}

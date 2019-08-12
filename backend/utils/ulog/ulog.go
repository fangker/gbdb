package ulog

import (
	"runtime"
	"strconv"
	"fmt"
	"reflect"
	"time"
)

var log_level = TRACE;

const (
	color_black   = uint8(iota + 30)
	color_red
	color_green
	color_yellow
	color_blue
	color_magenta
	color_cyan
	color_white
)

const (
	succ  = "[SUCC]"
	info  = "[INFO]"
	capt  = "[CAPT]"
	trac  = "[TRAC]"
	erro  = "[ERRO]"
	warn  = "[WARN]"
	debug = "[DEBUG]"
)

const (
	ERROR   = iota
	WARNING
	SUCCESS
	INFO
	CAPTION
	DEBUG
	TRACE
)

const (
	log_caller = true
)

func Trace(details ...interface{}) {
	if log_level >= TRACE {
		fmt.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n", color_cyan, background(0), 1, formatLog(trac, details...))
	}
}

func Info(details ...interface{}) {
	if log_level >= INFO {
		fmt.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n", color_white, background(0), 1, formatLog(info, details...))
	}
}

func Error(details ...interface{}) {
	if log_level >= ERROR {
		fmt.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n", color_red, background(0), 1, formatLog(erro, details...))
	}
}

func Success(details ...interface{}) {
	if log_level >= SUCCESS {
		fmt.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n", color_green, background(0), 1, formatLog(succ, details...))
	}
}

func Warn(details ...interface{}) {
	if log_level >= WARNING {
		fmt.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n", color_magenta, background(0), 1, formatLog(warn, details...))
	}
}

func Caption(details ...interface{}) {
	if log_level >= CAPTION {
		fmt.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n", color_black, background(color_white), 1, formatLog(capt, details...))
	}
}

func Debug(details ...interface{}) {
	if log_level >= DEBUG && checkScopeEnable() {
		fmt.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n", color_blue, background(0), 1, formatLog(debug, details...))
	}
}
func formatLog(prefix string, details ...interface{}) string {
	var detailsInfo string
	for _, value := range details {
		detailsInfo = detailsInfo + " " + fmt.Sprint("", value)
	}
	line := fmt.Sprintf("%s%s", prefix+" : ", detailsInfo)
	if (log_caller) {
		line = getTimeStempString() + line + caller()
	}
	return line
}
func caller() string {
	_, file, line, _ := runtime.Caller(3)
	return "  <file>" + string(file) + " <line> " + strconv.Itoa(line)
}

func background(color uint8) uint8 {
	return color + 10
}

func AnyViewToString(i interface{}) string {
	rs := reflect.ValueOf(i)
	if (rs.Kind() == reflect.Slice) {
		s := "[ "
		for i := 0; i < rs.Len(); i++ {
			s = s + fmt.Sprintf("%+v ", rs.Index(i))
			if (i != rs.Len()-1) {
				s = s + "\n"
			}
		}
		s = s + " ]"
		return s
	}
	if (rs.Kind() == reflect.Struct) {
		return fmt.Sprintf("%+v", rs)
	}
	return fmt.Sprintf("%+v", rs)
}

func getTimeStempString() string {
	return fmt.Sprintf("\x1b[%d;%d;%dm%s", 1, 1, 1, time.Now().Format("2006-01-02 03:04:05")+" | ")
}

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

package utils

import (
	"log"
	"time"
	"runtime"
	"strconv"
	//"fmt"
	"fmt"
)

const (
	LOG_DEBUG = true
)

const (
	color_black = uint8(iota + 30)
	color_red
	color_green
	color_yellow
	color_blue
	color_magenta
	color_cyan
	color_white

	info = "[INFO]"
	trac = "[TRAC]"
	erro = "[ERRO]"
	warn = "[WARN]"
	succ = "[SUCC]"
	capt = "[CAPT]"
)

const (
	log_caller = false
)

func Trace(content string,  details ...interface{}) {
	log.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n",color_cyan,background(0),1,formatLog(trac,content,details...))
}

func Info(content string, details ...interface{}){
	log.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n",color_white,background(0),1,formatLog(info,content,details...))
}

func Error(content string, details ...interface{}){
	log.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n",color_red,background(0),1,formatLog(erro,content,details...))
}

func Success(content string, details ...interface{}){
	log.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n",color_green,background(0),1,formatLog(succ,content,details...))
}

func Warn(content string, details ...interface{}){
	log.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n",color_magenta,background(0),1,formatLog(warn,content,details...))
}

func Caption(content string, details ...interface{}){
	log.Printf("\x1b[%d;%d;%dm%s  \x1b[;0m \n",color_black,background(color_white),1,formatLog(capt,content,details...))
}

func formatLog(prefix,content string,details ...interface{}) string {
	log.SetFlags(0)
	var detailsInfo string
	for _,value := range details{
		detailsInfo = detailsInfo+" "+fmt.Sprint("",value)
	}
	line:= fmt.Sprintf("%s|| %s",time.Now().Format("2006-01-02 03:04:05") + " " + prefix + ": "+ content,detailsInfo)
	if(log_caller){
		line = line + caller()
	}
	return line
}
func caller() string{
	_, file, line, _ := runtime.Caller(1)
	return "  <file>"+string(file)+ " <line> "+strconv.Itoa(line)
}

func background(color uint8) uint8{
 	return  color+10
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
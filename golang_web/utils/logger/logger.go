package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint8

const ( //表示日志的级别
	DEBUG = iota + 1
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	Level       LogLevel
	LevelStr    string
	FilePath    string
	FileName    string
	MaxFileSize uint32
	Out         io.Writer
	Formatter   LogFormatter
}

type logInfo struct {
	level     LogLevel
	strLevel  string
	message   string
	timeStamp string
	fileName  string
	funcName  string
	lineNo    int
}

//定义了一个函数类型LogFormatter，用于定义日志的格式化函数
type LogFormatter func(info *logInfo) string

//定义了一个默认的日志格式化函数DefaultFormatter，根据logInfo结构体中的信息生成日志的字符串表示
var FormatterInstance = func(info *logInfo) string {
	return fmt.Sprintf("[自定义日志]"+"[%s][%s][%s:%s:%d] %s\n",
		info.timeStamp,
		info.strLevel,
		info.fileName,
		info.funcName,
		info.lineNo,
		info.message,
	)
}

func NewFileLogger(filepath, filename, level string) *Logger {
	if !strings.HasSuffix(filename, ".log") {
		filename += ".log"
	}
	out, err := os.OpenFile(path.Join(filepath, filename), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return &Logger{
		Level:     parseLogLevel(level),
		LevelStr:  level,
		FilePath:  filepath,
		FileName:  filename,
		Out:       out,
		Formatter: FormatterInstance,
	}
}

func (log *Logger) writeLog(level LogLevel, strLevel, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	strNow := time.Now().Format("2006/01/02 15:04:05")
	//获取调用该函数的代码的文件名、函数名和行号
	fileName, funcName, lineNo := getInfo(3)
	fmt.Fprint(log.Out, log.Formatter(&logInfo{
		level:     level,
		strLevel:  strLevel,
		message:   message,
		timeStamp: strNow,
		fileName:  fileName,
		funcName:  funcName,
		lineNo:    lineNo,
	}))
}

func (log *Logger) Debug(format string, args ...interface{}) {
	if log.enable(DEBUG) {
		log.writeLog(DEBUG, "DEBUG", format, args...)
	}
}

func (log *Logger) Info(format string, args ...interface{}) {
	if log.enable(INFO) {
		log.writeLog(INFO, "INFO", format, args...)
	}
}

func (log *Logger) Warning(format string, args ...interface{}) {
	if log.enable(WARNING) {
		log.writeLog(WARNING, "WARN", format, args...)
	}
}

func (log *Logger) Error(format string, args ...interface{}) {
	if log.enable(ERROR) {
		log.writeLog(ERROR, "ERROR", format, args...)
	}
}

func (log *Logger) Fatal(format string, args ...interface{}) {
	if log.enable(FATAL) {
		log.writeLog(FATAL, "FATAL", format, args...)
	}
}

func (log *Logger) enable(level LogLevel) bool {
	return level >= log.Level
}

// get func name, file name and line number when print the log
func getInfo(skip int) (fileName, funcName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]

	return
}

// 用于将字符串形式的日志级别转换为LogLevel类型。
//根据输入的字符串，将其转换为小写并进行匹配，返回相应的LogLevel值。如果字符串不匹配任何级别，则抛出异常
func parseLogLevel(level string) LogLevel {

	level = strings.ToLower(level)
	switch level {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		panic("Invalid log level")
	}

}

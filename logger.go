package log

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

const (
	DefaultCallerDepth    = 2
	defaultDatetimeFormat = "2006-01-02 15:04:05.000"
)

var (
	logger           = log.New(os.Stdout, "", 0)
	timeFormat       = defaultDatetimeFormat
	useFullLevelFlag = false
	useColors        = true
	useFullFile      = false
)

func SetTimeFormat(format string) {
	timeFormat = format
}

func UseFullLevelFlag() {
	useFullLevelFlag = true
}

func NoColors() {
	useColors = false
}

func UseFullFile() {
	useFullFile = true
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Print(v...)
}

func Debugf(format string, v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf(format, v...)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Print(v...)
}

func Infof(format string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(format, v...)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Print(v...)
}

func Warnf(format string, v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf(format, v...)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Print(v...)
}

func Errorf(format string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalf(format, v...)
}

func setPrefix(level Level) {
	levelStr := getLevelStr(level)
	datetimeStr := fmt.Sprintf("%s", time.Now().Format(timeFormat))

	var logPrefix string
	pc, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if !useFullFile {
		file = path.Base(file)
	}
	if ok {
		fn := runtime.FuncForPC(pc)
		fileLineStr := fmt.Sprintf("[%s %-4d::%s]", file+":", line, fn.Name())
		logPrefix = fmt.Sprintf("%s %s %s : ", datetimeStr, levelStr, fileLineStr)
	} else {
		logPrefix = fmt.Sprintf("%s %s : ", datetimeStr, levelStr)
	}

	logger.SetPrefix(logPrefix)
}

func getLevelStr(level Level) string {
	levelFlag := shortLevelFlags[level]
	if useFullLevelFlag {
		levelFlag = fullLevelFlags[level]
	}

	result := `[` + levelFlag + `]`
	if !useColors {
		return result
	}

	switch level {
	case DEBUG:
		return Cyan(result)
	case WARNING:
		return Yellow(result)
	case ERROR:
		return Red(result)
	case FATAL:
		return Magenta(result)
	default:
		return Green(result)
	}
}

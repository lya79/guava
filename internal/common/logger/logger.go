package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	color_red     = uint8(iota + 91) // warn
	color_green                      // info
	color_yellow                     // debug
	color_blue                       //trace
	color_magenta                    // fatal
)
const (
	// INFO 用於描述應用運行過程
	INFO = "[INFO]"

	// DEBUG 除錯用途訊息
	DEBUG = "[DEBUG]"

	// TRACE 除錯用途訊息
	TRACE = "[TRACE]"

	// WARN 存在潛在的問題
	WARN = "[WARN]"

	// FATAL 非常嚴重的錯誤, 可能導致應用終止執行.
	FATAL = "[FATAL]"
)

var (
	prefix    string
	loggerMap map[string]*log.Logger
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	loggerMap = make(map[string]*log.Logger)
}

// SetPrefix 設定前綴
func SetPrefix(str string) {
	prefix = str
	setInfoLevel(str)
	enableLevel(TRACE, WARN, FATAL)
}

func setInfoLevel(str string) {
	log.SetPrefix(str + initColor(INFO, INFO) + " ")
	log.Println("logger level啟用:", INFO)
}

func initColor(level string, tag string) string {
	switch level {
	case INFO:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, tag)
	case DEBUG:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, tag)
	case TRACE:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, tag)
	case WARN:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(91), tag)
	case FATAL:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, tag)
	}
	return tag
}

// EnableDebug 啟用指定 debug level
func EnableDebug() {
	enableLevel(DEBUG)
}

// EnableLevel 啟用指定 level的 logger
func enableLevel(levels ...string) {
	for i := range levels {
		level := levels[i]
		if _, ok := loggerMap[level]; !ok {
			loggerMap[level] = log.New(os.Stdout, prefix+initColor(level, level)+" ", log.Ldate|log.Lmicroseconds|log.Llongfile)
			log.Println("logger level啟用:", level)
		}
	}

}

// Println 輸出指定 level訊息到終端機
func Println(level string, str string, f func(logger *log.Logger, str string)) {
	if !validLevel(level) {
		Println(WARN, "無效的 logger level:", f)
		return
	}

	if logger, ok := loggerMap[level]; ok {
		f(logger, str)
	}
}

func validLevel(level string) bool {
	return level == DEBUG || level == TRACE || level == WARN || level == FATAL
}

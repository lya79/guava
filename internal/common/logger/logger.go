package logger

import (
	"log"
	"os"
)

const (
	// INFO 用於描述應用運行過程
	INFO = "[INFO]"

	// DEBUG 除錯用途訊息
	DEBUG = "[DEBUG]"

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
	enableLevel(WARN, FATAL)
}

func setInfoLevel(str string) {
	log.SetPrefix(str + INFO + " ")
	log.Println("logger level啟用:", INFO)
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
			loggerMap[level] = log.New(os.Stdout, prefix+level+" ", log.Ldate|log.Lmicroseconds|log.Llongfile)
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
	return level == DEBUG || level == WARN || level == FATAL
}

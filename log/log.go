package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m", log.LstdFlags|log.Lshortfile) //使用 log.Lshortfile 支持显示文件名和代码行号
	infoLog  = log.New(os.Stdout, "\033[34m[info]\033[0m", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

var (
	Error = errorLog.Println
	Errorf = errorLog.Printf
	Info = infoLog.Println
	Infof = infoLog.Printf
)

//log levels

const (
	InfoLevel  = iota
	ErrorLevel
	Disable
)

func SetLevel(level int)  {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers{
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level{
		errorLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level{
		infoLog.SetOutput(ioutil.Discard)
	}
}

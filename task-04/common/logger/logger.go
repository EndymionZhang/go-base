package logger

import (
	"fmt"
	"github.com/endymion/go-base/task-04/common/setting"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var loggerSetting = setting.LoggerSetting

var (
	file               *os.File
	defaultPrefix      = ""
	DefaultCallerDepth = 2
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var logger *log.Logger

func SetUp() {
	path := loggerSetting.FilePath
	name := loggerSetting.FileName
	file, err := MustOpen(name, path)
	if err != nil {
		log.Fatalf("Fail to open log file: %v", err)
	}

	logger = log.New(file, defaultPrefix, log.LstdFlags)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

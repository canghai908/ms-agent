package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//Level a
type Level int

var (
	//F as
	F *os.File
	//DefaultPrefix s
	DefaultPrefix = ""
	//DefaultCallerDepth a
	DefaultCallerDepth = 2
	//logger ss
	logger *log.Logger
	//logPrefix s
	logPrefix = ""
	//levelFlags aa
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	//DEBUG a
	DEBUG Level = iota
	//INFO a
	INFO
	//WARNING a
	WARNING
	//ERROR a
	ERROR
	//FATAL a
	FATAL
)

//Setup setup
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
		return
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

//Debug a
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)

}

//Info a
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

//Warn a
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

//Error as
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

//Fatal ldevel
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

//setPrefix a
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}

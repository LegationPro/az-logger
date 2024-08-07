package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

type LogLevel int
type Levels map[LogLevel]string
type LevelColor map[string]func(a ...interface{}) string

const FILE_PERMS = 0755

const (
	Debug LogLevel = iota
	Info  LogLevel = iota
	Warn  LogLevel = iota
	Error LogLevel = iota
	Panic LogLevel = iota
	Fatal LogLevel = iota
)

var (
	debugColor = color.New(color.FgCyan).SprintFunc()
	infoColor  = color.New(color.FgGreen).SprintFunc()
	warnColor  = color.New(color.FgYellow).SprintFunc()
	errorColor = color.New(color.FgRed).SprintFunc()
	panicColor = color.New(color.FgHiRed).SprintFunc()
	fatalColor = color.New(color.FgMagenta).SprintFunc()
)

var levels Levels = Levels{
	Debug: "debug",
	Info:  "info",
	Warn:  "warn",
	Error: "error",
	Panic: "panic",
	Fatal: "fatal",
}

var colorLevels = LevelColor{
	"debug": debugColor,
	"iinfo": infoColor,
	"warn":  warnColor,
	"error": errorColor,
	"panic": panicColor,
	"fatal": fatalColor,
}

type logger struct {
	path string
}

var (
	l    *logger
	once sync.Once
)

type message struct {
	Level   string
	Content string
	Time    time.Time
}

func setup(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, FILE_PERMS)
		if err != nil {
			panic(err)
		}
	}
}

func NewLogger(path string) {
	once.Do(func() {
		setup(path)
		l = &logger{path}
	})
}

func GetLogger() *logger {
	if l == nil {
		panic("Logger not initialized")
	}

	return l
}

func (l *logger) Log(level LogLevel, content string) {
	file, err := os.OpenFile(
		fmt.Sprintf("%s/", l.path)+fmt.Sprintf("%s.log", levels[level]),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, FILE_PERMS)

	if err != nil {
		panic(err)
	}

	msg := message{
		Level:   levels[level],
		Content: content,
		Time:    time.Now(),
	}

	data, err := json.Marshal(msg)

	if colorLevels[levels[level]] != nil {
		log.Printf("%s", colorLevels[levels[level]](string(data)))
	} else {
		log.Printf("%s", string(data))
	}

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer file.Close()

	_, err = file.WriteString(string(data) + "\n")
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func (l *logger) Info(message string) {
	l.Log(Info, message)
}

func (l *logger) Debug(message string) {
	l.Log(Debug, message)
}

func (l *logger) Warn(message string) {
	l.Log(Warn, message)
}

func (l *logger) Error(message string) {
	l.Log(Error, message)
}

func (l *logger) Panic(message string) {
	l.Log(Panic, message)
}

func (l *logger) Fatal(message string) {
	l.Log(Fatal, message)
}

func (l *logger) Clean() {
	for _, level := range levels {
		filePath := fmt.Sprintf("%s/%s.log", l.path, level)

		if _, err := os.Stat(filePath); err == nil {
			file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, FILE_PERMS)
			if err != nil {
				panic(err)
			}

			defer file.Close()
		} else if os.IsNotExist(err) {
			continue
		} else {
			panic(err)
		}
	}
}

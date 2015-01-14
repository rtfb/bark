package log

import (
	"log"
	"os"
)

type Logger struct {
	l *log.Logger
}

const (
	defaultFlags = log.Ldate | log.Ltime | log.Lshortfile
)

func Create() *Logger {
	return &Logger{log.New(os.Stderr, "", defaultFlags)}
}

func CreateFile(file string) *Logger {
	writer := os.Stderr
	if file != "" {
		f, err := os.Create(file)
		if err != nil {
			panic("log.Create: " + err.Error())
		}
		writer = f
	}
	return &Logger{log.New(writer, "", defaultFlags)}
}

// TODO
//func CreateArgs(file *os.File, prefix string, flags int) *Logger {
//}

func (l *Logger) LogErr(err error, msg string, v ...interface{}) error {
	if err == nil {
		return nil
	}
	if msg != "" {
		l.l.Printf(msg, v...)
	}
	l.l.Println(err.Error())
	return err
}

func (l *Logger) LogIf(err error) error {
	if err == nil {
		return nil
	}
	l.l.Println(err.Error())
	return err
}

func (l *Logger) Print(v ...interface{}) {
	l.l.Print(v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.l.Println(v...)
}

func (l *Logger) Printf(fmt string, v ...interface{}) {
	l.l.Printf(fmt, v...)
}

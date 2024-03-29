package bark

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rtfb/httputil"
)

type Logger struct {
	l       *log.Logger
	colored bool
}

const (
	defaultFlags = log.Ldate | log.Ltime | log.Lshortfile
)

func Create() *Logger {
	return &Logger{
		l:       log.New(os.Stderr, "", defaultFlags),
		colored: true,
	}
}

func CreateFile(file string) (*Logger, error) {
	writer := os.Stderr
	if file != "" {
		f, err := os.Create(file)
		if err != nil {
			return nil, err
		}
		writer = f
	}
	return &Logger{
		l:       log.New(writer, "", defaultFlags),
		colored: false,
	}, nil
}

func AppendFile(file string) (*Logger, error) {
	writer := os.Stderr
	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}
		writer = f
	}
	return &Logger{
		l:       log.New(writer, "", defaultFlags),
		colored: false,
	}, nil
}

// TODO
//func CreateArgs(file *os.File, prefix string, flags int) *Logger {
//}

func (l *Logger) LogIff(err error, msg string, v ...interface{}) error {
	if err == nil {
		return nil
	}
	if msg != "" {
		l.l.Output(2, fmt.Sprintf(msg, v...))
	}
	return l.Log(err)
}

func (l *Logger) LogIf(err error) error {
	if err == nil {
		return nil
	}
	return l.Log(err)
}

func (l *Logger) Log(err error) error {
	l.l.Output(2, err.Error())
	return err
}

func (l *Logger) Print(v ...interface{}) {
	l.l.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Println(v ...interface{}) {
	l.l.Output(2, fmt.Sprintln(v...))
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.l.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) LogRq(req *http.Request, startTime time.Time) {
	var logEntry bytes.Buffer
	duration := time.Now().Sub(startTime)
	ip := httputil.GetIPAddress(req)
	format := "%s - \033[32;1m %s %s\033[0m - %v"
	if !l.colored {
		format = "%s - %s %s - %v"
	}
	fmt.Fprintf(&logEntry, format, ip, req.Method, req.URL.Path, duration)
	if len(req.Form) > 0 {
		format := " - \033[37;1mParams: %v\033[0m\n"
		if !l.colored {
			format = " - Params: %v\n"
		}
		fmt.Fprintf(&logEntry, format, req.Form)
	}
	l.l.Output(2, logEntry.String())
}

package gnose

import (
	"time"
	"fmt"
	"runtime"
	"strings"
	"os"
	"io"
	"reflect"
	"runtime/debug"
	"sync"
)

const (
	info = "[INFO]"
	warning = "[WARNING]"
	errs = "[ERROR]"
	deb = "[DEBUG]"
	exception = "[EXCEPTION]"
)

type Testlogger struct {
	Path   string
	Name   string
	locker sync.Mutex
}

func NewLogger(path, logname string) Testlogger {
	return Testlogger{Path:path, Name:logname}
}

func (tl Testlogger)format() (now string) {
	now = fmt.Sprint(time.Now().Format("2006-01-02 15:04:05.999999"))
	return
}

func (tl Testlogger)writeToLogFile(arg interface{}) (err error) {
	var logname string
	var logpath string
	logpath = tl.Path
	b := []byte(logpath)
	last := string(b[len(b) - 1])
	if last == "\\" || last == "/" {
		logname = fmt.Sprintf("%s%s", logpath, tl.Name)
	} else {
		if runtime.GOOS == "windows" {
			logname = fmt.Sprintf("%s\\%s", logpath, tl.Name)
		} else {
			logname = fmt.Sprintf("%s/%s", logpath, tl.Name)
		}
	}

	var f *os.File
	if CheckFileExist(logname) {
		f, err = os.OpenFile(logname, os.O_APPEND, 0666)
	} else {
		f, err = os.Create(logname)
	}
	if err != nil {
		return
	}
	defer f.Close()
	var msg string
	msg = fmt.Sprint(arg, "\n")
	_, err = io.WriteString(f, msg)
	return
}

func (tl Testlogger)timestamp() string {
	return tl.format()
}

func (tl Testlogger)middlestring() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		fmt.Println("error occur")
	}
	filenames := strings.Split(file, "/")
	length := len(filenames)
	filename := filenames[length - 1]
	return fmt.Sprintf("[%s line:%d]", filename, line)
}

func (tl Testlogger)output(msg string) {
	// output to os.stdout
	fmt.Println(msg)
	// write to os.File
	err := tl.writeToLogFile(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func (tl Testlogger)parse(arg ...interface{}) (temp string) {
	var useFormat bool
	n := len(arg)
	if n == 0 {
		temp = "-----------------------------"
	} else if n == 1 {
		temp = fmt.Sprint(arg[0])
	} else {
		if reflect.ValueOf(arg[0]).Kind() == reflect.String {
			if str, ok := arg[0].(string); ok {
				if strings.Contains(str, "%") {
					useFormat = true
					format := str
					temp = fmt.Sprintf(format, arg[1:]...)
				}
			}
		}
		if !useFormat {
			for i, s := range arg {
				if i < 2 {
					temp = temp + fmt.Sprint(s) + "\n"
				} else {
					temp = temp + fmt.Sprint(s)
				}
			}
		}

	}
	return
}

func (tl Testlogger)pprint(tp string, arg ...interface{}) {
	input := tl.parse(arg...)
	msg := fmt.Sprintf("%s %s %s : %v", tl.timestamp(), tl.middlestring(), tp, input)
	tl.output(msg)
}

func (tl Testlogger)Info(arg ...interface{}) {
	tl.pprint(info, arg...)
}

func (tl Testlogger)Warning(arg ...interface{}) {
	tl.pprint(warning, arg...)
}

func (tl Testlogger)Error(arg ...interface{}) {
	tl.pprint(errs, arg...)
}

func (tl Testlogger)Debug(arg ...interface{}) {
	tl.pprint(deb, arg...)
}

func (tl Testlogger)Exception(arg ...interface{}) {
	defer func() {
		tl.writeToLogFile(string(debug.Stack()))
	}()
	input := tl.parse(arg...)
	msg := fmt.Sprintf("%s %s %s : %v", tl.timestamp(), tl.middlestring(), exception, input)
	err := tl.writeToLogFile(msg)
	if err != nil {
		fmt.Println(err)
	}
	panic(msg)
}
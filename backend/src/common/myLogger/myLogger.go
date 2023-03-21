package logger

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

const (
	_           = iota
	LevelDebug  //调试信息
	LevelInfo   //普通信息
	LevelWarn   //告警信息
	LevelError  //错误信息
	LevelFatal  //致命错误，非业务信息
	LevelCheck  //上线稳定后关心信息
	LevelNoShow //什么都不显示

	MAX_BUFFER_SIZE = 65536
	MAX_FILE_SIZE   = 1024 * 1024 * 50 //50M
)

var (
	file_log_name string
	file          *os.File
	dir_log_name  = "logs"
	file_name     = ""
	file_log_flag = false
	show_func     = false
	show_level    = LevelDebug
	out_put_level = LevelDebug

	log_buff          = bytes.NewBuffer(make([]byte, MAX_BUFFER_SIZE))
	out_put_log_time  = time.Second / 2
	out_put_log_chan  = make(chan string, 8196)
	out_log_chan_crit = make(chan string, 1000)

	fin_chan_I   = make(chan struct{}, 1)
	fin_chan_O   = make(chan struct{})
	enter        = "\n"
	_file_format string
	_funcName    string
	_level_list  = []string{"NOSHOW", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "CHECK"}
	createdTime  time.Time
	fileIndex    = 0
)

//设置显示log等级
func SetShowLevel(level int) {
	show_level = getLevel(level)
}

//
func SetShowFunc(show bool) {
	show_func = show
}

// 设置输出log等级
func SetOutPutLevel(level int) {
	out_put_level = getLevel(level)
}

func getLevel(level int) int {
	switch level {
	case LevelInfo:
		return LevelInfo
	case LevelDebug:
		return LevelDebug
	case LevelError:
		return LevelError
	case LevelFatal:
		return LevelFatal
	case LevelCheck:
		return LevelCheck
	case LevelNoShow:
		return LevelNoShow

	}
	return LevelDebug
}

func getLevelNum(level string) int {
	switch level {
	case "debug":
		return LevelDebug
	case "error":
		return LevelError
	case "warn":
		return LevelWarn
	case "info":
		return LevelInfo
	case "noshow":
		return LevelNoShow
	case "check":
		return LevelCheck
	case "fatal":
		return LevelFatal
	}
	return LevelDebug
}

func init() {
	if runtime.GOOS == "windows" {
		enter = "\r\n"
	} else {
		enter = "\n"
	}

	_file_format = "%s\\%s_%s_%04d.log"
	if runtime.GOOS == "windows" {
		_file_format = "%s/%s_%s_%04d.log"
	}
}

func SetOutPutFileLog(log_file_name string, dirname string) {
	file_name = log_file_name
	if dirname == "" {
		dirname = fmt.Sprintf("%s_log", file_name)
	}

	dir_log_name = dirname
	checkFileSize()
	file_log_flag = true
	log_buff.Reset()
	go outPutLogLoop()
}

func checkFileSize() {
	//判断是否存在 判断大小
	var file_info os.FileInfo
	var name string
	var err error
	now := time.Now()
	if isSameDay(now, createdTime) {
		needPlus := true
		file_info, err = os.Stat(file_log_name)
		if err != nil {
			if file != nil {
				file.Close()
				file = nil
			}
		}

		if file_info == nil {
		} else if file_info.Size() < int64(MAX_FILE_SIZE) {
			needPlus = false
		}

		if needPlus {
			createdTime = now
			fileIndex = (fileIndex + 1) % 10000
			name = fmt.Sprintf(_file_format, dir_log_name, file_name, now.Format("20230303_16_34_23"), fileIndex)
			file_log_name = name
		}
	} else {
		createdTime = now
		fileIndex = 0
		name = fmt.Sprintf(_file_format, dir_log_name, file_name, now.Format("20230303_16_34_23"), fileIndex)
		if file != nil {
			file.Close()
			file = nil
		}
		file_log_name = name
	}
	return
}

func SetOutPutIntervalTime(interval int64) {
	if interval < 1 {
		return
	}

	out_put_log_time = time.Duration(interval)
}

func Debug(v ...interface{}) {
	if show_level <= LevelDebug || (file_log_flag && out_put_level <= LevelDebug) {
		mylog(LevelDebug, show_level <= LevelDebug, out_put_level <= LevelDebug, false, "%v", v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if show_level <= LevelDebug || (file_log_flag && out_put_level <= LevelDebug) {
		mylog(LevelDebug, show_level <= LevelDebug, out_put_level <= LevelDebug, false, format, v...)
	}
}

func Info(v ...interface{}) {
	if show_level <= LevelInfo || (file_log_flag && out_put_level <= LevelInfo) {
		mylog(LevelInfo, show_level <= LevelInfo, out_put_level <= LevelInfo, false, "%v", v...)
	}
}

func Infof(format string, v ...interface{}) {
	if show_level <= LevelInfo || (file_log_flag && out_put_level <= LevelInfo) {
		mylog(LevelInfo, show_level <= LevelInfo, out_put_level <= LevelInfo, false, format, v...)
	}
}

func Warn(v ...interface{}) {
	if show_level <= LevelWarn || (file_log_flag && out_put_level <= LevelWarn) {
		mylog(LevelWarn, show_level <= LevelWarn, out_put_level <= LevelWarn, false, "%v", v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if show_level <= LevelWarn || (file_log_flag && out_put_level <= LevelWarn) {
		mylog(LevelWarn, show_level <= LevelWarn, out_put_level <= LevelWarn, false, format, v...)
	}
}

func Error(v ...interface{}) {
	if show_level <= LevelError || (file_log_flag && out_put_level <= LevelError) {
		mylog(LevelError, show_level <= LevelError, out_put_level <= LevelError, false, "%v", v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if show_level <= LevelError || (file_log_flag && out_put_level <= LevelError) {
		mylog(LevelError, show_level <= LevelError, out_put_level <= LevelError, false, format, v...)
	}
}

func Fatal(v ...interface{}) {
	if show_level <= LevelFatal || (file_log_flag && out_put_level <= LevelFatal) {
		mylog(LevelFatal, show_level <= LevelFatal, out_put_level <= LevelFatal, false, "%v", v...)
	}
}
func Fatalf(format string, v ...interface{}) {
	if show_level <= LevelFatal || (file_log_flag && out_put_level <= LevelFatal) {
		mylog(LevelFatal, show_level <= LevelFatal, out_put_level <= LevelFatal, false, format, v...)
	}
}

func Check(v ...interface{}) {
	if show_level <= LevelCheck || (file_log_flag && out_put_level <= LevelCheck) {
		mylog(LevelCheck, show_level <= LevelCheck, out_put_level <= LevelCheck, false, "%v", v...)
	}
}
func Checkf(format string, v ...interface{}) {
	if show_level <= LevelCheck || (file_log_flag && out_put_level <= LevelCheck) {
		mylog(LevelCheck, show_level <= LevelCheck, out_put_level <= LevelCheck, false, format, v...)
	}
}

func mylog(mark int, show bool, out_put bool, flag bool, format string, v ...interface{}) {
	pfuncName, filestr, line, ok := runtime.Caller(2)
	if !ok {
		filestr = "???"
		line = 0
	}

	if show_func {
		_funcName = runtime.FuncForPC(pfuncName).Name()
	} else {
		_funcName = ""
	}

	_, filename := path.Split(filestr)
	var outstring string
	if flag {
		outstring = fmt.Sprintf("%s||%s|%-16s%v%s", time.Now().Format("2006-01-02_15:04:05.000"), _level_list[mark], fmt.Sprintf("%s:%d|%s", filename, line, _funcName), fmt.Sprintf(format, v...), enter)
	} else {
		outstring = fmt.Sprintf("%s||%s|%-16s%v%s", time.Now().Format("2006-01-02_15:04:05.000"), _level_list[mark], fmt.Sprintf("%s:%d|%s", filename, line, _funcName), fmt.Sprint(v...), enter)
	}

	if show {
		fmt.Print(outstring)
	}

	if file_log_flag && out_put {
		if mark > LevelWarn {
			out_log_chan_crit <- outstring
		} else {
			out_put_log_chan <- outstring
		}
	}
}

func myLogCheck(mark int, show bool, out_put bool, flag bool, format string, v ...interface{}) {
	pfuncName, filestr, line, ok := runtime.Caller(2)
	if !ok {
		filestr = "???"
		line = 0
	}

	if show_func {
		_funcName = runtime.FuncForPC(pfuncName).Name()
	} else {
		_funcName = ""
	}

	_, filename := path.Split(filestr)
	var outstring string
	if flag {
		outstring = fmt.Sprintf("%s||%s|%-16s%v%s", time.Now().Format("2006-01-02_15:04:05.000"), _level_list[mark], fmt.Sprintf("%s:%d|%s", filename, line, _funcName), fmt.Sprintf(format, v...), enter)
	} else {
		outstring = fmt.Sprintf("%s||%s|%-16s%v%s", time.Now().Format("2006-01-02_15:04:05.000"), _level_list[mark], fmt.Sprintf("%s:%d|%s", filename, line, _funcName), fmt.Sprintf("%v|%v", v...), enter)
	}

	if show {
		fmt.Print(outstring)
	}

	if file_log_flag && out_put {
		if mark > LevelWarn {
			out_log_chan_crit <- outstring
		} else {
			out_put_log_chan <- outstring
		}
	}
}

func outPutLogLoop() {
	t := time.Now().UnixNano()
	for file_log_flag {
		select {
		case <-time.After(out_put_log_time):
			if log_buff.Len() > 0 {
				outputLog()
				t = time.Now().UnixNano()
			}

		case buff, ok := <-out_put_log_chan:
			if ok {
				if log_buff.Len()+len(buff) > MAX_BUFFER_SIZE {
					outputLog()
					if len(buff) > MAX_BUFFER_SIZE {
						outputLongLog([]byte(buff))
						break
					}
					t = time.Now().UnixNano()
				}
				log_buff.Write([]byte(buff))
			}
		case buff, ok := <-out_log_chan_crit:
			if ok {
				if log_buff.Len() > 0 {
					outputLog()
				}
				outputLongLog([]byte(buff))
				t = time.Now().UnixNano()
			}
		case <-fin_chan_I:
			cnt := len(out_put_log_chan)
			for i := cnt; i > 0; i-- {
				buff := <-out_put_log_chan
				if log_buff.Len()+len(buff) > MAX_BUFFER_SIZE {
					outputLog()
					if len(buff) > MAX_BUFFER_SIZE {
						outputLongLog([]byte(buff))
						continue
					}
					t = time.Now().UnixNano()
				}
				log_buff.Write([]byte(buff))
			}
			outputLog()
			close(fin_chan_O)
			return
		}

		if log_buff.Len() > 0 && (time.Now().UnixNano()-t) > int64(out_put_log_time) {
			outputLog()
		}
	}
}

func tryCreateLogDir() bool {
	if _, err := os.Stat(dir_log_name); err != nil {
		if err := os.Mkdir(dir_log_name, 0755); err != nil {
			fmt.Println(err, "Mkdir")
			return true
		}
	}
	return false
}

func outputLog() {
	bCreate := tryCreateLogDir()
	if file == nil || bCreate {
		var err error
		file, err = os.OpenFile(file_log_name, os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			file, err = os.Create(file_log_name)
			if err != nil {
				fmt.Println("Error!!! Create", err)
				return
			}
		}
	}
	file.Write(log_buff.Bytes())
	log_buff.Reset()
	checkFileSize()
}

func outputLongLog(longLog []byte) {
	bCreate := tryCreateLogDir()
	if file == nil || bCreate {
		var err error
		file, err = os.OpenFile(file_log_name, os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			file, err = os.Create(file_log_name)
			if err != nil {
				fmt.Println("Error!!! Create", err)
				return
			}
		}
	}
	file.Write(longLog)
	checkFileSize()
}

func Fin_log() {
	fmt.Println("Fin Log start")
	fin_chan_I <- struct{}{}
	<-fin_chan_O
	fmt.Println("Fin Log End")

}

func Exit_log() {
	fmt.Println("Exit Log start")
	fin_chan_I <- struct{}{}
	<-fin_chan_O
}

func isSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type Level int8

/*
*
1. 記錄檔分級
*/
type Fields map[string]interface{}

const (
	// iota: 自增長常數
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}

	return ""
}

/*
*
2. 記錄檔標準化
*/
type Logger struct {
	newLogger  *log.Logger
	ctx        context.Context
	fields     Fields
	callers    []string
	timeFormat string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	l.SetFlags(0)
	return &Logger{newLogger: l, timeFormat: "2006-01-02T15:04:05.999999999Z"}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// 設定記錄檔公共欄位
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}

	for k, v := range f {
		ll.fields[k] = v
	}

	return ll
}

// 設定記錄檔內容屬性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// 設定某一層呼叫堆疊的資訊
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return ll
}

// 設定整個呼叫堆疊資訊
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}

	ll := l.clone()
	ll.callers = callers
	return ll
}

/*
*
3. 記錄檔格式化和輸出
*/
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	utcTime := time.Now().UTC()
	logLine := fmt.Sprintf("%s %s", utcTime.Format(l.timeFormat), content)

	switch level {
	case LevelDebug:
		l.newLogger.Print(logLine)
	case LevelInfo:
		l.newLogger.Print(logLine)
	case LevelWarn:
		l.newLogger.Print(logLine)
	case LevelError:
		l.newLogger.Print(logLine)
	case LevelFatal:
		l.newLogger.Fatal(logLine)
	case LevelPanic:
		l.newLogger.Panic(logLine)
	}
}

/*
*
4. 記錄檔分級輸出
*/
func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}
	return l
}

func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprintf(format, v...))
}

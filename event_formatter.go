package logger

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	timeFormat = "2006-Jan-02 15:04:05.000"
	sep        = "  "
)

// Formatter is used by the logger in order to get a final log message representation.
type Formatter interface {
	message(level LogLevel, t time.Time, args ...interface{}) (string, error)
}

// JSONFormatter represents log in json format. Arguments must be provided as a separate values.
type JSONFormatter struct {
	Pretty bool
}

func (s JSONFormatter) message(level LogLevel, t time.Time, args ...interface{}) (string, error) {
	if len(args) == 1 {
		args = append([]interface{}{"msg"}, args...)
	}
	msg := make(map[string]interface{})
	msg["level"] = level.String()
	msg["t"] = t.Format(timeFormat)
	var key string
	for i := 0; i < len(args); i++ {
		if i%2 == 0 {
			v, ok := args[i].(string)
			if !ok {
				v = "undef[" + strconv.Itoa(i) + "]"
			}
			key = v
			continue
		}
		msg[key] = args[i]
	}
	var (
		data []byte
		err error
	)
	if s.Pretty {
		data, err = json.MarshalIndent(msg, "", "  ")
	} else {
		data, err = json.Marshal(msg)
	}
	return string(data) + "\n", err
}

type defaultFormatter struct{}

func (d defaultFormatter) message(level LogLevel, t time.Time, args ...interface{}) (string, error) {
	tStr := t.Format(timeFormat)
	argsStr := sprint(args...)
	switch level {
	case LevelInfo:
		return "\033[1;32m" + "[" + level.String() + " ]" + "\033[0m" + " [" + tStr + "]" + sep + argsStr + "\n", nil
	case LevelWarn:
		return "\033[1;33m" + "[" + level.String() + " ]" + " [" + tStr + "]" + sep + argsStr + "\033[0m" + "\n", nil
	case LevelError, LevelFatal:
		return "\033[1;31m" + "[" + level.String() + "]" + " [" + tStr + "]" + sep + argsStr + "\033[0m" + "\n", nil
	case LevelDebug, LevelTrace:
		return "\033[0;90m" + "[" + level.String() + "]" + " [" + tStr + "]" + sep + argsStr + "\033[0m" + "\n", nil
	}
	return "", nil
}

func sprint(args ...interface{}) string {
	ss := make([]string, len(args))
	for i, arg := range args {
		ss[i] = interfaceType(arg)
	}
	return strings.Join(ss, ",")
}

func interfaceType(arg interface{}) string {
	switch x := arg.(type) {
	case bool:
		return strconv.FormatBool(x)
	case *bool:
		return strconv.FormatBool(*x)
	case string:
		return x
	case *string:
		return *x
	case int64:
		return strconv.FormatInt(x, 10)
	case *int64:
		return strconv.FormatInt(*x, 64)
	case int32:
		return strconv.FormatInt(int64(x), 10)
	case *int32:
		return strconv.FormatInt(int64(*x), 10)
	case int16:
		return strconv.FormatInt(int64(x), 10)
	case *int16:
		return strconv.FormatInt(int64(*x), 10)
	case int8:
		return strconv.FormatInt(int64(x), 10)
	case *int8:
		return strconv.FormatInt(int64(*x), 10)
	case int:
		return strconv.Itoa(x)
	case *int:
		return strconv.Itoa(*x)
	case uint64:
		return strconv.FormatUint(x, 10)
	case *uint64:
		return strconv.FormatUint(*x, 10)
	case uint32:
		return strconv.FormatUint(uint64(x), 10)
	case *uint32:
		return strconv.FormatUint(uint64(*x), 10)
	case uint16:
		return strconv.FormatUint(uint64(x), 10)
	case *uint16:
		return strconv.FormatUint(uint64(*x), 10)
	case uint8:
		return strconv.FormatUint(uint64(x), 10)
	case *uint8:
		return strconv.FormatUint(uint64(*x), 10)
	case uint:
		return strconv.FormatUint(uint64(x), 10)
	case *uint:
		return strconv.FormatUint(uint64(*x), 10)
	case float64:
		return strconv.FormatFloat(x, 'g', -1, 64)
	case *float64:
		return strconv.FormatFloat(*x, 'g', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(x), 'g', -1, 32)
	case *float32:
		return strconv.FormatFloat(float64(*x), 'g', -1, 32)
	default:
		return fmt.Sprint(x)
	}
}

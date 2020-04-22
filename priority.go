package logger

// LogLevel is a logging level.
type LogLevel uint8

func (ll LogLevel) String() string {
	switch ll {
	case 0:
		return "TRACE"
	case 1:
		return "DEBUG"
	case 2:
		return "INFO"
	case 3:
		return "WARN"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	default:
		return "OFF"
	}
}

// LogLevel priority constants.
const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	Off
)

func isLogAllowed(setLevel, provided LogLevel) bool {
	return setLevel <= provided
}

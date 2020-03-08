package log

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	fullLevelFlags  = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	shortLevelFlags = []string{"D", "I", "W", "E", "F"}
)

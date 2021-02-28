package restart

import "fmt"

type logger interface {
	Printf(format string, a ...interface{})
}

// defaultLogger - default logger (fmt logger)
type defaultLogger struct {
}

// Printf - Printf
func (l *defaultLogger) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

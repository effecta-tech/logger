package logger

type emptyLogger struct {
}

// NewLogrusLogger create new logger using logrus logger
func NewEmptyLogger() Logger {
	return &emptyLogger{}
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *emptyLogger) Debug(args ...interface{}) {
}

// Info uses fmt.Sprint to construct and log a message.
func (l *emptyLogger) Info(args ...interface{}) {
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *emptyLogger) Warn(args ...interface{}) {
}

// Error uses fmt.Sprint to construct and log a message.
func (l *emptyLogger) Error(args ...interface{}) {
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *emptyLogger) Panic(args ...interface{}) {
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *emptyLogger) Fatal(args ...interface{}) {
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *emptyLogger) Debugf(template string, args ...interface{}) {
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *emptyLogger) Infof(template string, args ...interface{}) {
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *emptyLogger) Warnf(template string, args ...interface{}) {
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *emptyLogger) Errorf(template string, args ...interface{}) {
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *emptyLogger) Panicf(template string, args ...interface{}) {
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *emptyLogger) Fatalf(template string, args ...interface{}) {
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (l *emptyLogger) WithFields(fields Fields) Logger {
	return l
}

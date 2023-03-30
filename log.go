package log

type (
	// LogPrefix interface allow to create a child logger with prefixed output.
	LogPrefix interface {
		// SetPrefix return a new Log object with the
		// prefix param set as internal logging message prefix.
		SetPrefix(prefix string) Log

		// GetPrefix return the current logger prefix.
		GetPrefix() string
	}

	// LogMessage interface allow to otput message.
	LogMessage interface {
		Debugf(format string, v ...interface{})
		Infof(format string, v ...interface{})
		Warnf(format string, v ...interface{})
		Errorf(format string, v ...interface{})
		Fatalf(format string, v ...interface{})
	}

	// Log interface backed previous interfaces into a single one.
	Log interface {
		LogPrefix
		LogMessage

		Printf(format string, args ...interface{})
	}
)

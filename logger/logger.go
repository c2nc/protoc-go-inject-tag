package logger

import (
    "os"
    "github.com/sirupsen/logrus"
    "strings"
)

var Logger = newLogger()

func newLogger() *logrus.Logger {
    var logger = logrus.New()
    
    // Log as JSON instead of the default ASCII formatter.
    logger.Formatter = &logrus.TextFormatter{
        FullTimestamp:          true,
        DisableLevelTruncation: true,
    }
    
    // Output to stdout instead of the default stderr
    logger.Out = os.Stdout
    
    // Only log the warning severity or above.
    logger.Level = logrus.DebugLevel
    
    return logger
}

func LogLevel(level string) bool {
    switch strings.ToLower(level) {
    case "debug":
        return Logger.Level == logrus.DebugLevel
    case "info":
        return Logger.Level == logrus.InfoLevel
    case "warn":
        return Logger.Level == logrus.WarnLevel
    case "error":
        return Logger.Level == logrus.ErrorLevel
    case "fatal":
        return Logger.Level == logrus.FatalLevel
    case "panic":
        return Logger.Level == logrus.PanicLevel
    }
    
    return false
}

//type fields logrus.Fields
func Field(key string, val string) logrus.Fields {
    return logrus.Fields{key: val}
}

func SetLogLevel(level string) {
    switch strings.ToLower(level) {
    case "debug":
        Logger.Level = logrus.DebugLevel
    case "info":
        Logger.Level = logrus.InfoLevel
    case "warn":
        Logger.Level = logrus.WarnLevel
    case "error":
        Logger.Level = logrus.ErrorLevel
    case "fatal":
        Logger.Level = logrus.FatalLevel
    case "panic":
        Logger.Level = logrus.PanicLevel
    }
}

/// DEBUG

func Debug(args ...interface{}) {
    Logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
    Logger.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
    Logger.Debugln(args...)
}

/// INFO

func Info(args ...interface{}) {
    Logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
    Logger.Infof(format, args...)
}

func Infoln(args ...interface{}) {
    Logger.Infoln(args...)
}

/// WARN

func Warn(args ...interface{}) {
    Logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
    Logger.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
    Logger.Warnln(args...)
}

/// ERROR

func Error(args ...interface{}) {
    Logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
    Logger.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
    Logger.Errorln(args...)
}

func WithError(err error) {
    Logger.WithError(err)
}

/// FATAL

func Fatal(args ...interface{}) {
    Logger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
    Logger.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
    Logger.Fatalln(args...)
}

/// PANIC

func Panic(args ...interface{}) {
    Logger.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
    Logger.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
    Logger.Panicln(args...)
}
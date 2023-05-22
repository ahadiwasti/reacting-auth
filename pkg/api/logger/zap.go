package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/kjk/dailyrotate"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	//"github.com/robfig/cron"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case Infos:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Debug:
		return zapcore.DebugLevel
	case Errors:
		return zapcore.ErrorLevel
	case Fatals:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

/* For daily based logs rotation
func newRollingFile(config Configuration) zapcore.WriteSyncer {
	year, month, day := time.Now().Date()
	lj_log := lumberjack.Logger{
		Filename:   config.FileLocation +strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)+".log",
		MaxSize:    config.MaxSize,    //megabytes
		MaxAge:     config.MaxAge,     //days
		MaxBackups: 2, //files
		LocalTime:  true,
		Compress: config.Compress,
	}

	c := cron.New()
	// c.AddFunc("* * * * * *", func() { lj_log.Rotate() })
	c.AddFunc("@daily", func() { lj_log.Rotate() })
	c.Start()

	return zapcore.AddSync(&lj_log)
}
*/

func newZapLogger(config Configuration) (Logger, error) {
	cores := []zapcore.Core{}
	if config.EnableConsole {
		level := getZapLevel(config.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(config.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	if config.EnableFile {
		level := getZapLevel(config.FileLevel)

		fileFormat := func(t time.Time) string {
			return fmt.Sprintf(config.FileLocation+"%d-%d-%d.log", t.Day(), t.Month(), t.Year())
		}

		w, err := dailyrotate.NewFileWithPathGenerator(fileFormat, onClose(true, 2))
		if err != nil {
			return nil, err
		}
		w.Location = time.Local

		writer := &dailyFile{w}
		core := zapcore.NewCore(getEncoder(config.FileJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()

	return &zapLogger{
		sugaredLogger: logger,
	}, nil
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}
func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}

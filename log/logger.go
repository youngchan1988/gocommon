/*
* Author: YoungChan
* Date: 2020-07-22 17:23:08
* LastEditors: YoungChan
* LastEditTime: 2020-09-02 18:27:43
* Description: file content
 */
package log

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
)

const tag = "Logger"

var logger zerolog.Logger

func init() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	logPath := flag.String("logPath", "/var/log", "set log file storage path")
	flag.Parse()

	//初始化log 本地文件存储设置
	var logf *rotatelogs.RotateLogs
	var err error
	logFile := *logPath + "/apptoy-go"
	logf, err = rotatelogs.New(
		logFile+".%Y%m%d%H%M.log",
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithRotationTime(1*time.Minute),
	)
	if err != nil {
		ErrorF(err, tag, "Initial RotateLogs failed: %v", err)
	}

	//初始化zerolog
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	var logWriter io.Writer

	if logf != nil {
		logWriter = zerolog.MultiLevelWriter(os.Stderr, logf)
	} else {
		logWriter = os.Stderr
	}

	logger = zerolog.New(zerolog.ConsoleWriter{Out: logWriter, NoColor: false,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		},
		FormatMessage: func(i interface{}) string{
			return fmt.Sprintf("message=\"%s\"", i)
		},
		FormatCaller: func(i interface{}) string {
			if i != nil {
				return fmt.Sprintf("caller=%s", i)
			}
			return ""
		}}).With().Timestamp().Logger()

}

//Debug debug level print
func Debug(tag string, msg string) {
	logger.Debug().Str("tag", tag).Msg(msg)
}

//DebugF debug level print format
func DebugF(tag string, format string, a ...interface{}) {
	logger.Debug().Str("tag", tag).Msgf(format, a...)
}

//Info info level print
func Info(tag string, msg string) {
	logger.Info().Str("tag", tag).Msg(msg)
}

//InfoF info level print format
func InfoF(tag string, format string, a ...interface{}) {
	logger.Info().Str("tag", tag).Msgf(format, a...)
}

//Warn warn level print
func Warn(tag string, msg string) {
	logger.Warn().Str("tag", tag).Msg(msg)
}

//WarnF warn level print format
func WarnF(tag string, format string, a ...interface{}) {
	logger.Warn().Str("tag", tag).Msgf(format, a...)
}

//Error error level print
func Error(err error, tag string, msg string) {
	logger.Error().Err(err).Caller(1).Str("tag", tag).Msg(msg)
}

//ErrorF error level print format
func ErrorF(err error, tag string, format string, a ...interface{}) {
	logger.Error().Err(err).Caller(1).Str("tag", tag).Msgf(format, a...)
}

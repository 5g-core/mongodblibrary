package logger

import (
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"

	"github.com/5g-core/logger_conf"
	"github.com/5g-core/logger_util"
)

var log *logrus.Logger
var MongoDBLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	n5GCLogHook, err := logger_util.NewFileHook(logger_conf.N5GCLogfle, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(n5GCLogHook)
	}

	selfLogHook, err := logger_util.NewFileHook(logger_conf.LibLogDir+"mongodb_library.log",
		os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	MongoDBLog = log.WithFields(logrus.Fields{"component": "LIB", "category": "MonDB"})
}

func SetLogLevel(level logrus.Level) {
	MongoDBLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	MongoDBLog.Infoln("set report call :", bool)
	log.SetReportCaller(bool)
}

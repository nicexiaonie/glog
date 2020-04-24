package glog

import (
	"github.com/nicexiaonie/glog/hook"
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Path     string
	Filename string
	Level    string
	// 格式化  json  text
	Format string
	// 输出  stdin file
	Output       string
	ReportCaller bool
	// 日志分割  "20060102150405"
	Split    string
	Lifetime time.Duration
	Hook     []logrus.Hook
}

func New(c *Config) *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(c.ReportCaller)

	switch c.Format {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
		break
	case "text":
		log.SetFormatter(&logrus.TextFormatter{})
		break
	default:
		break
	}

	switch c.Level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	if len(c.Hook) > 0 {
		for _, v := range c.Hook {
			log.AddHook(v)
		}
	}

	switch c.Output {
	case "file":
		fileOut := hook.FileOut{
			FilePath: c.Path,
			FileName: c.Filename,
			Split:    c.Split,
			Lifetime: c.Lifetime,
		}
		fileOut.Init()
		log.AddHook(fileOut)
		break
	default:
		break
	}

	return log

}

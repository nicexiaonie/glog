package glog

import (
	"github.com/nicexiaonie/glog/hook"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Path     string
	Filename string
	Level    string
	Format   string
	// 输出  stdin file
	Output       string
	ReportCaller bool
	// 日志分割  "20060102150405"
	Split string
	//Lifetime time.Duration
	//Rotation time.Duration
}

func New(c *Config) *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(c.ReportCaller)

	log.SetFormatter(&logrus.JSONFormatter{})
	if c.Format == "text" {
		log.SetFormatter(&logrus.TextFormatter{})
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

	switch c.Output {
	case "file":
		fileOut := hook.FileOut{
			FilePath: c.Path,
			FileName: c.Filename,
			Split:    c.Split,
		}
		log.AddHook(fileOut)
		break
	default:
		break
	}
	return log

}

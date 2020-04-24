package glog

import (
	"github.com/nicexiaonie/glog/hook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
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
	//Rotation time.Duration
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

	switch c.Output {
	case "file":
		fileOut := hook.FileOut{
			FilePath: c.Path,
			FileName: c.Filename,
			Split:    c.Split,
		}
		log.AddHook(fileOut)
		c.lifetime()
		break
	default:
		break
	}

	return log

}

func (current *Config) lifetime() {

	if current.Lifetime == 0 {
		return
	}
	//fmt.Println(111)
	go func(current *Config) {
		for {
			time.Sleep(current.Lifetime)
			//fmt.Printf("sleep %s \n", current.Lifetime )
			files, _ := ioutil.ReadDir(current.Path)
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if current.Filename != file.Name()[0:len(current.Filename)] {
					continue
				}
				d := time.Now().Sub(file.ModTime())
				//fmt.Printf("sub %s \n", d )
				if d > (current.Lifetime) {
					fileName := current.Path + file.Name()
					//fmt.Printf("del %s \n", d )
					_ = os.Remove(fileName)
				}
				//fmt.Println(current.Lifetime * time.Second)
				//fmt.Printf("\n\n\n")
			}
		}
	}(current)
}

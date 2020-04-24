package hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type FileOut struct {
	FilePath string
	FileName string
	Split    string
}

func (FileOut) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (current FileOut) Fire(entry *logrus.Entry) error {

	msg, _ := entry.String()
	_ = os.MkdirAll(current.FilePath, os.ModePerm)
	fileName := current.getFileName()

	fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("failed to open logfile:", current.FilePath, err)
		return err
	}
	defer fd.Close()

	fd.Write([]byte(msg))
	return nil
}
func (current FileOut) getFileName() string {
	date := ""
	if len(current.Split) > 0 {
		date = time.Now().Format(current.Split)
	}
	fileName := current.FilePath + current.FileName + date
	return fileName

}

package hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

type FileOut struct {
	FilePath string
	FileName string
	Split    string
	Lifetime time.Duration
}

func (current *FileOut) Init() {

	if current.Lifetime != 0 {
		//fmt.Println(111)
		go func(current *FileOut) {
			for {

				//fmt.Printf("sleep %s \n", current.Lifetime )
				files, _ := ioutil.ReadDir(current.FilePath)
				for _, file := range files {
					if file.IsDir() {
						continue
					}
					if len(file.Name()) <= len(current.FileName) {
						continue
					}
					if current.FileName != file.Name()[0:len(current.FileName)] {
						continue
					}
					d := time.Now().Sub(file.ModTime())
					//fmt.Printf("sub %s \n", d )
					if d > (current.Lifetime) {
						fileName := current.FilePath + file.Name()
						//fmt.Printf("del %s \n", d )
						_ = os.Remove(fileName)
					}
					//fmt.Println(current.Lifetime * time.Second)
					//fmt.Printf("\n\n\n")
				}
				time.Sleep(current.Lifetime)
			}
		}(current)
	}

	go func() {
		for {
			time.Sleep(time.Second)
			linkName := current.FilePath + current.FileName
			tmpLinkName := current.FilePath + current.FileName + "_symlink"
			_ = os.Symlink(current.getFileName(), tmpLinkName)
			_ = os.Rename(tmpLinkName, linkName)
		}
	}()

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

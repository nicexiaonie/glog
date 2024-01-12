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
					//fmt.Printf("sub %s lifetime:%s \n", d, current.Lifetime)
					if d > (current.Lifetime) {
						fileName := current.FilePath + file.Name()
						fmt.Printf("del %s fileName: %s \n", d, fileName)
						_ = os.Remove(fileName)
					}
					//fmt.Println(current.Lifetime * time.Second)
					//fmt.Printf("\n\n\n")
				}
				time.Sleep(time.Second * 10)
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
	err := os.MkdirAll(current.FilePath, os.ModePerm)
	if err != nil {
		fmt.Println("failed: ", err)
		return err
	}
	fileName := current.getFileName()
	//fmt.Println("failed to open logfile:", current.FilePath, )
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

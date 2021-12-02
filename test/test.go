package main

import (
	"fmt"
	"github.com/nicexiaonie/glog"
	"runtime"
	"time"
)

func main() {

	logger := glog.New(&glog.Config{
		Path:         "/usr/glog/",
		Level:        "debug",
		Filename:     "app.log",
		Format:       "text",
		Output:       "file",
		ReportCaller: true,
		Split:        ".2006010215",
		Lifetime:     5 * time.Second,
	})
	logger.WithField("A", 345).Errorf("adfet3tewtg")

	fmt.Println(runtime.NumGoroutine())

	for {
		time.Sleep(time.Second)
		fmt.Println(runtime.NumGoroutine())
	}

	select {

	}
}

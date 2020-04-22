package main

import "github.com/nicexiaonie/glog"

func main()  {


	logger := glog.New(&glog.Config{
		Path: "/Users/nieyuanpei/Project/github/glog/log/",
		Level: "debug",
		Filename: "app.log",
		Format: "json",
		Output: "file",
		ReportCaller: true,
		Split: "2006010215",
	})
	logger.WithField("A",345).Errorf("adfet3tewtg")



}

package main

import (
	service2 "autoTestFromGo/service"
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"time"
)

const (
	seleniumPath = `./chromedriver.exe`
	port            = 9515
)
func main() {

	service, err := selenium.NewChromeDriverService(seleniumPath, port)
	if err != nil{
		log.Fatal(err)
	}
	//注意这里，server关闭之后，chrome窗口也会关闭
	defer func() {
		err := service.Stop()
		if err != nil{
			log.Fatal("selenium 释放时发生错误!-->",err)
		}
	}()

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	err = wd.Get("http://121.204.251.4:56899/")

	err = service2.Login(wd, "admin", "huvn*jCzMytim68Fmx")
	if err != nil{
		log.Print(err)
	}
	time.Sleep(time.Minute)

}

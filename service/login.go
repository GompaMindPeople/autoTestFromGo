package service

import (
	"autoTestFromGo/model"
	"errors"
	"github.com/tebeka/selenium"
	"log"
)
var login =  model.Login{}


//登录的服务,该服务仅用于 快快云
func Login(wd selenium.WebDriver,userName,password string)error{
	err := login.GetUserName(wd).SendKeys(userName)
	err = login.GetPassword(wd).SendKeys(password)
	if err != nil{
		log.Print("发送按键时,发生报错--->",err)
		return errors.New("登录发生错误!!")
	}
	err = login.GetLoginButton(wd).Click()
	if err != nil{
		log.Print("点击登录按键时,发生报错--->",err)
		return errors.New("登录发生错误!!")
	}
	return nil
}
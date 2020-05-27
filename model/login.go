package model

import (
	"github.com/tebeka/selenium"
)

type Login struct {
	BaseModel
}



func (l *Login)GetLoginButton(wd selenium.WebDriver)selenium.WebElement{
	//element, err := wd.FindElement(selenium.ByID, "login")
	element:= l.GetEleByString(wd,selenium.ByXPATH,"//*[@id=\"account-login-button\"]")
	return element
}

func (l *Login)GetUserName(wd selenium.WebDriver)selenium.WebElement{
	element:= l.GetEleByString(wd,selenium.ByXPATH,"//*[@id=\"account-name\"]/input")
	return element
}


func (l *Login)GetPassword(wd selenium.WebDriver)selenium.WebElement{
	element:= l.GetEleByString(wd,selenium.ByXPATH,"//*[@id=\"account-password\"]/input")
	return element
}
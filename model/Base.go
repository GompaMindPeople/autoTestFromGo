package model

import (
	"github.com/tebeka/selenium"
	"log"
)

type BaseModel struct {

}


//通过 字符串的方式 制定 对应的 查询类型 以及查询的 元素条件
func (l *BaseModel)GetEleByString(wd selenium.WebDriver,ByString,Value string)selenium.WebElement{
	element, err := wd.FindElement(ByString, Value)
	if err != nil{
		log.Print("获取"+Value+"的时候发生错误-->",err)
		return nil
	}
	return element
}

func (l *BaseModel)GetEleById(wd selenium.WebDriver,Value string)selenium.WebElement{
	element, err := wd.FindElement(selenium.ByXPATH, Value)
	if err != nil{
		log.Print("获取"+Value+"的时候发生错误-->",err)
		return nil
	}
	return element
}
package util

import "github.com/astaxie/beego"

/**
 * 根据开发模式进行判断是否输出日志
 */
func LogInfo(v ...interface{}) {

	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Info(v)
	}
}

/**
 * 错误
 */
func LogError(v ...interface{}) {
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Error(v)
	}
}

func LogWarn(v ...interface{}) {
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Warn(v)
	}
}

func LogDebug(v ...interface{}) {
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Debug(v)
	}
}

func LogNotice(v ...interface{}) {
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Notice(v)
	}
}

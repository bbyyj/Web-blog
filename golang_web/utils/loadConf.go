package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

//包被导入时自动执行，它用于执行包的初始化操作
func init() {
	//使用Viper来加载和解析这些配置文件
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("an error occurs when load conf file:%v", err))
	}

	//调用CreateLogger()函数来创建日志记录器实例。这个函数会根据配置文件中的设置创建日志记录器
	CreateLogger()
}

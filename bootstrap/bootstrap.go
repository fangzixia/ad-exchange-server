package bootstrap

import (
	"ad-exchange-server/config"
	"ad-exchange-server/infra/logger"
)

func Init() {
	InitInfra()
}

func InitInfra() {
	logger.InitLogger()
	config.InitConfig()
	logger.Logger.Info("基础设施初始化完成")
	logger.MediaLogger.Info("基础设施初始化完成")
	logger.ForwardLogger.Info("基础设施初始化完成")
}

func InitBusinessComponents() {

}
func InitRuleChain() {

}
func InitBudgetAdapters() {

}

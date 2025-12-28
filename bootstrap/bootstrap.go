package bootstrap

import (
	"ad-exchange-server/config"
	"ad-exchange-server/infra/logger"
	"log"
)

func Init() {

}

func InitInfra() {
	logger.InitLogger()
	config.InitConfig()
	log.Println("基础设施初始化完成")
}

func InitBusinessComponents() {

}
func InitRuleChain() {

}
func InitBudgetAdapters() {

}

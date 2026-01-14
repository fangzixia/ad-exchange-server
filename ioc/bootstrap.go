package ioc

import (
	"ad-exchange-server/config"
)

func Init() {
	InitInfra()
}

func InitInfra() {
	config.InitConfig()
}

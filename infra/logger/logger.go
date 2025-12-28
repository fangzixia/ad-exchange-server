package logger

import (
	"log"
	"os"
)

// InitLogger 初始化日志
func InitLogger() {
	// 简化实现，实际可使用Zap/Logrus
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

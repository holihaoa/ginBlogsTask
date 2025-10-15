package core

import (
	"fmt"
	"ginBlogsTask/server/global"
	"ginBlogsTask/server/initialize"
	"time"
)

func RunServer() {

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}

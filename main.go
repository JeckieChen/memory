package main

import (
	"memory/config"
	"memory/db"
	"memory/service"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	db.InitDBEngin()
	service.InitCasbinService()
	engine := gin.Default()
	config.RegisterRouter(engine)
	engine.Run(":8000")
}

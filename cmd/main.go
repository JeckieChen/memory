package main

import (
	"memory/internal/db"
	"memory/internal/router"
	"memory/internal/service"

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
	router.RegisterRouter(engine)
	engine.Run(":8000")
}

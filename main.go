package main

import (
	"fmt"
	"github.com/juliofilizzola/bot_discord/adpter/input/routes"
	"github.com/juliofilizzola/bot_discord/application/convert"
	"github.com/juliofilizzola/bot_discord/application/domain/repository"

	"github.com/juliofilizzola/bot_discord/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/bot_discord/adpter/input/controller"
	"github.com/juliofilizzola/bot_discord/application/services"
	discord2 "github.com/juliofilizzola/bot_discord/config/discord"
	"github.com/juliofilizzola/bot_discord/config/env"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/postgres"
)

func init() {
	env.Env()
}

func main() {
	env.SetEnvTerminal()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal("Error setting trusted proxies:", err)
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	webController := initDependencies()
	_, err = db.ConnectDB()
	routes.InitRoutes(&r.RouterGroup, webController)

	if err = r.Run(env.Port); err != nil {
		fmt.Println("", err)

		log.Fatal(err)
	}
}

func initDependencies() controller.WebhookControllerInterface {
	discord, err := discord2.StartDiscord()
	if err != nil {
		fmt.Printf("xza", err)
		log.Fatal(err)
	}
	connectDB, err := db.ConnectDB()
	if err != nil {
		return nil
	}

	repoUse := repository.NewUserRepository(connectDB)
	repoPr := repository.NewPRRepository(connectDB)
	service := services.NewWebhookDomainService(discord, repoPr, repoUse)
	serviceUse := services.NewUserService(repoUse)
	convert.Init(serviceUse)
	return controller.NewWebhookControllerInterface(service)
}

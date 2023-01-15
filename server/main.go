package main

import (
	"os"
	"server/db"
	"server/pkg/api/feed"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)


func main() {
	router := gin.Default()
	
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Fatal error when retrieving env variables", err)
		os.Exit(1)
	}

	feed.PostsHandler(router)
	// auth.NewAuthHandler(router)
	// feed.NewTopicHandler(router)
	// interaction.NewInteractionHandler(router)
	// chat.NewChatHandler(router)
	// hub := chat.InitHub()
	// go hub.Run()


	db.InitDB() 
	// RedisAddr := os.Getenv("REDISADDR")
	// err = db.NewRedis(RedisAddr)
	// if err != nil{
	// 	log.Fatal("Fatal error when connecting to Redis", err)
	// 	os.Exit(1)
	// }


	ListenAddr := os.Getenv("LISTENADDR")
	err = router.Run(ListenAddr)
	if err != nil {
		log.Fatal("Fatal error when running server", err)
		os.Exit(1)
	}
}


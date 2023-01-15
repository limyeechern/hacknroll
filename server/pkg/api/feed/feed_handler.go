package feed

import (

	// "server/pkg/api/auth"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/db"
	"server/pkg/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func PostsHandler(router *gin.Engine) {
	router.Use(CORS())
	topic := router.Group("api/v1")
	// topic := router.Group("api/v1/topic",)
	{
		topic.POST("/newpost", newPost)
		topic.GET("/feed/:page", getFeed)
		topic.GET("/ping", ping)
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


func newPost(c *gin.Context) {
	var newPost models.NewPost

	if err := c.BindJSON(&newPost); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("error from binding json is ", err)
        return
    }

	newPost.Date = time.Now()
	key := "data"
	value := newPost.Body
	jsonData := fmt.Sprintf(`{"%s": "%s"}`, key, value)
	data := []byte(jsonData)

	resp, err := http.Post("http://localhost:8081/getpredictions", "application/json", bytes.NewBuffer(data))

	if err != nil {
		log.Error("error is ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
	}

	var MLData []models.RedditData
	err = json.Unmarshal(body, &MLData)

	if err != nil {
		log.Error("error in line 76 is ", err)
	}

	newPost.RedditData = MLData


	err = db.NewPost(newPost)

	if err != nil {
		log.Panic("Error from db.NewProfile while registering new user: ", err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message" : "new post added"})
}


func getFeed(c *gin.Context) {
	page, err := strconv.Atoi(c.Param(("page")))
	if err != nil {
		log.Info("error is ", err)
	}
	results, pagination := db.GetFeed(page)

	response := models.Response{
		Code : http.StatusOK,
		Pagination : pagination,
		Data: results,
	}

	c.IndentedJSON(http.StatusOK, gin.H{"body" : response})
}


// func newUnopenedTopic(c *gin.Context) {
// 	profileHeader := c.GetHeader("Profile")

// 	var newTopic models.NewTopic

// 	if err := c.BindJSON(&newTopic.UnopenedSettings); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		log.Error("error from binding json is ", err)
//         return
//     }

// 	id, _ := primitive.ObjectIDFromHex(profileHeader)
// 	newTopic.Participants = append(newTopic.Participants, id)
// 	newTopic.IsOpened = false

// 	_, err := db.AddNewTopic(newTopic)

// 	if err != nil {
// 		log.Panic("Error from db.NewProfile while registering new user: ", err)
// 	}

// 	c.IndentedJSON(http.StatusOK, gin.H{"message" : "new topic added"})
// }

// func mockUnopenedTopic(c *gin.Context) {
// 	// profileHeader := c.GetHeader("Profile")
// 	profileHeader := "63b280ec6c39b617695c72a6"

// 	var newTopic models.NewTopic

// 	if err := c.BindJSON(&newTopic.UnopenedSettings); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		log.Error("error from binding json is ", err)
//         return
//     }

// 	id, _ := primitive.ObjectIDFromHex(profileHeader)
// 	newTopic.Participants = append(newTopic.Participants, id)
// 	newTopic.IsOpened = false

// 	_, err := db.AddNewTopic(newTopic)

// 	if err != nil {
// 		log.Panic("Error from db.NewProfile while registering new user: ", err)
// 	}

// 	c.IndentedJSON(http.StatusOK, gin.H{"message" : "new topic added"})
// }

// func getTopicsInConnect(c *gin.Context) {
// 	page, err := strconv.Atoi(c.Param(("page")))
// 	if err != nil{
// 		log.Error("Error while getting parameter is ", err)
// 		c.IndentedJSON(http.StatusBadRequest, models.Response{
// 			Code : http.StatusBadRequest, 
// 			Pagination: models.Pagination{}, 
// 			Data: models.Topic{},
// 		},
// 	)
// 	}
// 	mongoId := c.GetHeader("Profile")

// 	topics, pagination := db.GetTopicsInConnect(page, mongoId)
// 	response := models.Response{
// 		Code : http.StatusOK,
// 		Pagination : pagination,
// 		Data: topics,
// 	}
// 	c.IndentedJSON(http.StatusOK, response)
// }


func ping(c *gin.Context) {
	log.Info("ping")
	c.IndentedJSON(http.StatusOK, "pong ")
}

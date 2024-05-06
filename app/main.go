package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "gipogram")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbNames := []string{"users", "posts", "comments", "likes", "follow_requests", "connections"}

	for _, dbName := range dbNames {
		name := "../internal/db/" + dbName + ".sql"
		sqlFile, err := os.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(string(sqlFile))
		log.Println(string(sqlFile))
		if err != nil {
			log.Fatal(err)
		}
	}

	router.GET("/gipogram/register", func(c *gin.Context) {

	})

	router.POST("/gipogram/create-post", func(c *gin.Context) {

	})

	router.POST("/gipogram/send-follow-request", func(c *gin.Context) {

	})

	router.POST("/gipogram/accept-user-follow-request", func(c *gin.Context) {

	})

	router.POST("/gipogram/decline-user-follow-request", func(c *gin.Context) {

	})

	router.POST("/gipogram/contents/:post_id", func(c *gin.Context) {
		postId := c.Param("post_id")
		_ = postId
	})	

	address := "localhost:7777"
	log.Println("Server is listening on", address)
	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}

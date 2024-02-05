package main

import (
	"Test_1/Credentials"
	"Test_1/DB"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("Web/*/*.html")

	data := router.Group("/db")
	{
		data.GET("/getdbitem", DB.GetDBItem)
		data.POST("/additem", DB.AddItem)
	}

	info := router.Group("/info")
	{
		info.GET("/getinfo", DB.GetFrontPageInfo)
	}

	login := router.Group("/login")
	{
		login.GET("/loadlogin", Credentials.Login)
		login.POST("/login", Credentials.Login)
	}

	// Open DB (MYSQL)
	db, err := DB.OpenDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	/*router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		// AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           12 * time.Hour,
	}))*/

	router.Static("/js", "./Web/JS")
	router.Static("/css", "./Web/CSS")
	router.Static("/assets", "./Assets")
	router.Static("/images", "./Assets/Images")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mainpage.html", nil)
	})

	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Nigga this shit broke: " + err.Error())
	}
}

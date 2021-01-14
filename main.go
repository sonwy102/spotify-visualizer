package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// func main() {
//   // Set the router as the default one shipped with Gin
//   router := gin.Default()

//   // Serve frontend static files
//   // router.Use(static.Serve("/", static.LocalFile("./views", true)))

// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

// }

func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()
  
  // Serve frontend static files
  router.Use(static.Serve("/", static.LocalFile("./client", true)))

  consumerKey := os.Getenv("TWITTER_API_KEY")
	consumerSecret := os.Getenv("TWITTER_API_SECRET_KEY")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
  
  config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's Name:%+v\n", user.ScreenName)

  
  // Start and run the server
  router.Run(":5000")
}
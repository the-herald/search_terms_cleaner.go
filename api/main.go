package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"Search_Terms_Cleaner/shared"
	"Search_Terms_Cleaner/search_terms_cleaner"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load.")
	}

	router := gin.Default()

	router.GET("/accounts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"accounts": shared.GetAllAccounts(),
		})
	})

	router.POST("/run-cleaner", func(c *gin.Context) {
		accountInput := c.Query("account_input")
		if accountInput == "" {
			c.JSON(400, gin.H{"error": "Missing account_input"})
			return
		}

		results, err := search_terms_cleaner.RunCleaner(accountInput)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"scans": results})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	log.Printf("Server starting on port %s...", port)
	router.Run(":" + port)
}

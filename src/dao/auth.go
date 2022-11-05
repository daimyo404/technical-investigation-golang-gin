package dao

import (
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthCosmosDb() {

	err := godotenv.Load("/../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	endpoint := os.Getenv("COSMOSDB_ENDPOINT")
	key := os.Getenv("COSMOSDB_PRIMARY_KEY")
	databaseName := os.Getenv("COSMOSDB_DATABASE_NAME")
	containerName := os.Getenv("COSMOSDB_CONTAINER_NAME")
	partitionKey := "id"

	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		log.Fatal("Failed to create a credential: ", err)
	}

	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		log.Fatal("Failed to create Azure Cosmos DB db client: ", err)
	}

	value = readItem(client, databaseName, containerName, item.CustomerId, item.ID)
	if err != nil {
		log.Printf("readItem failed: %s\n", err)
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("hoge")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/fuga", fuga)

	router.Run()
}

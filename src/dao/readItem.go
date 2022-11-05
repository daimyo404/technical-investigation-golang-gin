package dao

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

func readItem(client *azcosmos.Client, databaseName, containerName, partitionKey, itemId string) string {

	containerClient, err := client.NewContainer(databaseName, containerName)
	//TODO: 一旦回避
	if err != nil {
		panic("Error loading .env file")
	}

	pk := azcosmos.NewPartitionKeyString(partitionKey)

	// Read an item
	ctx := context.TODO()
	itemResponse, err := containerClient.ReadItem(ctx, pk, itemId, nil)
	//TODO: 一旦回避
	// if err != nil {
	// 	return err
	// }

	return string(itemResponse.Value)

}

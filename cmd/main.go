package main

import (
	"context"
	"fmt"

	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse"
)

func main() {
	candidateID := "67f01a7f-64e2-4e40-b781-04113f1af7c5"

	client, err := megaverse.NewClient(candidateID, nil)
	if err != nil {
		fmt.Println("Error creating client:", err)
	}

	// cli.Astral.Generate(nil) // AstralService.Generate()
	// megaverse.AstralService.Delete()
	ctx := context.Background()
	res, err := client.Astral.GetMap(ctx)
	if err != nil {
		fmt.Println("Error getting map:", err)
	}

	fmt.Println(res)

	res2, err := client.Astral.GetGoalMap(ctx)
	if err != nil {
		fmt.Println("Error getting goal map:", err)
	}

	fmt.Println(res2)
}

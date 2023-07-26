package main

import (
	"context"
	"fmt"

	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse"
	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse/astral"
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

	polyanet := astral.NewPolyanet(0, 0)

	err = client.Astral.Generate(ctx, polyanet)
	if err != nil {
		fmt.Println("Error Generating astral object:", err)
	}

	cometh := astral.NewCometh(0, 1, astral.Cometh_Up)

	err = client.Astral.Generate(ctx, cometh)
	if err != nil {
		fmt.Println("Error Generating astral object:", err)
	}

	soloon := astral.NewSoloon(0, 2, astral.Soloon_Red)

	err = client.Astral.Generate(ctx, soloon)
	if err != nil {
		fmt.Println("Error Generating astral object:", err)
	}

	res, err = client.Astral.GetMap(ctx)
	if err != nil {
		fmt.Println("Error getting map:", err)
	}

	fmt.Println(res)

	err = client.Astral.Delete(ctx, polyanet)
	if err != nil {
		fmt.Println("Error Deleting astral object:", err)
	}

	err = client.Astral.Delete(ctx, cometh)
	if err != nil {
		fmt.Println("Error Deleting astral object:", err)
	}

	err = client.Astral.Delete(ctx, soloon)
	if err != nil {
		fmt.Println("Error Deleting astral object:", err)
	}

	res, err = client.Astral.GetMap(ctx)
	if err != nil {
		fmt.Println("Error getting map:", err)
	}

	fmt.Println(res)

}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse"
	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse/astral"
)

// TODO: add option to change the clientID or read from env file.
const candidateID = "67f01a7f-64e2-4e40-b781-04113f1af7c5"

func main() {
	phase := flag.String("phase", "test", "Specify the phase (phase1, phase2, or test). Default: test")

	flag.Parse()

	// Check the value of the "phase" flag and execute the corresponding code.
	switch *phase {
	case "test":
		executeTests()
	case "phase1":
		executePhase1()
	case "phase2":
		executePhase2()
	default:
		log.Fatalf("Invalid phase specified: %s. Must be 'phase1', 'phase2', or 'test'", *phase)
	}
}

func executePhase1() {
	client, err := megaverse.NewClient(candidateID, nil)
	if err != nil {
		log.Println("Error creating client:", err)
	}

	ctx := context.Background()
	goalMap, err := client.Astral.GetGoalMap(ctx)
	if err != nil {
		log.Println("Error getting goal map:", err)
	}

	cleanMap(client, ctx, goalMap)

	for row, columns := range goalMap.Goal {
		for column, astralType := range columns {
			astralObject := createAstralObject(astralType, row, column)
			if astralObject != nil {
				log.Println(row, column, astralType)
				err = client.Astral.Generate(ctx, astralObject)
				if err != nil {
					log.Println("Error Generating astral object:", err)
				}
			} else {
				client.Astral.Delete(ctx, astral.NewPolyanet(row, column))
			}
		}
	}
}

func executePhase2() {
	client, err := megaverse.NewClient(candidateID, nil)
	if err != nil {
		fmt.Println("Error creating client:", err)
	}

	ctx := context.Background()
	goalMap, err := client.Astral.GetGoalMap(ctx)
	if err != nil {
		fmt.Println("Error getting goal map:", err)
	}

	astralMap := make([][]astral.AstralObject, len(goalMap.Goal))
	for row := range astralMap {
		astralMap[row] = make([]astral.AstralObject, len(goalMap.Goal[1]))
	}

	for row, comuns := range goalMap.Goal {
		for colum, astralType := range comuns {
			fmt.Println(row, colum, astralType)
		}
	}
}

func cleanMap(client *megaverse.Client, ctx context.Context, goalMap *astral.GoalMap) {
	for row, columns := range goalMap.Goal {
		for column := range columns {
			err := client.Astral.Delete(ctx, astral.NewPolyanet(row, column))
			if err != nil {
				log.Println("Error Deleting astral object:", err)
			}
		}
	}
}

func createAstralObject(astralType string, row, column int) astral.AstralObject {
	switch {
	case astralType == astral.POLYANET:
		return astral.NewPolyanet(row, column)
	case strings.Contains(astralType, astral.COMETH):
		return astral.NewCometh(row, column, astral.Cometh_Right)
	case strings.Contains(astralType, astral.SOLOON):
		return astral.NewSoloon(row, column, astral.Soloon_Red)
	default:
		return nil
	}
}

func executeTests() {

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

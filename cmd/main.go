package main

import (
	"fmt"
	"lem-in/api/lemin"
	"lem-in/internal/util"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <input_file>")
		return
	}
	filename := os.Args[1]

	// Парсинг входных данных из файла и создание муравейника
	numberOfAnts, startRoom, endRoom, rooms, tunnels, err := util.ParseInput(filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		// Добавить вывод ошибки
		fmt.Println(err)
		return
	}

	// Создание муравейника
	farm := lemin.NewAntFarm(startRoom, endRoom, rooms, tunnels)
	// Находим все возможные пути
	paths := farm.FindAllPaths()

	// Находим группы непересекающихся путей
	groups := lemin.FindNonIntersectingPathGroups(paths)
	if len(groups) == 0 {
		fmt.Println("No viable path groups found.")
		return
	}

	// Выбор лучшей группы путей
	bestGroupIndex, minSteps, bestPaths := lemin.ChooseBestGroup(numberOfAnts, groups)
	bestGroup := groups[bestGroupIndex]

	// Показываем лучшую группу путей
	fmt.Printf("Best group chosen: Group %d with minimum steps required: %d\n", bestGroupIndex+1, minSteps)
	fmt.Println("Paths in the best group:")
	for i, path := range bestPaths {
		fmt.Printf("Path %d: %v\n", i+1, path)
	}

	// Распределение муравьев по лучшей группе путей
	ants := lemin.DistributeAnts(numberOfAnts, &bestGroup)

	// Симуляция движения муравьев
	lemin.SimulateAntsMovement(ants, farm.Start, farm.End)
}

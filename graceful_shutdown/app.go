package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const blackColor string = "\033[1;30m%s\033[0m"

var colors = []string{
	"\033[1;31m%s\033[0m",
	"\033[1;32m%s\033[0m",
	"\033[1;33m%s\033[0m",
	"\033[1;34m%s\033[0m",
	"\033[1;35m%s\033[0m",
	"\033[1;36m%s\033[0m",
}

type Character struct {
	Name        string
	Description string
}

func main() {
	printHello()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Starting random Dragon Ball characters service...")

	shutdown := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		shutdown <- true
	}()

	characterSize, characterList := readFile()

	quit := make(chan struct{})

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {

			select {
			case <-ticker.C:
				printMessage(characterSize, characterList)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	<-shutdown

	close(quit)

	fmt.Println("Process gracefully stopped.")
}

func printHello() {
	dat, err := os.ReadFile("ascii_art.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dat))
}

func readFile() (int, []Character) {
	file, err := os.Open("dragon_ball.csv")

	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		panic(err)
	}

	characterList := buildCharacterList(data)

	file.Close()

	return len(characterList), characterList
}

func buildCharacterList(data [][]string) []Character {
	var characterList []Character

	for row, line := range data {
		if row == 0 {
			continue
		}

		var character Character

		for col, field := range line {
			if col == 0 {
				character.Name = field
			} else if col == 1 {
				character.Description = field
			}
		}

		characterList = append(characterList, character)
	}

	return characterList
}

func printMessage(characterSize int, characterList []Character) {
	color := colors[rand.Intn(len(colors))]
	characterIndex := rand.Intn(characterSize)
	character := characterList[characterIndex]

	fmt.Printf(color, fmt.Sprintf("%s %s", "🐉", character.Name))
	fmt.Printf(blackColor, fmt.Sprintf(" %s\n", character.Description))
}

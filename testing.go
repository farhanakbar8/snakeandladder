package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 101

type allInformatioon struct {
	head, tail, bottom, top, position int
}

var SnLBox [N] allInformatioon

func main() {
	var play, playerName string
	var position, dice, tempPos int
	snakeAndladder := 0
	position = 0
	playerName = welcomeMessage()
	if snakeAndladder == 0 {
		firstThingToDo()
		snakeAndladder =+ 1
	}
	fmt.Printf("Press %q to roll the dice and type %q to end the game\n", "Enter", "QUIT")
	fmt.Scanln(&play)
	for play != "QUIT" {
		if position < 100 {
			dice = rollDice()
			fmt.Printf("%s's dice is %d\n", playerName, dice)
			position = position + dice
			if position > 100 {
				tempPos = position - 100
				position = 100 - tempPos
			}
			SnLBox[position].position = position
			position = checkingBox(position, playerName)
			fmt.Printf("%s's position is at %d\n", playerName, position)
			if position == 100 {
				fmt.Println("You Win!")
			} else {
				fmt.Scanln(&play)
			}
		} else {
			play = "QUIT"
		}
	}
	fmt.Println("Snakes and Ladders :")
	//for i := 0; i < N; i++ {
	//	fmt.Println(SnLBox[i], i)
	//}
	fmt.Println("Thank you for playing this game :D")
}

func rollDice() int {

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func firstThingToDo() {
	fmt.Println("Snakes :")
	randomingSnakes()
	fmt.Println("Ladders :")
	randomingLadders()
}

func randomingLadders() {
	var boxindex1 = 0
	var boxindex2 = 0
	i := 0
	rand.Seed(time.Now().UnixNano())
	for i < 7 {
		boxindex1 = (rand.Intn(80)) + 10
		boxindex2 = (rand.Intn(80)) + 10
		if SnLBox[boxindex1].top == 0 && SnLBox[boxindex1].bottom == 0 && SnLBox[boxindex2].top == 0 && SnLBox[boxindex2].bottom == 0 && SnLBox[boxindex1].head == 0 && SnLBox[boxindex1].tail == 0 && SnLBox[boxindex2].head == 0 && SnLBox[boxindex2].tail == 0 {
			if boxindex1 > boxindex2 {
				for (boxindex2 > (boxindex1 / 10) * 10) {
					boxindex2 = rand.Intn(80) + 10
				}
				SnLBox[boxindex2].top = boxindex1
				SnLBox[boxindex2].bottom = boxindex2 
				fmt.Println("The bottom of the ladder is at :" ,SnLBox[boxindex2].bottom, "And the top of the ladder is at :" ,SnLBox[boxindex2].top)
				i++
			} else if boxindex2 > boxindex1 {
				for (boxindex1 > (boxindex2 / 10) * 10) {
					boxindex1 = rand.Intn(80) + 10
				}
				SnLBox[boxindex1].top = boxindex2
				SnLBox[boxindex1].bottom = boxindex1
				fmt.Println("The bottom of the ladder is at :" ,SnLBox[boxindex1].bottom, "And the top of the ladder is at :", SnLBox[boxindex1].top)
				i++
			} else {
				boxindex1 = (rand.Intn(80)) + 10
				boxindex2 = (rand.Intn(80)) + 10
			}
		}
	}
}

func randomingSnakes() {
	var boxindex1 = 0
	var boxindex2 = 0
	i := 0
	rand.Seed(time.Now().UnixNano())
	for i < 10 { 
		boxindex1 = (rand.Intn(65)) + 26
		boxindex2 = (rand.Intn(65)) + 26
		if SnLBox[boxindex1].top == 0 && SnLBox[boxindex1].bottom == 0 && SnLBox[boxindex2].top == 0 && SnLBox[boxindex2].bottom == 0 && SnLBox[boxindex1].head == 0 && SnLBox[boxindex1].tail == 0 && SnLBox[boxindex2].head == 0 && SnLBox[boxindex2].tail == 0 {
			if boxindex1 > boxindex2 {
				for (boxindex2 > (boxindex1 / 10) * 10) {
					boxindex2 = rand.Intn(65) + 26
				}
				SnLBox[boxindex1].head = boxindex1
				SnLBox[boxindex1].tail = boxindex2
				fmt.Println("The head of the snake is at :" ,SnLBox[boxindex1].head, "The tail of the snake is at :", SnLBox[boxindex1].tail)
				i++
			} else if boxindex2 > boxindex1 {
				for (boxindex1 > (boxindex2 / 10) * 10) {
					boxindex1 = rand.Intn(65) + 26
				}
				SnLBox[boxindex2].head = boxindex2
				SnLBox[boxindex2].tail = boxindex1
				fmt.Println("The head of the snake is at :" ,SnLBox[boxindex2].head, "The tail of the snake is at :" ,SnLBox[boxindex2].tail)
				i++
			} else {
				boxindex1 = (rand.Intn(65)) + 25
				boxindex2 = (rand.Intn(65)) + 25
			}
		}
	}
}

func checkingBox(position int, playerName string) int {
	if SnLBox[position].position == SnLBox[position].bottom {
		fmt.Printf("%s hit a ladder at %d You go to %d\n Press any key to continue\n", playerName, position, SnLBox[position].top)
		position = SnLBox[position].top
	} else if SnLBox[position].position == SnLBox[position].head {
		fmt.Printf("%s hit a head of snake at %d You go to %d\n Press any key to continue\n", playerName, position, SnLBox[position].tail)
		position = SnLBox[position].tail
	}
	return position
}

func welcomeMessage() string {
	var playerName string
	fmt.Println("Welcome to Dynamic Crumb Monster Snake and Ladder Game! :D")
	fmt.Print("Enter your name : ")	
	fmt.Scanln(&playerName)
	return playerName
}
// This Code Made by Muhammad Farhan Akbar
// Student of Informatics Major at Telkom University, IF-43-INT class
// Final project for Basic Algorithm and Programming

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 101 // there are 100 boxes. In Go the index start from 0, so it is 101 instead of 100

type allInformatioon struct {
	head, tail, bottom, top, crumb, monster, position int
}

var SnLBox [N]allInformatioon

func main() {
	var play, playerName string
	var playerPosition, dice, points, snakeAndladder int
	snakeAndladder = 0
	playerPosition = 0
	points = 0
	playerName = welcomeMessage()
	checkingSnakeandLadder(&snakeAndladder)
	printInstruction()
	fmt.Scanln(&play)
	for play != "QUIT" {
		if playerPosition < 100 {
			dice = rollDice()
			fmt.Printf("%s's dice is %d\n", playerName, dice)
			playerPosition = playerPosition + dice
			checkabove100(&playerPosition)
			SnLBox[playerPosition].position = playerPosition
			CrumbandMonster(playerPosition)
			checkingBoxv2(&playerPosition, &points, &playerName)
			playerPosition = checkingBox(playerPosition, playerName)
			fmt.Printf("%s's position is at %d with %d point(s)\n", playerName, playerPosition, points)
			WinorNot(playerPosition, points, playerName, &play)
		} else {
			play = "QUIT"
		}
	}
	endMessage()
}

// if the player lands on square 100, print the winnning message. If not, the game continue.
func WinorNot(playerPosition, points int, playerName string, play *string) {
	if playerPosition == 100 {
		fmt.Println("_______________________________________")
		fmt.Println("_______________________________________")
		fmt.Printf("%s Win!\n%s has %d point(s)\n", playerName, playerName, points)
	} else {
		fmt.Scanln(play)
	}
}

// Print how to play the game
func printInstruction() {
	fmt.Println("\n------------------HOW TO PLAY THE GAME-------------------------")
	fmt.Printf("Press %s to roll the dice and type %q to end the game\n", "ENTER", "QUIT")
	fmt.Println("---------------------------------------------------------------")
}

// Print the welcome message and enter the player's name
func welcomeMessage() string {
	var playerName string
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("---Welcome to Dynamic Crumb Monster Snake and Ladder Game! :D---")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------")
	fmt.Print("Enter your name : ")
	fmt.Scanln(&playerName)
	return playerName
}

// Last message for player who finished the game
func endMessage() {
	fmt.Println("Thank you for playing this game :D")
	fmt.Println("_______________________________________")
	fmt.Println("_______________________________________")
	fmt.Println("Made by Muhammad Farhan Akbar\nIF-43-INT, SID : 1301192246\nInformatics Major at Telkom University")
}

// Check the player position, if more than 100, the player position have to go backwards
func checkabove100(playerPosition *int) {
	var tempPos int
	if *playerPosition > 100 {
		tempPos = *playerPosition - 100
		*playerPosition = 100 - tempPos
	}
}

// Check if snake and ladder already random or not
func checkingSnakeandLadder(snakeAndladder *int) {
	if *snakeAndladder == 0 {
		firstThingToDo()
		*snakeAndladder++
	}
}

// Roll the dice as usual
func rollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

// Randoming ladder and snake only once before the play begins
func firstThingToDo() {
	fmt.Println("Snakes :")
	randomingSnakes()
	fmt.Println("Ladders :")
	randomingLadders()
}

// Fill the array with 7 random ladders
func randomingLadders() {
	var boxindex1 = 0
	var boxindex2 = 0
	i := 0
	rand.Seed(time.Now().UnixNano())
	for i < 7 {
		boxindex1 = (rand.Intn(70)) + 20
		boxindex2 = (rand.Intn(79)) + 11
		if boxindex1 > boxindex2 {
			for boxindex2 > (boxindex1/10)*10 {
				boxindex2 = rand.Intn(79) + 11
			}
			if SnLBox[boxindex1].top == 0 && SnLBox[boxindex1].bottom == 0 && SnLBox[boxindex2].top == 0 && SnLBox[boxindex2].bottom == 0 && SnLBox[boxindex1].head == 0 && SnLBox[boxindex1].tail == 0 && SnLBox[boxindex2].head == 0 && SnLBox[boxindex2].tail == 0 {
				SnLBox[boxindex2].top = boxindex1
				SnLBox[boxindex2].bottom = boxindex2
				fmt.Println("The bottom of the ladder is at :", SnLBox[boxindex2].bottom, "And the top of the ladder is at :", SnLBox[boxindex2].top)
				i++
			}
		} else if boxindex2 > boxindex1 {
			for boxindex1 > (boxindex2 / 10) * 10 {
				boxindex1 = rand.Intn(79) + 11
			}
			if SnLBox[boxindex1].top == 0 && SnLBox[boxindex1].bottom == 0 && SnLBox[boxindex2].top == 0 && SnLBox[boxindex2].bottom == 0 && SnLBox[boxindex1].head == 0 && SnLBox[boxindex1].tail == 0 && SnLBox[boxindex2].head == 0 && SnLBox[boxindex2].tail == 0 {
				SnLBox[boxindex1].top = boxindex2
				SnLBox[boxindex1].bottom = boxindex1
				fmt.Println("The bottom of the ladder is at :", SnLBox[boxindex1].bottom, "And the top of the ladder is at :", SnLBox[boxindex1].top)
				i++
			}
		} else {
			boxindex1 = (rand.Intn(70)) + 20
			boxindex2 = (rand.Intn(79)) + 11
		}
	}
}

// Fill the array with 10 random snakes
func randomingSnakes() {
	var boxindex1 = 0
	var boxindex2 = 0
	i := 0
	rand.Seed(time.Now().UnixNano())
	for i < 10 {
		boxindex1 = (rand.Intn(60)) + 30
		boxindex2 = (rand.Intn(64)) + 26
		if boxindex1 > boxindex2 {
			for boxindex2 > (boxindex1 / 10) * 10 {
				boxindex2 = rand.Intn(64) + 26
			}
			if SnLBox[boxindex1].top == 0 && SnLBox[boxindex1].bottom == 0 && SnLBox[boxindex2].top == 0 && SnLBox[boxindex2].bottom == 0 && SnLBox[boxindex1].head == 0 && SnLBox[boxindex1].tail == 0 && SnLBox[boxindex2].head == 0 && SnLBox[boxindex2].tail == 0 {
				SnLBox[boxindex1].head = boxindex1
				SnLBox[boxindex1].tail = boxindex2
				fmt.Println("The head of the snake is at :", SnLBox[boxindex1].head, "The tail of the snake is at :", SnLBox[boxindex1].tail)
				i++
			}
		} else if boxindex2 > boxindex1 {
				for boxindex1 > (boxindex2 / 10) *10 {
				boxindex1 = rand.Intn(64) + 26
			}
			if SnLBox[boxindex1].top == 0 && SnLBox[boxindex1].bottom == 0 && SnLBox[boxindex2].top == 0 && SnLBox[boxindex2].bottom == 0 && SnLBox[boxindex1].head == 0 && SnLBox[boxindex1].tail == 0 && SnLBox[boxindex2].head == 0 && SnLBox[boxindex2].tail == 0 {
				SnLBox[boxindex2].head = boxindex2
				SnLBox[boxindex2].tail = boxindex1
				fmt.Println("The head of the snake is at :", SnLBox[boxindex2].head, "The tail of the snake is at :", SnLBox[boxindex2].tail)
				i++
			}
		} else {
			boxindex1 = (rand.Intn(60)) + 30
			boxindex2 = (rand.Intn(64)) + 26
		}
	}
}

// Check if the player standing on the ladder or snake
func checkingBox(playerPosition int, playerName string) int {
	if SnLBox[playerPosition].position == SnLBox[playerPosition].bottom {
		fmt.Printf("%s hit a ladder at %d %s go to %d\n", playerName, playerPosition, playerName, SnLBox[playerPosition].top)
		playerPosition = SnLBox[playerPosition].top
	} else if SnLBox[playerPosition].position == SnLBox[playerPosition].head {
		fmt.Printf("%s hit a head of snake at %d %s go to %d\n", playerName, playerPosition, playerName, SnLBox[playerPosition].tail)
		playerPosition = SnLBox[playerPosition].tail
	}
	return playerPosition
}

// Check if the player lands on the crumb or monster square
func checkingBoxv2(playerPosition, points *int, playerName *string) {
	var temp int
	if SnLBox[*playerPosition].position == SnLBox[*playerPosition].crumb {
		temp = rand.Intn(10) + 1
		rand.Seed(time.Now().UnixNano())
		fmt.Printf("%s lands on crumb square, %s get %d point(s)\n", *playerName, *playerName, temp)
		*points = *points + temp
	} else if SnLBox[*playerPosition].position == SnLBox[*playerPosition].monster {
		temp = (*points * 10) / 100
		if temp < 1 && *points != 0 {
			temp = 1
		}
		*points = *points - temp
		fmt.Printf("%s lands on monster square, %s point(s) reduce to %d point(s)\n", *playerName, *playerName, *points)
	}
}

// Random the crumb and monster everytime the player throw a dice, there is 20% chance for crumb and 10% chance for monster
func CrumbandMonster(playerPosition int) {
	var boxindex = 0
	rand.Seed(time.Now().UnixNano())
	boxindex = rand.Intn(100) + 1
	if boxindex <= 20 && boxindex > 10 {
		SnLBox[playerPosition].crumb = playerPosition
	} else if boxindex <= 10 {
		SnLBox[playerPosition].monster = playerPosition
	}
}
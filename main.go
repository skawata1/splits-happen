package main

import (
	"fmt"
	"math/rand"
	"time"
	"flag"
	"strings"
	"strconv"
)

func main() {
	var bowlingInput string
	flag.StringVar(&bowlingInput, "bowlingInput", "random", "input for bowling scores or else it is randomly generated")
	flag.Parse()

	var bowlingFrames []string

	if bowlingInput == "random" {
		fmt.Println("Bowling Scores Randomly Generated")
		bowlingFrames = generateRandomBowling()
	} else {
		fmt.Println("Bowling Scores Inputted")
		bowlingFrames = strings.Split(bowlingInput,"")
	}

	fmt.Print(bowlingFrames)
	fmt.Println("\tTotal: ", calculateBowlingScore(bowlingFrames))
}

func bowlBall(pinsLeft int, ballNumber int) (string, int) {
	rand.Seed(time.Now().UnixNano())
	ballValue := rand.Intn(pinsLeft +1)

	if ballValue == 10 {
		if ballNumber == 1 {
			return "X", 0
		} else {
			return "/", 0
		}
	} else if ballValue == pinsLeft && ballNumber == 2 {
		return "/", 0
	} else if ballValue == 0 {
		return "-", pinsLeft
	} else {
		return strconv.Itoa(ballValue), pinsLeft - ballValue
	}
}

func calculateBowlingScore(ballsThrown []string) int{
	var totalScore int

	for i, ball := range ballsThrown {
		if ball == "X" {
			totalScore += 10

			if i+3 < len(ballsThrown){
				totalScore += getMarkValue(ballsThrown[i+1], ballsThrown[i]) + getMarkValue(ballsThrown[i+2],ballsThrown[i+1])
			}
		} else if ball == "/" {
			totalScore += getMarkValue(ballsThrown[i], ballsThrown[i-1])

			if i+2 < len(ballsThrown) {
				totalScore += getMarkValue(ballsThrown[i+1], ballsThrown[i])
			}
		} else {
			totalScore += getMarkValue(ballsThrown[i], 0)
		}
	}

	return totalScore
}

func generateRandomBowling() []string {
	var ballsThrown []string

	for i:=1; i<=10; i++ {
		ball1, pinsLeft := bowlBall(10, 1)
		ballsThrown = append(ballsThrown, ball1)

		var ball2 string

		if ball1 != "X" {
			ball2, _ = bowlBall(pinsLeft, 2)
			ballsThrown = append(ballsThrown, ball2)

			if i == 10 && ball2 == "/" {
				ball3, _ := bowlBall(10, 1)
				ballsThrown = append(ballsThrown, ball3)
			}
		}

		if i == 10 && ball1 == "X" {
			ball2, pinsLeft = bowlBall(10, 1)

			ballsThrown = append(ballsThrown, ball2)

			var ball3 string
			if ball2 != "X" {
				ball3, _ = bowlBall(pinsLeft, 2)

				ballsThrown = append(ballsThrown, ball3)
			} else {
				ball3, _ = bowlBall(10, 1)

				ballsThrown = append(ballsThrown, ball3)
			}
		}

		time.Sleep(500)
	}

	return ballsThrown
}

func getMarkValue(mark interface{}, previousMark interface{}) int {
	if mark == "X" {
		return 10
	} else if mark == "/" {
		return 10 - getMarkValue(previousMark, 0)
	} else if mark == "-" {
		return 0
	} else {
		i, _ := strconv.Atoi(mark.(string))
		return i
	}
}
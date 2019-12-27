package main

import (
	"errors"
	"fmt"
)

// Creating global variables + board.
// These variables are now accessible to all functions in this file.
const g = "G" // a gopher character.
const e = "E" // an enemy.
const s = "-" // a space on the board.

// Board shouldn't be a const. You'll be adjusting it as you go.
var board = [][]string{{e, e, e, e, g},
	{e, e, e, e, e},
	{e, e, g, e, e},
	{e, e, e, e, e},
	{g, e, e, e, e}}

func stringToLocation(s string) ([2]int, error) {
	//Takes a string (eg "A1") and returns an integer array (eg [0 	0])
	m0 := make(map[byte]int)
	m1 := make(map[byte]int)

	m0['A'] = 0
	m0['B'] = 1
	m0['C'] = 2
	m0['D'] = 3
	m0['E'] = 4

	m1['1'] = 0
	m1['2'] = 1
	m1['3'] = 2
	m1['4'] = 3
	m1['5'] = 4

	row, ok1 := m0[s[0]]
	column, ok2 := m1[s[1]]
	//if the ok value is false, it means the key (_) wasn't in the map.
	if ok1 == false || ok2 == false {
		return [2]int{-1, -1}, errors.New("You supplied a position that isn't on the board")
	}

	return [2]int{row, column}, nil
}

func LocationToString(location [2]int) (string, error) {
	err := isLegalLocation(location)
	if err != nil{
		return ":(", err
	}
	m0 := make(map[int]string)
	m1 := make(map[int]string)

	m0[0] = "A"
	m0[1] = "B"
	m0[2] = "C"
	m0[3] = "D"
	m0[4] = "E"

	m1[0] = "1"
	m1[1] = "2"
	m1[2] = "3"
	m1[3] = "4"
	m1[4] = "5"

	row := m0[location[0]]
	column := m1[location[1]]

	return row + column, nil
}

func at(location [2]int) (string, error) {
	err := isLegalLocation(location)
	
	if err != nil{
		return ":(", err
	}
	
	// Everything is okay. Let's return the board's contents.
	return board[location[0]][location[1]], nil
}

func isLegalLocation(location [2]int) (error) {
	if (location[0] >= 0 && location[0] <= 4) && (location[1] >= 0 && location[1] <= 4) {
		return nil
	}
	return errors.New("You've supplied a position that isn't on the board")

}

func adjacentLocation(location [2]int, direction string) ([2]int, error){
	//return the location next to (up, down, left, right) from the one supplied
	err := isLegalLocation(location)
	if err != nil{
		return [2]int{-1,-1}, errors.New("You've supplied a position that isn't on the board")
	}

	row := location[0]
	column := location[1]

	if direction == "up"{
		return [2]int{row-1, column}, nil
	} else if direction == "down"{
		return [2]int{row+1, column}, nil
	} else if direction == "left"{
		return [2]int{row, column-1}, nil
	} else if direction == "right"{
		return [2]int{row, column+1}, nil
	}
	return location, errors.New("")
}

func main() {
	fmt.Println("")
	

}

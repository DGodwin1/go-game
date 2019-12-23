package main
import (
	"errors"
	)

// Creating global variables + board outside of any function. 
// These variables are now accessible to all functions in this file.
const g = "G" // a gopher character.
const e = "E" // an enemy. 
const s = "-" // a space on the board.

// Board shouldn't be a const as you'll be adjusting it as you go.
var board = [][]string {{e, e, e, e, g},
				{e, e, e, e, e},
				{e, e, g, e, e},
				{e, e, e, e, e},
				{g, e, e, e, e}}

func stringToLocation(s string) ([2]int, error) {
	//Takes a string s (eg "A1") and returns an integer array (eg [0,0])

	//declare some maps for position [0] and [1] from the string.
	m0 := make(map[byte]int)
	m1 := make(map[byte]int)

	//make the mapping
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
		return [2]int{-1,-1}, errors.New("You supplied a position that isn't on the board")
	}

	return [2]int{row, column}, nil
}

func locationToString(location [2]int)(string){
	return ""
}

func main(){
	
}


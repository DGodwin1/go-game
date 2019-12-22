package main
import ("fmt")

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

func stringToLocation(s string) ([2]int){
	//TODO: add a check in case the location added isn't legal. 

	//Takes a string s (eg "A1") and returns an integer array (eg [0,0])
	m0 := make(map[byte]int)
	m0['A'] = 0
	m0['B'] = 1
	m0['C'] = 2
	m0['D'] = 3
	m0['E'] = 4

	m1 := make(map[byte]int)
	m1['1'] = 0
	m1['2'] = 1
	m1['3'] = 2
	m1['4'] = 3
	m1['5'] = 4
	return [2]int{m0[s[0]],m1[s[1]]}
}

func main(){
	fmt.Println(stringToLocation("A1")) //[0,0]
	fmt.Println(stringToLocation("F1")) //needs to error
	fmt.Println(stringToLocation("A8")) //needs to error
	
}
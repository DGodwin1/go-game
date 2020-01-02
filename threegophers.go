package main

import (
	"errors"
	"fmt"
)

// Creating global variables + board.
// These variables are now accessible to all functions in this file.
const g string = "G" // a gopher character.
const e string = "E" // an enemy.
const s string = "-" // a space on the board.

// Board shouldn't be a const. You'll be adjusting it as you go.
var board = [][]string{{e, e, e, e, g},
	{e, s, e, e, e},
	{e, e, g, e, e},
	{e, e, e, e, e},
	{g, e, e, e, e}}

func StringToLocation(s string) ([2]int, error) {
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
	//if the ok value is false, it means the key wasn't in the map.
	if ok1 == false || ok2 == false {
		return [2]int{-1, -1}, errors.New("can't make location from string as location wouldn't technically be on the board")
	}

	return [2]int{row, column}, nil
}

func LocationToString(location [2]int) (string, error) {
	_, err := isLegalLocation(location)
	if err != nil {
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
	_, err := isLegalLocation(location)

	if err != nil {
		return ":(", err
	}

	// Everything is okay. Let's return the board's contents.
	return board[location[0]][location[1]], nil
}

func isLegalLocation(location [2]int) (bool, error) {
	//is the location being given on the board?
	if (location[0] >= 0 && location[0] <= 4) && (location[1] >= 0 && location[1] <= 4) {
		return true, nil
	}
	return false, errors.New("the position you've supplied isn't legal I'm afraid")
}

func adjacentLocation(location [2]int, direction string) ([2]int, error) {
	//return the location next to (up, down, left, right) from the one supplied

	_, err := isLegalLocation(location) //is the starting position okay?
	if err != nil {
		return [2]int{-1, -1}, err
	}
	row := location[0]
	column := location[1]

	if direction == "up" {
		return [2]int{row - 1, column}, nil
	} else if direction == "down" {
		return [2]int{row + 1, column}, nil
	} else if direction == "left" {
		return [2]int{row, column - 1}, nil
	} else if direction == "right" {
		return [2]int{row, column + 1}, nil
	}
	return location, errors.New("hmmm. seems there was an issue. it might be down to what direction you gave. maybe try again?")
}

func isLegalMoveByGopher(location [2]int, direction string) (bool, error) {
	/*Tests if the G at the location can move in the direction specified.
	Assume the input will always be in correct range. ie, not 5,5
	Returns an error if the starting character is not a gopher */

	//is the character a Gopher?
	c, err := at(location)
	if err != nil || c != "G" {
		return false, err
	}

	//where does the Gopher want to move?
	gDest, err := adjacentLocation(location, direction)
	if err != nil {
		return false, err
	}

	//is the destination a legal position?
	_, err = isLegalLocation(gDest)
	if err != nil {
		return false, errors.New("you're trying to move to a location that isn't on the board I'm afraid")
	}

	//destination is legal. is the character at the destination an enemy that we can eat?
	cd, err := at(gDest)
	if err != nil {
		return false, errors.New(board[gDest[0]][gDest[1]])
	} else if cd != "E" {
		return false, errors.New("terribly sorry. you're trying to eat something that isn't an enemy and you're not allowed to do that.")
	} else {
		return true, nil
	}
}

func isLegalMoveByEnemy(location [2]int, direction string) (bool, error) {
	//Can the enemy move into an empty space? Let's find out.

	//is the character at the present location an E?
	character, err := at(location)
	if err != nil || character != "E" {
		return false, err
	}

	//where does the enemy want to move?
	eDest, err := adjacentLocation(location, direction)
	if err != nil {
		return false, err
	}
	//is the destination a legal position?
	_, err = isLegalLocation(eDest)
	if err != nil {
		return false, errors.New("you're trying to move to a location that isn't on the board I'm afraid")
	}

	//destination is legal. is there a space free at the destination that we can nab?
	isSpace, err := at(eDest)
	if err != nil {
		return false, err
	} else if isSpace != "-" {
		return false, errors.New("whoops! you're trying to move into something that isn't an empty.")
	} else {
		return true, nil
	}
}

func isLegalMove(location [2]int, direction string) (bool, error) {
	//get the character from the initial location.
	//once you've found that out, you can throw that to the other isLegal functions
	character, err := at(location)
	if err != nil {
		return false, err
	}
	if character == "E" { //if an enemy, check they can make a move.
		return isLegalMoveByEnemy(location, direction)
	} else if character == "G" {
		return isLegalMoveByGopher(location, direction)
	} else {
		return false, errors.New("hmmm. can't check if it's a legal move becuase it probably isn't a 'real' character.")
	}
}

func hasOneMoveAvailable(location [2]int) (bool, error) {
	//Tests whether the player at the location has at least one move available.
	
	//is the location legit to test?
	_, err := isLegalLocation(location)
	if err != nil{
		return false, err
	}

	
	
	directions := []string{"left", "right", "up", "down"}
	for _, i := range(directions){
		fmt.Println(location, i) //need to make sure that the adjacentlocation is legal.

		//then check if the destination of the move is legit for the character.
	}
	return true, errors.New("hello")
	
}
    
func main() {
	fmt.Println(hasOneMoveAvailable([2]int{0, 4}))
	// fmt.Println(isLegalMoveByGopher([2]int{0, 4}, "right")) //can't do it.
	// fmt.Println(isLegalMoveByGopher([2]int{0, 4}, "left"))  //can do it.
	// fmt.Println(isLegalMoveByGopher([2]int{0, 3}, "left"))  //nope. it's not a gopher
	// fmt.Println(isLegalMoveByGopher([2]int{2, 2}, "up"))    //nope. it's not an enemy
	// fmt.Println(isLegalMoveByEnemy([2]int{1, 1}, "right")) //yep.
	// fmt.Println(isLegalMoveByEnemy([2]int{2, 2}, "right")) //nope it's a gopher
	// fmt.Println(isLegalMoveByEnemy([2]int{0, 0}, "left")) //nope. you're going off the board.
	// fmt.Println(isLegalMoveByEnemy([2]int{0, 0}, "left")) //nope. you're going off the board.
}

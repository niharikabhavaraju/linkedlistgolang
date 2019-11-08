package main

import "fmt"

// clue has hint to move forward in treasure hunt game until gold is found
type clue struct {
	name        string
	description string
	next        *clue
}

// treasurehunt game has clues and goal to find gold
type treasurehunt struct {
	name        string
	goal        string
	head        *clue
	currentClue *clue
}

// createtreasurehunt creates treasurehuntgame
func createtreasurehunt(name string, goal string) *treasurehunt {
	return &treasurehunt{
		name: name,
		goal: goal,
	}
}

// addClue adds clues in treasurehunt game
func (t *treasurehunt) addClue(name, description string) error {
	c := &clue{
		name:        name,
		description: description,
	}
	if t.head == nil {
		t.head = c
	} else {
		currentNode := t.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = c
	}
	return nil
}

// gerCurrentClue returns the current clue in game
func (t *treasurehunt) getCurrentClue() *clue {
	t.currentClue = t.head
	return t.currentClue
}

// nextClue returns the next clue in game
func (t *treasurehunt) nextClue() *clue {
	t.currentClue = t.currentClue.next
	return t.currentClue
}

// getAllClues returns all clues in game
func (t *treasurehunt) getAllClues() error {
	currentNode := t.head
	if currentNode == nil {
		fmt.Println("Treasure hunt is empty,not defined.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}

// playGame plays the game iterating through clues until gold is found
func (t *treasurehunt) playGame() {

	for {
		t.nextClue()
		if t.currentClue.name == "gold" {
			break
		}
		fmt.Println("go to next clue...")
		fmt.Printf("Clue: %s by %s\n", t.currentClue.name, t.currentClue.description)
	}

	fmt.Println("Found " + t.currentClue.name)
}

func main() {
	//create treasurehunt game
	game := createtreasurehunt("find gold", "gold")
	fmt.Println("treasurehunt created")
	fmt.Println()

	//add clues to the game
	fmt.Print("Adding clues to the game...\n\n")
	game.addClue("clue1", "desc1")
	game.addClue("clue2", "desc2")
	game.addClue("clue3", "desc3")
	game.addClue("clue4", "desc4")
	game.addClue("gold", "gold")

	fmt.Println("Showing all clues in game")
	//shows all clues in game
	game.getAllClues()
	fmt.Println()

	//shows current clue in game
	game.getCurrentClue()
	fmt.Printf("Now playing: %s by %s\n", game.currentClue.name, game.currentClue.description)
	fmt.Println()

	//play game until gold is found
	game.playGame()
}

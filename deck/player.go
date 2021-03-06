package deck

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PlayerAction byte

const (
	STAND      PlayerAction = 'S'
	HIT        PlayerAction = 'H'
	DOUBLE     PlayerAction = 'D'
	SPLIT      PlayerAction = 'F'
	NULLACTION PlayerAction = '0'
)

type Player struct {
	chips map[int]int
	hands []*Hand
}

func GetPlayerInput() PlayerAction {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter action ([H]it, [S]tand, [D]ouble, or [F]ork/Split): ")
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, "\n", "")
	fmt.Println("Input string: \"" + input + "\"")

	if input == "s" || input == "stand" {
		return STAND
	} else if input == "h" || input == "hit" {
		return HIT
	} else {
		return NULLACTION
	}
}

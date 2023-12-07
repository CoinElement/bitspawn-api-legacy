package enum

import "fmt"

type Console string

const (
	Nintendo    Console = "NINTENDO"
	NintendoWii Console = "NINTENDO_WII"
	XboxX       Console = "XBOX_SERIES_X"
	XboxOne     Console = "XBOX_ONE"
	PS4         Console = "PLAYSTATION_4"
	PS5         Console = "PLAYSTATION_5"
	PC          Console = "PC"
	Cross       Console = "CROSS_PLAY"
)

func (c Console) IsValid() bool {
	if c != Nintendo && c != NintendoWii && c != XboxX && c != XboxOne && c != PS4 && c != PS5 && c != PC && c != Cross {
		return false
	}
	return true
}

func (c Console) ToString() string {
	return string(c)
}

func ConsoleJoin(consoles []Console) string {
	strConsoles := ""
	for _, console := range consoles {
		strConsoles = fmt.Sprintf("%s,%s", strConsoles, console.ToString())
	}
	if len(strConsoles) > 0 {
		strConsoles = strConsoles[1:]
	}
	return strConsoles
}

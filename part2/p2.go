package part2

import (
	"fmt"
	"os"
	"strings"
)

func P1() {

	const (
		PlayerRock     = "X"
		PlayerPaper    = "Y"
		PlayerScisiors = "Z"
		OponentRock    = "A"
		OponentPaper   = "B"
		OponentScisors = "C"
		Rock           = 1
		Paper          = 2
		Scicosrs       = 3
		Win            = 6
		Draw           = 3
		Loss           = 0
	)
	val, _ := os.ReadFile("./part2/input.txt")

	stringVal := string(val)

	games := strings.Split(stringVal, "\n")

	sum := 0
	for _, game := range games {
		plays := strings.Split(game, " ")
		if len(plays) == 1{

            fmt.Println("asdfasdfasdfadsaf")
			break
		}
		oponent := plays[0]
		player := plays[1]

		switch oponent {
		case OponentRock:
			if player == PlayerRock {
				sum += Rock + Draw
			} else if player == PlayerPaper {
				sum += Paper + Win
			} else if player == PlayerScisiors {
				sum += Scicosrs + Loss
			}
		case OponentPaper:
			if player == PlayerRock {
				sum += Rock + Loss
			} else if player == PlayerPaper {
				sum += Paper + Draw
			} else if player == PlayerScisiors {
				sum += Scicosrs + Win
			}

		case OponentScisors:
			if player == PlayerRock {
				sum += Rock + Win
			} else if player == PlayerPaper {
				sum += Paper + Loss
			} else if player == PlayerScisiors {
				sum += Scicosrs + Draw
			}

		}

	}

	fmt.Println(sum)
}

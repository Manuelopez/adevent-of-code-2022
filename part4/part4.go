package part4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	first  string
	second string
}

func P2() {

	val, _ := os.ReadFile("./part4/input.txt")

	stringVal := string(val)

	pairs := strings.Split(stringVal, "\n")

	var elves []Pair
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		groups := strings.Split(pair, ",")
		elves = append(elves, Pair{first: groups[0], second: groups[1]})
	}

	sum := 0
	for _, elve := range elves {
		sections1 := strings.Split(elve.first, "-")
		sections2 := strings.Split(elve.second, "-")

		g1L, _ := strconv.ParseInt(sections1[0], 10, 64)
		g1H, _ := strconv.ParseInt(sections1[1], 10, 64)
		g2L, _ := strconv.ParseInt(sections2[0], 10, 64)
		g2H, _ := strconv.ParseInt(sections2[1], 10, 64)

		if g1L >= g2L && g1H <= g2H {
			sum += 1
		} else if g2L >= g1L && g2H <= g1H {
			sum += 1
		} else if g1H >= g2L && g2H >= g1H{
			sum += 1
		}else if g2H >= g1L && g1H >= g2H{
            sum += 1
        }
	}

	fmt.Println(sum)

}
func P1() {

	val, _ := os.ReadFile("./part4/input.txt")

	stringVal := string(val)

	pairs := strings.Split(stringVal, "\n")

	var elves []Pair
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		groups := strings.Split(pair, ",")
		elves = append(elves, Pair{first: groups[0], second: groups[1]})
	}

	sum := 0
	for _, elve := range elves {
		sections1 := strings.Split(elve.first, "-")
		sections2 := strings.Split(elve.second, "-")

		g1L, _ := strconv.ParseInt(sections1[0], 10, 64)
		g1H, _ := strconv.ParseInt(sections1[1], 10, 64)
		g2L, _ := strconv.ParseInt(sections2[0], 10, 64)
		g2H, _ := strconv.ParseInt(sections2[1], 10, 64)

		if g1L >= g2L && g1H <= g2H {
			sum += 1
		} else if g2L >= g1L && g2H <= g1H {
			sum += 1
		}

	}

	fmt.Println(sum)

}

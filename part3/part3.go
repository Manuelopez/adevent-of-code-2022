package part3

import (
	"fmt"
	"os"
	"strings"
)

func P1() {

	val, _ := os.ReadFile("./part3/input.txt")
	stringVal := string(val)

	sacks := strings.Split(stringVal, "\n")

	sum := 0

	for _, sack := range sacks {

		sbyte := []byte(sack)
		s1 := sbyte[0 : len(sbyte)/2]
		s2 := sbyte[len(sbyte)/2 : len(sbyte)]

		var item byte
		for _, item1 := range s1 {
			for _, item2 := range s2 {
				if item1 == item2 {
					item = item1
					break
				}
			}
			if item > 0 {
				break
			}
		}

		if item != 0 {

			sum += int(getValue(item))
		}
	}
	fmt.Println(sum)
}

func P2() {
	val, _ := os.ReadFile("./part3/input.txt")
	stringVal := string(val)

	sacks := strings.Split(stringVal, "\n")

	sum := 0
	for i := 0; i+2 < len(sacks); i += 3 {
		g1 := []byte(sacks[i])
		g2 := []byte(sacks[i+1])
		g3 := []byte(sacks[i+2])

		feq := make(map[byte]int)

		for _, c := range g1 {
            feq[c] = 1
		}
		for _, c := range g2 {
            if val, _ := feq[c]; val == 1{
                feq[c] =2
            }
		}
		for _, c := range g3 {
            if val, _ := feq[c]; val ==2{
                feq[c]= 3
            }
		}

		for key, val := range feq {
			if key != 0 {
				if val == 3 {
					sum += int(getValue(key))
				}
			}
		}

	}
	fmt.Println(sum)
}

func getValue(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 96
	} else {
		return (c - 64) + 26
	}
}

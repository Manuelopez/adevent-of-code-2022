package part5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P1() {
	val, _ := os.ReadFile("./part5/input.txt")

	stringVal := string(val)

	input := strings.Split(stringVal, "\n\n")
	stack := input[0]
    moves := input[1]

	rows := strings.Split(stack, "\n")

    m := make(map[int][]string)
    lastValue := 0
	for i := len(rows) - 1; i >= 0; i-- {
		if i == len(rows)-1 {
			for j := 0; j < len([]byte(rows[i])); j += 4 {
				a, _ := strconv.Atoi(string(rows[i][j+1]))
                m[a] = make([]string, 0)
                lastValue = a
			}
			continue
		}

		for j := 0; j < len([]byte(rows[i])); j += 4 {
            a := j / 4
            if string(rows[i][j+1]) == " "{
                continue
            }
            m[a+1] = append(m[a+1], string(rows[i][j+1]))
		}
	}

    lineMoves := strings.Split(moves, "\n")

    for _, move := range lineMoves{
        words := strings.Split(move, " ")
        if(len(words) < 5){
            continue
        }
        amountToTake, _ := strconv.Atoi(words[1])
        fromStack, _ := strconv.Atoi(words[3])
        placeInStack, _ := strconv.Atoi(words[5])

        for i := 0; i < amountToTake; i++{
            value := m[fromStack][len(m[fromStack]) -1]
            m[fromStack] = m[fromStack][0:len(m[fromStack]) -1]
            m[placeInStack] = append(m[placeInStack], value)
        }
    }

    s := ""
    for i := 0; i < lastValue; i++{
        s += m[i+1][len(m[i+1]) -1]
    }
    fmt.Println(s)
}

func P2() {
	val, _ := os.ReadFile("./part5/input.txt")

	stringVal := string(val)

	input := strings.Split(stringVal, "\n\n")
	stack := input[0]
    moves := input[1]

	rows := strings.Split(stack, "\n")

    m := make(map[int][]string)
    lastValue := 0
	for i := len(rows) - 1; i >= 0; i-- {
		if i == len(rows)-1 {
			for j := 0; j < len([]byte(rows[i])); j += 4 {
				a, _ := strconv.Atoi(string(rows[i][j+1]))
                m[a] = make([]string, 0)
                lastValue = a
			}
			continue
		}

		for j := 0; j < len([]byte(rows[i])); j += 4 {
            a := j / 4
            if string(rows[i][j+1]) == " "{
                continue
            }
            m[a+1] = append(m[a+1], string(rows[i][j+1]))
		}
	}

    lineMoves := strings.Split(moves, "\n")

    for _, move := range lineMoves{
        words := strings.Split(move, " ")
        if(len(words) < 5){
            continue
        }
        amountToTake, _ := strconv.Atoi(words[1])
        fromStack, _ := strconv.Atoi(words[3])
        placeInStack, _ := strconv.Atoi(words[5])

        pushValues := make([]string, 0)
        for i := 0; i < amountToTake; i++{
            value := m[fromStack][len(m[fromStack]) -1]
            m[fromStack] = m[fromStack][0:len(m[fromStack]) -1]
            pushValues = append(pushValues, value)
        }

        for i := 0; i < amountToTake; i++{
            m[placeInStack] = append(m[placeInStack], pushValues[len(pushValues) -1 ])
            pushValues = pushValues[0:len(pushValues) - 1]
        }
    }

    s := ""
    for i := 0; i < lastValue; i++{
        s += m[i+1][len(m[i+1]) -1]
    }
    fmt.Println(s)
}

package part1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func P1(){
    val, _ := os.ReadFile("./part1/input.txt")

    stringVal := string(val);

    elves := strings.Split(stringVal, "\n");

    elvesFood := make(map[int][]string)

    currentElve := 0
    for _, cal := range elves{
        if cal == ""{
            currentElve += 1
        }else{
            elvesFood[currentElve] = append(elvesFood[currentElve], cal)
        }
    }


    max := 0
    for _, foods := range elvesFood{
        count := 0
        for _, food := range foods{

            val, _ := strconv.Atoi(food)
            count += val

        }
        if count > max {
            max = count
        }


    }
    fmt.Println(max)

}

func P2(){

    val, _ := os.ReadFile("./part1/input.txt")

    stringVal := string(val);

    elves := strings.Split(stringVal, "\n");

    elvesFood := make(map[int][]string)

    currentElve := 0
    for _, cal := range elves{
        if cal == ""{
            currentElve += 1
        }else{
            elvesFood[currentElve] = append(elvesFood[currentElve], cal)
        }
    }



    allFoods := make([]int, 0)
    for _, foods := range elvesFood{
        count := 0
        for _, food := range foods{

            val, _ := strconv.Atoi(food)
            count += val

        }
        allFoods = append(allFoods, count)
    }
    sort.Slice(allFoods, func(i, j int) bool { return allFoods[i] > allFoods[j]})
    fmt.Println(allFoods[0] + allFoods[1] + allFoods[2])
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k, rq, cq int32 = 5, 3, 4, 3
	// obstacles := [][]int32{
	// 	{3, 4},
	// 	{4, 2},
	// 	{2, 3},
	// }
	// var n, k, rq, cq int32 = 5, 0, 4, 3
	obstacles := [][]int32{}
	file, err := os.Open("input/queen-attack-2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var i int32 = -2
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		t1, _ := strconv.ParseInt(strings.Split(text, " ")[0], 10, 64)
		t2, _ := strconv.ParseInt(strings.Split(text, " ")[1], 10, 64)
		if i == -2 {
			// fmt.Println(text[0])
			// fmt.Println(text[1])
			n = int32(t1)
			k = int32(t2)
			i++
		} else if i == -1 {
			rq = int32(t1)
			cq = int32(t2)
			i++
		} else {
			obstacles = append(obstacles, []int32{int32(t1), int32(t2)})
		}
	}

	// fmt.Println(n, k, rq, cq)
	// fmt.Println(len(obstacles))
	// k = 0
	fmt.Println(queensAttack(n, k, rq, cq, obstacles))
}

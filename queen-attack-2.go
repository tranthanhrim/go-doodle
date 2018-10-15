package main

import (
	"fmt"
	"math"
)

var obsUp, obsDown, obsLeft, obsRight, obsUpLeft, obsUpRight, obsDownLeft, obsDownRight int32

func checkObstructed(n int32, rq int32, cq int32, r int32, c int32) int32 {
	fmt.Println(r, c, r-rq, c-cq)
	if r == rq {
		if cq > c && c > obsLeft {
			temp := obsLeft
			obsLeft = c
			return c - temp
		} else if n-c+1 > obsRight {
			temp := obsRight
			obsRight = n - c + 1
			return n - c + 1 - temp
		}
		return 0
	}

	if c == cq {
		if rq > r && r > obsUp {
			temp := obsUp
			obsUp = r
			return r - temp
		} else if n-r+1 > obsDown {
			temp := obsDown
			obsDown = n - r + 1
			return n - r + 1 - temp
		}
		return 0
	}

	// {2, 2} {1, 1}
	if rq-r == cq-c && rq-r > 0 {
		var offset int32
		if r < c {
			offset = r
		} else {
			offset = c
		}
		fmt.Println(offset, obsDownLeft)
		if offset > obsDownLeft {
			temp := obsDownLeft
			obsDownLeft = offset
			return offset - temp
		}
		return 0
	}

	// {2, 2} {3, 3}
	if rq-r == cq-c && rq-r < 0 {
		var offset int32
		if r > c {
			offset = n - r + 1
		} else {
			offset = n - c + 1
		}
		fmt.Println(offset)
		if offset > obsUpRight {
			temp := obsUpRight
			obsUpRight = offset
			return offset - temp
		}
		return 0
	}

	// {2, 2} {1, 3}
	if math.Abs(float64(rq-r)) == math.Abs(float64(cq-c)) && rq > r {
		var offset int32
		if r < n-c+1 {
			offset = r
		} else {
			offset = n - c + 1
		}
		if offset > obsDownRight {
			temp := obsDownRight
			obsDownRight = offset
			return offset - temp
		}
		return 0
	}

	// {2, 2} {3, 1}
	if math.Abs(float64(rq-r)) == math.Abs(float64(cq-c)) && rq < r {
		var offset int32
		if c < n-r+1 {
			offset = c
		} else {
			offset = n - r + 1
		}
		fmt.Println("obsUpLeft", offset, obsUpLeft)
		if offset > obsUpLeft {
			temp := obsUpLeft
			obsUpLeft = offset
			return offset - temp
		}
		return 0
	}
	return 0
}

func calcAvailPts(n int32, r int32, c int32) int32 {
	var top int32 = n - r
	var bot int32 = r - 1
	var left int32 = c - 1
	var right int32 = n - c
	var result int32 = (n - 1) * 2
	if bot > 0 {
		if bot < left {
			result += bot
		} else {
			result += left
		}
		if bot < right {
			result += bot
		} else {
			result += right
		}
	}
	if top > 0 {
		if top < left {
			result += top
		} else {
			result += left
		}
		if top < right {
			result += top
		} else {
			result += right
		}
	}
	return result
}

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, rq int32, cq int32, obstacles [][]int32) int32 {
	var availPts int32 = calcAvailPts(n, rq, cq)
	var i int32

	for ; i < k; i++ {
		var numObstrcuted int32 = checkObstructed(n, rq, cq, obstacles[i][0], obstacles[i][1])
		// fmt.Println("numObstrcuted", obstacles[i][0], obstacles[i][1], numObstrcuted)
		fmt.Println(numObstrcuted)
		fmt.Println(obsUp, obsDown, obsLeft, obsRight, obsUpLeft, obsUpRight, obsDownLeft, obsDownRight)
		fmt.Println("")
		availPts = availPts - numObstrcuted
	}
	return availPts
}

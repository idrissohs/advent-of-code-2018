package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns number of squares already visited
func addPoints(lstart, tstart, wDimension, hDimension int, cur map[int]map[int]struct{}, visited map[int]map[int]struct{}) int {
	ans := 0
	for i := tstart; i < tstart+hDimension; i++ {
		for j := lstart; j < lstart+wDimension; j++ {
			if _, ok := cur[i]; !ok {
				cur[i] = map[int]struct{}{}
			}
			if _, ok := visited[i]; !ok {
				visited[i] = map[int]struct{}{}
			}
			if _, ok := cur[i][j]; ok {
				if _, ok := visited[i][j]; !ok {
					ans++
					visited[i][j] = struct{}{}
				}
			} else {
				cur[i][j] = struct{}{}
			}
		}
	}
	return ans
}

// checkPoints checks if the square has been visited yet
func checkPoints(lstart, tstart, wDimension, hDimension int, visited map[int]map[int]struct{}) bool {
	for i := tstart; i < tstart+hDimension; i++ {
		for j := lstart; j < lstart+wDimension; j++ {
			if _, ok := visited[i][j]; ok {
				return false
			}
		}
	}
	return true
}
func main() {
	res := 0
	cur := map[int]map[int]struct{}{}
	visited := map[int]map[int]struct{}{}
	fh, err := os.Open("./day3Puzzle.txt")
	if err != nil {
		fmt.Println("Error")
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		start := strings.Split(line[2], ",")
		start[1] = start[1][:len(start[1])-1]
		dimension := strings.Split(line[3], "x")
		leftStart, _ := strconv.Atoi(start[0])
		topStart, _ := strconv.Atoi(start[1])
		wDimension, _ := strconv.Atoi(dimension[0])
		hDimension, _ := strconv.Atoi(dimension[1])
		//fmt.Printf("%d, %d, %d, %d", leftStart, topStart, wDimension, hDimension)
		res += addPoints(leftStart, topStart, wDimension, hDimension, cur, visited)
	}
	fmt.Println(res)

	// part 2
	fh.Seek(0, 0)
	scanner = bufio.NewScanner(fh)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		start := strings.Split(line[2], ",")
		start[1] = start[1][:len(start[1])-1]
		dimension := strings.Split(line[3], "x")
		leftStart, _ := strconv.Atoi(start[0])
		topStart, _ := strconv.Atoi(start[1])
		wDimension, _ := strconv.Atoi(dimension[0])
		hDimension, _ := strconv.Atoi(dimension[1])
		if checkPoints(leftStart, topStart, wDimension, hDimension, visited) {
			fmt.Println(line[0])
		}
	}
}

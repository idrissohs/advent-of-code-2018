package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type logLine struct {
	date    time.Time
	logType string
}

type logs []logLine

func (p logs) Len() int {
	return len(p)
}

func (p logs) Less(i, j int) bool {
	return p[i].date.Before(p[j].date)
}

func (p logs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	format := "2006-01-02 15:04"
	curGuard := ""
	isEnd := false
	guardTime := map[string]map[int]int{}
	startMin, endMin := 0, 0
	var logLines logs
	fh, err := os.Open("./day4Puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		curTime, err := time.Parse(format, line[1:17])
		if err != nil {
			fmt.Println(err)
		}
		logLines = append(logLines, logLine{curTime, line[18:]})
	}
	sort.Sort(logLines)
	for _, v := range logLines {
		if strings.Contains(v.logType, "Guard") {
			curLine := strings.Fields(v.logType)
			curGuard = curLine[1]
		} else if strings.Contains(v.logType, "falls") {
			startMin = v.date.Minute()
		} else {
			endMin = v.date.Minute()
			isEnd = true
		}
		if isEnd {
			for i := startMin; i <= endMin; i++ {
				if _, ok := guardTime[curGuard]; !ok {
					guardTime[curGuard] = map[int]int{}
				}
				guardTime[curGuard][i]++
			}
			isEnd = false
		}
	}
	maxMin := 0
	maxGuard := 0
	guardID := ""
	curMax := 0
	curMaxMinID := 0
	maxMinID := 0
	for i, v := range guardTime {
		maxMin = 0
		curMax = 0
		curMaxMinID = 0
		for j, v2 := range v {
			if v2 > maxMin {
				maxMin = v2
				curMaxMinID = j
			}
			curMax += v2
		}
		if curMax > maxGuard {
			maxGuard = curMax
			guardID = i
			maxMinID = curMaxMinID
		}
	}
	fmt.Println(guardID)
	fmt.Println(maxMinID)

	// part 2
	guardID2 := ""
	maxMin2 := 0
	min2 := 0
	for i, v := range guardTime {
		for j, v2 := range v {
			if v2 > maxMin2 {
				maxMin2 = v2
				guardID2 = i
				min2 = j
			}
		}
	}
	fmt.Println(guardID2)
	fmt.Println(min2)
}

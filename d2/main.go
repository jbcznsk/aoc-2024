package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

    f, err := os.Open(os.Args[1])	
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    fs := bufio.NewScanner(f)
    fs.Split(bufio.ScanLines)

    var safeCount int
    safeCount = 0

    for fs.Scan() {
        report := strings.Fields(fs.Text())
        fmt.Print(report)
        if isSafe(report) {
            fmt.Println(" : SAFE")
            safeCount++
        } else {
            fmt.Println(" : UNSAFE")
        }
    }

    fmt.Println(safeCount)
}

func isSafe(level []string) bool {
    
    if isIncreasing(level, 1, 3) || isDecreasing(level, 1, 3) {
        return true
    }

    return false

}

func isIncreasing(report []string, minf int, maxf int) bool {

    before, _ := strconv.Atoi(report[0])
    before = before - minf

    for _, level := range report {
        levelInt, _ := strconv.Atoi(level)
        diff := levelInt - before
        if (diff < minf || diff > maxf) {
            return false
        }
        before = levelInt
    }

    return true
}

func isDecreasing(report []string, minf int, maxf int) bool {
    before, _ := strconv.Atoi(report[0])
    before = before + minf

    for _, level := range report {
        levelInt, _ := strconv.Atoi(level)
        diff := before - levelInt 
        if (diff < minf || diff > maxf) {
            return false
        }
        before = levelInt
    }
    return true
}
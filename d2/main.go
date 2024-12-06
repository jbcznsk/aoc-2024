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
        fmt.Println(report)
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

    isInc := true

    for _, level := range report {
        levelInt, _ := strconv.Atoi(level)
        diff := levelInt - before
        if (diff < minf || diff > maxf) {
            isInc =  false
        }
        before = levelInt
    }

    if !isInc {
        fmt.Println("Trying subreports")
        isValidSubReport := true
        validSubReportCount := 0
        for i := 0; i < len(report); i++ {
            isValidSubReport = true
            newRep := removeElement(report, i)
            before, _ := strconv.Atoi(newRep[0])
            before = before - minf
            fmt.Println(newRep)
            fmt.Println(isValidSubReport)
            for _, level := range newRep {
                levelInt, _ := strconv.Atoi(level)
                diff := levelInt - before
                if (diff < minf || diff > maxf) {
                    fmt.Println(levelInt)
                     fmt.Println(before)
                    isValidSubReport = false
                }
                before = levelInt
            }
            if isValidSubReport {
                fmt.Println("OK")
                validSubReportCount++
            } else {
                fmt.Println("NOK")
            }      
        }
        if validSubReportCount > 0 {
            return true
        } else {
            return false
        }
    } else {
        return true
    }

}

func isDecreasing(report []string, minf int, maxf int) bool {
    before, _ := strconv.Atoi(report[0])
    before = before + minf

    isDec := true

    for _, level := range report {
        levelInt, _ := strconv.Atoi(level)
        diff := before - levelInt 
        if (diff < minf || diff > maxf) {
            isDec = false
        }
        before = levelInt
    }

    if !isDec {
        isValidSubReport := true
        validSubReportCount := 0
        for i := 0; i < len(report); i++ {
            isValidSubReport = true
            newRep := removeElement(report, i)
            before, _ := strconv.Atoi(newRep[0])
            before = before + minf
            for _, level := range newRep {
                levelInt, _ := strconv.Atoi(level)
                diff := before - levelInt 
                if (diff < minf || diff > maxf) {
                    isValidSubReport = false
                }
                before = levelInt
            }
            if isValidSubReport {
                validSubReportCount++
            }
        }
        if validSubReportCount > 0 {
            return true
        } else {
            return false
        }
    } else {
        return true
    }

}

func removeElement(slice []string, pos int) []string {
    newSlice := make([]string, 0)
    newSlice = append(newSlice, slice[:pos]...)
    newSlice = append(newSlice, slice[pos+1:]...)
    return newSlice
}
package main

import (
	"bufio"
	"fmt"
	"strconv"
	// "io"
	// "log"
	"os"
	"strings"
	"math"
	"sort"
)

func main() {
	

	var v1 []int
	var v2 []int

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	soma := 0
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		partes := strings.Fields(fileScanner.Text())
		n1, err1 := strconv.Atoi(partes[0])
		n2, err2 := strconv.Atoi(partes[1])

		if err1 != nil || err2 != nil {
			fmt.Println(err1)
			fmt.Println(err2)
			return
		}

		v1 = append(v1, n1)
		v2 = append(v2, n2)
	}

	sort.Ints(v1)
	sort.Ints(v2)

	for i:= 0; i < len(v1); i++ {
		soma = soma + int(math.Abs(float64(v1[i]) - float64(v2[i])))
	}



	fmt.Println(soma)
}

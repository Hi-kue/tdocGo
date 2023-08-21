package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mealCost, err := strconv.ParseFloat(readLine(reader), 64)
	checkErr(err)

	tipPercent, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkErr(err)

	taxPercent, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkErr(err)

	meals(mealCost, tipPercent, taxPercent)
}

func meals(mealCost float64, tipPercent int64, taxPercent int64) {
	tip, tax := mealCost*(float64(tipPercent)/100), mealCost*(float64(taxPercent)/100)
	totalCost := math.Round(mealCost + tip + tax)

	fmt.Println(totalCost)
}

func readLine(reader *bufio.Reader) string {
	line, _, err := reader.ReadLine()
	checkErr(err)

	return strings.TrimRight(string(line), "\r\n")
}

func checkErr(err error) {
	if err == io.EOF {
		return
	}

	if err != nil {
		panic(err)
	}
}

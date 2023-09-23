package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// CB=Cheese burger
const (
	inputFileName  = "cheeseburger_corollary_2_input.txt"
	outputFileName = "output.txt"
	bunsPerSOrDCB  = 2
	pattiesPerDCB  = 2
	cheesePerDCB   = 2
	NO             = "NO"
	YES            = "YES"
)

func main() {
	tCases, sc, closeFile := parseInput(inputFileName)
	defer closeFile()
	outFile, err := os.OpenFile(outputFileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	for i := 1; i <= tCases; i++ {
		sc.Scan()
		inputs := strings.Fields(sc.Text())
		perSCBCost, _ := strconv.Atoi(inputs[0])
		perDCBCost, _ := strconv.Atoi(inputs[1])
		totalBudget, _ := strconv.Atoi(inputs[2])
		if totalBudget < perDCBCost && totalBudget < perSCBCost {
			outFile.WriteString(fmt.Sprintf("Case #%d: %d\n", i, 0))
			continue
		}
		sCBCount := totalBudget / perSCBCost
		dCBCount := totalBudget / perDCBCost

		if totalBudget == 5 {
			sCBCount = 1
			dCBCount = 1
		} else if sCBCount >= 2*dCBCount {
			rem := totalBudget % perSCBCost
			dCBCount = rem / perDCBCost
		} else {
			rem := totalBudget % perDCBCost
			sCBCount = rem / perSCBCost
		}

		totalBuns := (sCBCount * bunsPerSOrDCB) + (dCBCount * bunsPerSOrDCB)
		totalPatties := (sCBCount * 1) + (dCBCount * pattiesPerDCB)
		totalCheese := (sCBCount * 1) + (dCBCount * cheesePerDCB)

		kMaxDecker := calculateDeckerKValue(totalBuns, totalPatties, totalCheese)

		// fmt.Printf("Case #%d: %d\n", i, kMaxDecker)
		outFile.WriteString(fmt.Sprintf("Case #%d: %d\n", i, kMaxDecker))
	}

}

func parseInput(fileName string) (int, *bufio.Scanner, func() error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	sc.Scan()
	testCases, _ := strconv.Atoi(sc.Text())
	return testCases, sc, file.Close
}

func calculateDeckerKValue(buns, patties, cheese int) int {
	//if buns are more than patties
	if buns > patties && buns > cheese && patties == cheese {
		return patties
	}
	//all values will be equal at the very least
	k := patties - 1
	if k < 0 {
		k = 0
	}
	return k
}

package main

import (
    "bufio"
    "log"
    "os"
		"fmt"
		"strings"
)

type Sheet struct {
	rows    []string
  index   map[string]bool
}

func main() {
	var userRefs []string = make([]string, 1)
	var settlementIds []string = make([]string, 1)
  var inputFileName = "./duplicates.csv"
	userRefs, settlementIds = readFile(inputFileName)

  fmt.Println("Found ", (len(userRefs) - 1),  " records in ", inputFileName)

	var sheets []Sheet = createSheets(50)

	for i := 1; i < len(userRefs); i++ {
    fmt.Print(".")
		addToASheet(sheets, userRefs[i], settlementIds[i])
	}

  fmt.Println("")
  for i := 0; i < len(sheets); i++ {
    if (len(sheets[i].rows) == 0) {
      break
    }

    writeCsv(sheets[i], i, userRefs[0], settlementIds[0])
  }
}

func writeCsv(sheet Sheet, index int, column1Heading string, column2Heading string) {
  var filename string = fmt.Sprintf("processed_%d.csv", index + 1)
  file, err := os.Create(filename)

  if err != nil {
    log.Fatalf("failed creating file: %s", err)
  }

  datawriter := bufio.NewWriter(file)
  datawriter.WriteString(column1Heading + "," + column2Heading + "\n")

  for _, row := range sheet.rows {
    _, _ = datawriter.WriteString(row + "\n")
  }

  fmt.Println("Finished writing ", filename,  " with", len(sheet.rows), " record(s)")

  datawriter.Flush()
  file.Close()
}

func addToASheet(sheets []Sheet, userRef string, settlementId string) {
  for i := 0; i < len(sheets); i++ {
		if sheets[i].index[userRef] == false {
        sheets[i].index[userRef] = true
        sheets[i].rows = append(sheets[i].rows, userRef + "," + settlementId)
        break
    }
	}
}

func readFile(path string)([]string, []string) {
	file, err := os.Open("./duplicates.csv")

	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	var userRefs []string = make([]string, 0)
	var settlementIds []string = make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row string = scanner.Text()

    if (len(row) > 0) {
      var columns []string = strings.Split(row, ",")
      userRefs = append(userRefs, columns[0])
  		settlementIds = append(settlementIds, columns[1])
    }
	}

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}

	return userRefs, settlementIds
}

func createSheets(numberOfSheets int)([]Sheet) {
  var sheets []Sheet = make([]Sheet, numberOfSheets)

	for i := 0; i < numberOfSheets; i++ {
    sheets[i] = Sheet {index: make(map[string]bool), rows: make([]string, 0)}
	}

  return sheets
}

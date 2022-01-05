package main

import (
	"fmt"
	"os"

	"github.com/exactlylabs/parquet2csv/pkg/convert"
)

func printUsage() {
	fmt.Println("Usage: parquet2csv <parquet_file>")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	parquetFilePath := os.Args[1]

	convert.Parquet2CSV(parquetFilePath)
}

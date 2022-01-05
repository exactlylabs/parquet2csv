package convert

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	goparquet "github.com/fraugster/parquet-go"
	"github.com/fraugster/parquet-go/parquet"
)

// Take a parquet file and convert it to a CSV file
func Parquet2CSV(parquetFilePath string) {
	fl, err := os.Open(parquetFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not open the file: %q\n", err)
		return
	}
	defer fl.Close()

	reader, err := goparquet.NewFileReader(fl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read the parquet header: %q\n", err)
		return
	}

	schema := reader.GetSchemaDefinition()
	for _, child := range schema.RootColumn.Children {
		if child.Children != nil {
			fmt.Fprintf(os.Stderr, "unsupported nested schema: %q\n", child.SchemaElement.Name)
			return
		}
	}

	names := make([]string, len(schema.RootColumn.Children))
	for i, col := range schema.RootColumn.Children {
		names[i] = col.SchemaElement.Name
	}

	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	w.Write(names)

	for {
		data, err := reader.NextRow()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading row data failed with error: %q, skipping row\n", err)
			continue
		}

		row := make([]string, len(names))
		for i, v := range names {
			t := schema.RootColumn.Children[i].SchemaElement.Type
			if t != nil && *t == parquet.Type_BYTE_ARRAY {
				row[i] = string(data[v].([]byte))
			} else {
				row[i] = fmt.Sprint(data[v])
			}
		}

		w.Write(row)
	}
}

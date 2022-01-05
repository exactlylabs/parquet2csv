### Parquet2CSV

Usage:

    parquet2csv <PATH_TO_PARQUET_FILE>

This will output a CSV file to `STDOUT`.

Currently, this only works with files without nested fields.

It also assumes all `BYTE_ARRAY` columns are to be writen as a `UTF8` string.

To install locally:

    go install github.com/exactlylabs/parquet2csv/cmd/parquet2csv@latest


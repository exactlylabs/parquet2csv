### Parquet2CSV

Parquet2CSV is a simple tool to convert a Parquet file to a CSV written in Go/ Golang

Usage:

    parquet2csv <PATH_TO_PARQUET_FILE>

This will output a CSV file to `STDOUT`.

Currently, this only works with files without nested fields.

It also assumes all `BYTE_ARRAY` columns are to be writen as a `UTF8` string.

To install locally:

    go install github.com/exactlylabs/parquet2csv/cmd/parquet2csv@latest

or download a pre-built binary from https://github.com/exactlylabs/parquet2csv/releases


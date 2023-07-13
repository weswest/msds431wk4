package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func run(inputFile, outputFile string, N int) error {
	outfile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outfile.Close()

	for i := 0; i < N; i++ {
		file, err := os.Open(inputFile)
		if err != nil {
			return err
		}

		reader := csv.NewReader(file)
		header, err := reader.Read()
		if err != nil {
			return err
		}

		columns := make(map[string][]float64)
		for _, field := range header {
			columns[field] = []float64{}
		}

		reader.ReuseRecord = true // This saves on memory allocation

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			for j, val := range record {
				value, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return err
				}
				columns[header[j]] = append(columns[header[j]], value)
			}
		}

		for _, field := range header {
			sort.Float64s(columns[field])
		}

		file.Close()

		fmt.Fprint(outfile, strings.Join(header, "\t")+"\n")

		metrics := []string{"count", "mean", "std", "min", "25%", "50%", "75%", "max"}
		for _, metric := range metrics {
			fmt.Fprint(outfile, metric)
			for _, field := range header {
				switch metric {
				case "count":
					fmt.Fprintf(outfile, "\t%f", float64(len(columns[field])))
				case "mean":
					fmt.Fprintf(outfile, "\t%f", stat.Mean(columns[field], nil))
				case "std":
					fmt.Fprintf(outfile, "\t%f", stat.StdDev(columns[field], nil))
				case "min":
					fmt.Fprintf(outfile, "\t%f", floats.Min(columns[field]))
				case "25%":
					fmt.Fprintf(outfile, "\t%f", stat.Quantile(0.25, stat.Empirical, columns[field], nil))
				case "50%":
					fmt.Fprintf(outfile, "\t%f", stat.Quantile(0.5, stat.Empirical, columns[field], nil))
				case "75%":
					fmt.Fprintf(outfile, "\t%f", stat.Quantile(0.75, stat.Empirical, columns[field], nil))
				case "max":
					fmt.Fprintf(outfile, "\t%f", floats.Max(columns[field]))
				}
			}
			fmt.Fprintln(outfile)
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the input file name as a command-line argument.")
	}
	inputFile := os.Args[1]

	if err := run(inputFile, "housesOutputGo.txt", 100); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

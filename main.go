package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func main() {
	N := 100

	// Create the file that will be populated with the output
	outfile, err := os.Create("housesOutputGo.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outfile.Close()

	for i := 0; i < N; i++ {

		// Open the file
		file, err := os.Open("housesInput.csv")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Capture the first row of the csv.  This is used to grab the field names
		reader := csv.NewReader(file)
		// reader.ReuseRecord = true     // This saves on memory allocation
		header, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create a map of the columns.  This will be used to store the data
		columns := make(map[string][]float64)
		for _, field := range header {
			columns[field] = []float64{}
		}

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for j, val := range record {
				value, err := strconv.ParseFloat(val, 64)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
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
}

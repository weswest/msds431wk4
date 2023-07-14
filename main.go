// Package main provides the implementation for running the CSV processing program.
package main

import (
	"encoding/csv" // Importing the encoding/csv package for reading and writing CSV files.
	"fmt"          // Importing the fmt package for formatted I/O functions.
	"io"           // Importing the io package for basic interfaces to I/O primitives.
	"log"          // Importing the log package for simple logging.
	"os"           // Importing the os package for interacting with the operating system.
	"sort"         // Importing the sort package for sorting primitives.
	"strconv"      // Importing the strconv package for conversions to and from string representations of basic data types.
	"strings"      // Importing the strings package for manipulation of string objects.

	"gonum.org/v1/gonum/floats" // Importing the gonum/floats package for operations on slices of float64s.
	"gonum.org/v1/gonum/stat"   // Importing the gonum/stat package for statistical computations.
)

// run reads the CSV data from inputFile, computes statistics for each column,
// and writes the results to outputFile. It repeats this process N times.
func run(inputFile, outputFile string, N int) error {
	// Create the output file.
	outfile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	// Ensure the output file is closed after the function returns.
	defer outfile.Close()

	// Repeat the process N times.
	for i := 0; i < N; i++ {
		// Open the input file.
		file, err := os.Open(inputFile)
		if err != nil {
			return err
		}

		// Create a new CSV reader for the input file.
		reader := csv.NewReader(file)
		// Read the header row from the CSV file.
		header, err := reader.Read()
		if err != nil {
			return err
		}

		// Initialize a map to hold the data for each column.
		columns := make(map[string][]float64)
		for _, field := range header {
			columns[field] = []float64{}
		}

		// Enable reusing of the record buffer to save memory.
		reader.ReuseRecord = true

		// Read the data rows from the CSV file.
		for {
			record, err := reader.Read()
			// Break the loop when we reach the end of the file.
			if err == io.EOF {
				break
			}
			// Return any other error.
			if err != nil {
				return err
			}

			// Parse the data and add it to the map.
			for j, val := range record {
				value, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return err
				}
				columns[header[j]] = append(columns[header[j]], value)
			}
		}

		// Sort the data in each column.
		for _, field := range header {
			sort.Float64s(columns[field])
		}

		// Close the input file.
		file.Close()

		// Write the header row to the output file.
		fmt.Fprint(outfile, strings.Join(header, "\t")+"\n")

		// Define the metrics we're going to compute.
		metrics := []string{"count", "mean", "std", "min", "25%", "50%", "75%", "max"}

		// Compute and write the metrics for each column.
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

	// Return nil to indicate success.
	return nil
}

// main is the entry point of the program.
func main() {
	// Check if the user has provided an input file name.
	if len(os.Args) < 2 {
		log.Fatal("Please provide the input file name as a command-line argument.")
	}
	// Get the input file name from the command-line argument.
	inputFile := os.Args[1]

	// Run the program and handle any error.
	if err := run(inputFile, "housesOutputGo.txt", 100); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

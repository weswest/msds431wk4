# Overview

This repo is for the Northwestern Masters in Data Science [MSDS-431 Learning Data Engineering with Go](https://msdsgo.netlify.app/data-engineering-with-go/) week 4 assignment.  See details of the assignment at the end of this Readme.

This program essentially replicates the .describe() method in python, taking a csv of numbers and outputting summary statistics (count, mean, stdev, min, max, and quartile values) for each numeric column.  The file name is read in as an input at time of execution, so in the terminal enter 

```
./msds431wk4 <FILENAME> 
```

Note that this is a toy project, with the output file (housesOutputGo.txt) as a fixed value.  If you try to read in a csv with non-text data it will just return an error.

Finally - an intent with this assignment was to test the speed of the program versus R and python, so in addition to the hardcoded nature of its application, the program writes the output of the .describe function 100 times to the text file.

# Assignment Details

This program is structured around mimicking the python .describe() function which provides a set of summary statistics for each column of numeric data in a dataframe.  With this as the premise, the gonum/stat package was used as the source for statistical calculations.  This is a very popular go package with nearly 7k stars, and more than 6000 commits, including some in the last week.

Most of the program is, as expected, oriented around getting the data read in and structured correctly to execute the various statistical commands.  The output is relatively hard-coded as a series of tab-delimited strings - this is an area for potential refinement in the future.

## Comparison of Results against R and Python

Here are the results of the program from each of the three systems tested:

### R's summary(houses)

| value          | income          | age           | rooms         | bedrooms       | pop           | hh             |
| -------------- | --------------- | ------------- | ------------- | -------------- | ------------- | -------------- |
| Min. : 14999   | Min. : 0.4999   | Min. : 1.00   | Min. : 2      | Min. : 1.0     | Min. : 3      | Min. : 1.0     |
| 1st Qu.:119600 | 1st Qu.: 2.5634 | 1st Qu.:18.00 | 1st Qu.: 1448 | 1st Qu.: 295.0 | 1st Qu.: 787  | 1st Qu.: 280.0 |
| Median :179700 | Median : 3.5348 | Median :29.00 | Median : 2127 | Median : 435.0 | Median : 1166 | Median : 409.0 |
| Mean :206856   | Mean : 3.8707   | Mean :28.64   | Mean : 2636   | Mean : 537.9   | Mean : 1425   | Mean : 499.5   |
| 3rd Qu.:264725 | 3rd Qu.: 4.7432 | 3rd Qu.:37.00 | 3rd Qu.: 3148 | 3rd Qu.: 647.0 | 3rd Qu.: 1725 | 3rd Qu.: 605.0 |
| Max. :500001   | Max. :15.0001   | Max. :52.00   | Max. :39320   | Max. :6445.0   | Max. :35682   | Max. :6082.0   |

### Python's houses.describe()

|       | value       | income   | age       | rooms       | bedrooms   | pop         | hh         |
| ----- | ----------- | -------- | --------- | ----------- | ---------- | ----------- | ---------- |
| count | 20640       | 20640    | 20640     | 20640       | 20640      | 20640       | 20640      |
| mean  | 206855.8169 | 3.870671 | 28.639486 | 2635.763081 | 537.898014 | 1425.476744 | 499.53968  |
| std   | 115395.6159 | 1.899822 | 12.585558 | 2181.615252 | 421.247906 | 1132.462122 | 382.329753 |
| min   | 14999       | 0.4999   | 1         | 2           | 1          | 3           | 1          |
| 25%   | 119600      | 2.5634   | 18        | 1447.75     | 295        | 787         | 280        |
| 50%   | 179700      | 3.5348   | 29        | 2127        | 435        | 1166        | 409        |
| 75%   | 264725      | 4.74325  | 37        | 3148        | 647        | 1725        | 605        |
| max   | 500001      | 15.0001  | 52        | 39320       | 6445       | 35682       | 6082       |

### Go's homebrew msds431wk4 housesInput.csv

|       | value       | income   | age       | rooms       | bedrooms   | pop         | hh         |
| ----- | ----------- | -------- | --------- | ----------- | ---------- | ----------- | ---------- |
| count | 20640       | 20640    | 20640     | 20640       | 20640      | 20640       | 20640      |
| mean  | 206855.8169 | 3.870671 | 28.639486 | 2635.763081 | 537.898014 | 1425.476744 | 499.53968  |
| std   | 115395.6159 | 1.899822 | 12.585558 | 2181.615252 | 421.247906 | 1132.462122 | 382.329753 |
| min   | 14999       | 0.4999   | 1         | 2           | 1          | 3           | 1          |
| 25%   | 119600      | 2.5625   | 18        | 1447        | 295        | 787         | 280        |
| 50%   | 179700      | 3.5347   | 29        | 2127        | 435        | 1166        | 409        |
| 75%   | 264700      | 4.7431   | 37        | 3148        | 647        | 1725        | 605        |
| max   | 500001      | 15.0001  | 52        | 39320       | 6445       | 35682       | 6082       |

### Go MINUS Python

|       | value | income    | age | rooms  | bedrooms | pop | hh |
| ----- | ----- | --------- | --- | ------ | -------- | --- | -- |
| count | 0     | 0         | 0   | 0      | 0        | 0   | 0  |
| mean  | 0     | 0         | 0   | 0      | 0        | 0   | 0  |
| std   | 0     | 0         | 0   | 0      | 0        | 0   | 0  |
| min   | 0     | 0         | 0   | 0      | 0        | 0   | 0  |
| 25%   | 0     | \-0.0009  | 0   | \-0.75 | 0        | 0   | 0  |
| 50%   | 0     | \-0.0001  | 0   | 0      | 0        | 0   | 0  |
| 75%   | \-25  | \-0.00015 | 0   | 0      | 0        | 0   | 0  |
| max   | 0     | 0         | 0   | 0      | 0        | 0   | 0  |

The one area where results differ is in the quantile calculations, although these are still close enough for our purposes.  I would guess there is a minor difference in execution related to rounding when the quantile doesn't land on a specific number.

## Execution Speeds

As described in the assignment, we were encouraged to use Ricardo Gerardi's *Powerful Command-Line Applications in Go: Build Fast and Maintainable Tools* as a guide for optimizing our program and to compare our go program to R and Python.

The "time <command>" structure was used to test execution speed of the python, R, and go programs.  Something to note is that if you run "time go run" then the resultant time will be the combination of the program executing and the time necessary to build the go executable.

Also of note is that there are two versions of the go program.  The first was a minimally viable product that ran correctly; the second is the beginnings of optimization to make the program run a little faster.  The main change was in managing the .csv file after it was read in: I followed Gerardi's recommendation to use reader.ReuseRecord = true as a way to reduce the memory footprint while iterating through a csv, which had a measurable reduction in execution time.

|                | real  | user  | sys   | user+sys | Comment                                  |
| -------------- | ----- | ----- | ----- | -------- | ---------------------------------------- |
| python         | 7.592 | 4.713 | 0.79  | 5.503    | Using code provided                      |
| R              | 8.697 | 8.25  | 0.362 | 8.612    | Using code provided                      |
| go run v1      | 6.241 | 5.158 | 1.01  | 6.168    | First iteration of go code               |
| go compiled v1 | 4.884 | 4.77  | 0.213 | 4.983    |                                          |
| go run v2      | 5.011 | 4.852 | 0.672 | 5.524    | v2 follows Ricardi recommendation to set |
| go compiled v2 | 4.676 | 4.432 | 0.179 | 4.611    | reader.ReuseRecord = true to save memory |

# Assignment details

## Management Problem
[It's déjà vu all over again.] Managers of a technology startup are keen on limiting the number of computer languages supported by the company. They want to see employees using Go as their primary programming language. 

As with the Week 1 assignment, we will be comparing R, Python/Pandas, and Go programs for addressing the management problem. But in this assignment, we are asked to develop a Go command-line application.

Data scientists often use R or Python/Pandas to read comma-delimited text files and compute summary statistics. These are straightforward tasks that could easily be done using modules from the Go standard library, working within a command-line interface (CLI).

## Assignment Requirements

We take the role of company data scientists, learning how to use Go on the command line.

The input data for this assignment are provided in a comma-delimited text file (housesInput.csv Download housesInput.csv) derived from a study of California housing prices (Miller 2015). There are R (runHouses.R Download runHouses.R) and Python/Pandas (runHouses.py Download runHouses.py) scripts for reading these data and computing summary statistics for each of seven input variables. Output from running these scripts should be a plain text files of summary statistics: housesOutputR.txt Download housesOutputR.txtand housesOutputPy.txt Download housesOutputPy.txtshow output from running the R and Python/Pandas scripts, respectively.

We need to enter console/terminal commands for running the R and Python/Pandas scripts on the CLI of our development platform (Windows 10/11, MacOS, or Linux). We run the scripts as batch jobs and determine CPU processing times using console/terminal commands.

When running benchmark trials, ensure (to the extent possible) that no other application programs are running. This includes web browsers and background processes not needed for the program being benchmarked.

Write a Go command-line application to perform the same operations as the R and Python programs: (1) reading the comma-delimited text file of California housing data, (2) computing summary statistics for each of the seven input variables, and (3) listing the results in a plain text file. Ensure that you compute at least three summary statistics in your Go command-line application: minimum, maximum, and mean. Note that Gerardi (2021) provides exemplary Go command-line code for reading comma-delimited text files and for improving the performance of command-line applications. He also suggests using the Cobra CLI framework when developing command-line applications. 

Using the command-line interface/window, run the Go command-line application, repeating the data file reading, summary statistics calculations, and writing to the text output file one hundred times. Determine the CPU processing time across the one hundred runs. As time permits, employ methods described in Gerardi (2021) to improve the efficiency of the Go command-line application. 

Compare the processing time of the most efficient Go command-line application with the processing times of the R and Python/Pandas scripts.

Ensure that the results from the Go application match (within rounding error) the results from the R and Python/Pandas scripts.

While working on this assignment, we should be writing idiomatic Go, using tools for code formatting, linting, and vetting as described by Mat Ryer, Jaana Dogan, and Johnny Boursiquot in this Go TimeLinks to an external site. recording from 2019:
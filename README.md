# Overview

This repo is for the Northwestern Masters in Data Science [MSDS-431 Learning Data Engineering with Go](https://msdsgo.netlify.app/data-engineering-with-go/) week 4 assignment.  See details of the assignment at the end of this Readme

# Assignment Details

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
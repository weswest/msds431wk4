package main

import (
	"os"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = run("housesInput.csv", "housesOutputGo.txt", 100)
	}
}

func TestRunValidFile(t *testing.T) {
	err := run("housesInput.csv", "housesOutputGo.txt", 100)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestRunInvalidFile(t *testing.T) {
	err := run("invalidFile.csv", "housesOutputGo.txt", 100)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}

func TestRunInvalidData(t *testing.T) {
	// Create dummy CSV file with invalid data
	file, err := os.Create("dummy.csv")
	if err != nil {
		t.Fatalf("Failed to create dummy CSV file: %v", err)
	}
	defer os.Remove("dummy.csv")

	_, err = file.WriteString("height,width,color\n3,5,red\nfour,4,blue\n1,1,green\n")
	file.Close()
	if err != nil {
		t.Fatalf("Failed to write to dummy CSV file: %v", err)
	}

	err = run("dummy.csv", "housesOutputGo.txt", 100)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}

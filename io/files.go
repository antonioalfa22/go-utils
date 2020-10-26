package io

import (
	"bufio"
	"github.com/antonioalfa22/go-utils/errors"
	"os"
)

func ReadFile(filename string) []string {
	f, err := os.Open(filename)
	errors.CheckSoft(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines,scanner.Text())
	}
	return lines
}

func WriteFile(lines []string, filename string) {
	f, err := os.Create(filename)
	errors.CheckSoft(err)
	defer f.Close()
	for _, line := range lines{
		_, err := f.WriteString(line+"\n")
		errors.CheckSoft(err)
	}
	_ = f.Sync()
}
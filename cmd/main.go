package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("no arguments given. Please add filename and (optional) line number")
	}
	lineNumber := 0
	filename := args[0]
	log.Printf("using file: %s", filename)
	if len(args) > 1 {
		if s, err := strconv.ParseInt(args[1], 10, 32); err == nil {
			lineNumber = int(s)
		} else {
			panic(fmt.Sprintf("error parsing linenumber: %v", err))
		}
		log.Printf("using line %d", lineNumber)
	}

	lines, err := readLines(filename)
	if err != nil {
		panic(fmt.Sprintf("error reading file: %v", err))
	}
	if lineNumber > len(lines) {
		panic(fmt.Sprintf(`linenumber "%d" not found in file "%s"`, lineNumber, filename))
	}
	log.Printf("starting command line: %s \r\n", lines[lineNumber])
	largs := strings.Split(lines[lineNumber], " ")

	cargs := make([]string, 0)
	cargs = append(cargs, "/k")
	cargs = append(cargs, largs...)

	cmd := exec.Cmd{
		Path: "c:\\windows\\system32\\cmd.exe",
		Args: cargs,
		SysProcAttr: &syscall.SysProcAttr{
			CreationFlags:    0x10,
			NoInheritHandles: true,
		},
	}
	err = cmd.Start()
	if err != nil {
		log.Printf("cant start program: %v", err)
	}
	time.Sleep(5000)
}

func readLines(filePath string) ([]string, error) {
	readFile, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines, nil
}

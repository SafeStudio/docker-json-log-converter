package lc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ansidev/djlc/file"
	"log"
	"os"
)

func FromFile(input string) *DockerLogConverter {
	return &DockerLogConverter{
		inputFile: input,
	}
}

type DockerLogConverter struct {
	inputFile  string
	outputFile string
}

func (c *DockerLogConverter) ToFile(output string) *DockerLogConverter {
	c.outputFile = output

	return c
}

func (c DockerLogConverter) Convert() {
	if !file.IsExists(c.inputFile) {
		fmt.Println("Input file", c.inputFile, "does not exist")
		return
	}

	err := c.initOutputFile()
	if err != nil {
		return
	}

	in, err := os.Open(c.inputFile)
	if err != nil {
		log.Fatalf("Failed to open file %s", c.inputFile)
		return
	}

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	var dockerLogRow DockerLogRow
	for scanner.Scan() {
		text := scanner.Text()
		err := json.Unmarshal([]byte(text), &dockerLogRow)

		if err != nil {
			log.Printf("ERROR %s", text)
		} else {
			log.Print(dockerLogRow.Log)
		}
	}

	in.Close()
}

func (c *DockerLogConverter) initOutputFile() error {
	logFile, err := os.OpenFile(c.outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatal("Cannot write to output file", c.outputFile)
		return err
	} else {
		log.SetOutput(logFile)
		log.SetFlags(0)
		return nil
	}
}

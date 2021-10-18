package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Mapper struct {
	matrix [][]string
}

func (m *Mapper) Translate(s string) string {
	for _, row := range m.matrix {
		if row[0] == s {
			if len(row) < 2 {
				return " "
			}
			return row[1]
		}
	}
	return s
}

func main() {
	m, err := NewMapper("mappings.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile("z340.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range string(b) {
		fmt.Print(m.Translate(fmt.Sprintf("%c", c)))
	}
}

func NewMapper(filename string) (mapper *Mapper, e error) {
	m := make([][]string, 63)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			e = err
		}
	}()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	symbols := scanner.Text()
	for i, c := range symbols {
		m[i] = []string{fmt.Sprintf("%c", c)}
	}
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			m[i] = append(m[i], fmt.Sprintf("%c", c))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &Mapper{m}, nil
}

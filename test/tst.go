package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var (
	ErrOpenFile = errors.New("file open error")
)

func ReadFile(fl string) ([]string, error) {
	file, err := os.Open(fl)
	if err != nil {
		return nil, ErrOpenFile
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	text := make([]string, 0)
	for sc.Scan() {
		text = append(text, sc.Text())
	}

	return text, nil
}

func Formater(s string) func(s string) (string, error) {
	i := 0
	return func(s string) (string, error) {
		i++
		slc := strings.Split(s, "|")
		if slices.Contains(slc, "") {
			panic(fmt.Sprintf("empty field on string %d", i))
		}
		return fmt.Sprintf("%d\nName: %s\nAddress: %s\nCity: %s\n\n\n", i, slc[0], slc[1], slc[2]), nil
	}
}

func WriteFile(name string, data []string) error {

	fl, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()

	wr := bufio.NewWriter(fl)

	Formater := Formater("0|0|0")
	for _, row := range data {
		s, _ := Formater(row)

		_, err := wr.WriteString(s)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer wr.Flush()
	return nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			dt, err := ReadFile("data/data_out.txt")
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range dt {
				fmt.Println(v)
			}
		}
	}()

	data, err := ReadFile("07_task_in.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = WriteFile("data/data_out.txt", data)
	if err != nil {
		log.Fatal(err)
	}

}

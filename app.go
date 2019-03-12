package csv

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Csv struct {
	Len     int
	Content string
}

type ReadCsv struct {
	Title []string
	Row   [][]string
}

func New() *Csv {
	return &Csv{Len: 0, Content: ""}
}

func Read(path string) (*ReadCsv, error) {

	fi, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	newCsv := &ReadCsv{Title: []string{}, Row: [][]string{}}

	defer fi.Close()

	br := bufio.NewReader(fi)

	isFirstLine := true

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		line := strings.Split(string(a), ",")

		if isFirstLine {
			newCsv.Title = append(newCsv.Title, line...)
		} else {
			newCsv.Row = append(newCsv.Row, line)
		}

		isFirstLine = false
	}

	return newCsv, nil
}

func (c *Csv) WriteTitle(s []string) *Csv {

	c.Len = len(s)
	c.Content = strings.Join(s, ",")
	c.Content += "\n"

	return c
}

func (c *Csv) WriteRow(s []string) *Csv {

	c.Content += strings.Join(s, ",")
	c.Content += "\n"
	return c
}

func (c *Csv) String() string {
	return c.Content
}

func (c *Csv) Byte() []byte {
	return []byte(c.Content)
}

func (c *Csv) ToFile(path string) {
	buffer := c.Byte()
	ioutil.WriteFile(path, buffer, 0754)
}

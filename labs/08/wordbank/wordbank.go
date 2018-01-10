package wordbank

import (
	"bufio"
	"math/rand"
	"os"
)

type WordBank struct {
	bank []string
}

func New(f *os.File) (*WordBank, error) {
	wb := &WordBank{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wb.bank = append(wb.bank, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return wb, nil
}

func (wb *WordBank) GetWord() string {
	index := rand.Intn(len(wb.bank))
	return wb.bank[index]
}

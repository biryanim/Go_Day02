package wc

import (
	"bufio"
	"os"
	"sync"
)

func Counter(filename string, wg *sync.WaitGroup, strategy func(data []byte, atEOF bool) (advance int, token []byte, err error)) (int, error) {
	defer wg.Done()
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(strategy)
	var count int
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

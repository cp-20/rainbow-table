package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func createChain(start string) string {
	current := start

	for i := 0; i < chainLength; i++ {
		hash := sha256.Sum256([]byte(current))
		hashStr := hex.EncodeToString(hash[:])
		current = reduce(hashStr, i)
	}

	return current
}

func worker(id int, chainsToGenerate int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	for i := 0; i < chainsToGenerate; i++ {
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		start := randomString(rand, 6)
		end := createChain(start)
		results <- fmt.Sprintf("%s:%s\n", start, end)

		if i%1000 == 0 {
			fmt.Printf("Worker %d generated %d chains\n", id, i)
		}
	}
}

func createTable() {
	f, err := os.Create(rainbowTableFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	chainsPerWorker := numberOfChains / goroutineNum
	results := make(chan string, numberOfChains)

	var wg sync.WaitGroup
	wg.Add(goroutineNum)

	for i := 0; i < goroutineNum; i++ {
		go worker(i, chainsPerWorker, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for chain := range results {
		_, err := f.WriteString(chain)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Rainbow table created successfully.")
}

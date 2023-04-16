package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func reverseHash(targetHash string) (string, error) {
	file, err := os.Open(rainbowTableFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	for i := chainLength - 1; i >= 0; i-- {
		reduced := reduce(targetHash, i)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				return "", fmt.Errorf("invalid line in rainbow table: %s", line)
			}

			if parts[1] == reduced {
				plainText := parts[0]
				for j := 0; j <= i; j++ {
					hash := sha256.Sum256([]byte(plainText))
					hashStr := hex.EncodeToString(hash[:])
					if hashStr == targetHash {
						return plainText, nil
					}
					plainText = reduce(hashStr, j)
				}
			}
		}

		_, err := file.Seek(0, 0)
		if err != nil {
			return "", err
		}
	}

	return "", fmt.Errorf("failed to reverse hash")
}

func reverse() {
	targetHash := "c96c6d5beed8889afcc8d3e7b3a3a5966b1fae204ee9ab9a01f44d2f8c3d3b31"
	plainText, err := reverseHash(targetHash)
	if err != nil {
		fmt.Printf("Failed to reverse hash: %s\n", err)
	} else {
		fmt.Printf("Hash reversed: %s -> %s\n", targetHash, plainText)
	}
}

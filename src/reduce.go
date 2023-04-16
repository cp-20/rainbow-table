package main

import "math/rand"

func reduce(hash string, pos int) string {
	var value int64
	for _, r := range hash {
		value += int64(r)
	}

	randSeed := value + int64(pos)
	rand := rand.New(rand.NewSource(randSeed))
	return randomString(rand, 6)
}

func randomString(rand *rand.Rand, length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}
	return string(b)
}

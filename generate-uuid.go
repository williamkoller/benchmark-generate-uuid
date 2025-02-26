package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

type UUID [16]byte

// bom pq usa apenas bytes
func NewWithoutString(r io.Reader) (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(r, uuid[:])
	if err != nil {
		var emptyUUID [16]byte
		return emptyUUID, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return uuid, nil
}

// ruim pq usa string :(
func generateUUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return ""
	}

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		NewWithoutString(rand.Reader)
	}
	elapsed := time.Since(start)
	fmt.Printf("Go: %s\n", elapsed)
}

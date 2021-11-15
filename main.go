package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

const printable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"

func main() {

	hash := flag.String("hash", "", "hash to be decrypted")
	hashType := flag.String("type", "sha1", "hash algorithm")
	minLen := flag.Int("min", 1, "min length of message")
	maxLen := flag.Int("max", 10, "max length of message")
	charSet := flag.String("charset", printable, "char set for possible message")
	threadsNum := flag.Int("threads", 256, "max number of threads")
	filePath := flag.String("file", "", "path to file with hashes")
	flag.Parse()

	if *hash != "" {
		now := time.Now()
		defer func() {
			fmt.Println(time.Now().Sub(now))
		}()
		crackHash(*hashType, *hash, *charSet, *minLen, *maxLen, *threadsNum)
	} else if *filePath != "" {
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Println("Failed opening file!")
		}
		now := time.Now()
		defer func() {
			fmt.Println(time.Now().Sub(now))
		}()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			crackHash(*hashType, scanner.Text(), *charSet, *minLen, *maxLen, *threadsNum)
			fmt.Print()
		}

		file.Close()
	} else {
		flag.Usage()
	}
}

func crackHash(hashType string, hash string, charSet string, minLen int, maxLen int, threads int) string {
	cracker := NewHASHCracker(hashType, hash, []uint8(charSet), uint8(minLen), uint8(maxLen))
	res := cracker.Crack(uint32(threads))
	if res == "" {
		fmt.Println("=========> Message wasn't decrypted try to increase max length or change char set.")
	} else {
		fmt.Println("=========> Message:", res)
	}
	return res
}

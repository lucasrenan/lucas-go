package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("nasa-logs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	fileMd5sum, err := os.Open("nasa-logs/md5sum.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileMd5sum.Close()

	scanner := bufio.NewScanner(fileMd5sum)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		fileMd5 := words[0]
		fileName := words[1]

		fileLog, err := os.Open(fmt.Sprintf("%s/%s", "nasa-logs/", fileName))
		if err != nil {
			log.Fatal(err)
		}
		defer fileLog.Close()

		hash := md5.New()

		nothing, err := io.Copy(hash, fileLog)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(nothing)

		hashInBytes := hash.Sum(nil)[:16]

		if fileMd5 == hex.EncodeToString(hashInBytes) {
			fmt.Printf("File %s is matching md5: %s\n", fileName, fileMd5)
		} else {
			fmt.Printf("File %s is not matching md5: %s\n", fileName, fileMd5)
		}

	}
}

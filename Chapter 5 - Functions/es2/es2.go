package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, err
}

func fileLen(name string) (int, error) {
	f, closer, err := getFile(name)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	data := make([]byte, 2048)

	tot := 0
	for {
		count, err := f.Read(data)
		tot += count
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}

	return tot, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	len, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Args[1], "lungo", len, "bytes")
}

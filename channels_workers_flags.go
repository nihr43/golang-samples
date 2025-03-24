package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func compressWorker(jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for f := range jobs {
		fmt.Println("compressing", f)
		inputFile, err := os.Open(f)
		if err != nil {
			log.Fatal(f, "failed:", err)
		}
		defer inputFile.Close()
		outputPath := f + ".gz"
		outputFile, _ := os.Create(outputPath)
		defer outputFile.Close()

		gzipWriter := gzip.NewWriter(outputFile)
		defer gzipWriter.Close()

		_, err = io.Copy(gzipWriter, inputFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	dir := flag.String("dir", ".", "directory")
	threads := flag.Int("threads", 2, "number of workers to spawn")
	flag.Parse()

	jobs := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < *threads; i++ {
		wg.Add(1)
		go compressWorker(jobs, &wg)
	}

	entries, err := os.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		path := *dir + "/" + entry.Name()
		jobs <- path
	}

	close(jobs)
	wg.Wait()
}

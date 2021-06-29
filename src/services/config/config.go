package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func SetEnvFromFile(path string) {
	file, e := os.Open(path)

	if e != nil {
		log.Fatalf("Unable to find: %s", e)
	}

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanLines)

	for input.Scan() {
		line := input.Text()
		enVar := strings.SplitN(line, "=", 2)
		os.Setenv(enVar[0], enVar[1])
	}

	file.Close()
}

func GetEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("Unable to retreive environment variable: %s", key)
	}

	return value
}

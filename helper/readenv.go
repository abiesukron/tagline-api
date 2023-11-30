package helper

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type EnvFileOption struct {
	Key   string
	Value string
}

func ReadFileEnv(dir string) []EnvFileOption {
	file, err := os.OpenFile(dir, os.O_RDONLY, 0666)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	var envFileOption []EnvFileOption

	for {
		line, _, err := reader.ReadLine()

		count := len(strings.Split(string(line), "="))

		if count > 1 {
			envFileOption = append(envFileOption, EnvFileOption{
				Key:   strings.Split(string(line), "=")[0],
				Value: strings.Split(string(line), "=")[1],
			})
		}

		if err == io.EOF {
			break
		}
	}

	return envFileOption

}

func ReadFileEnvToString(dir string, key string) string {
	file, err := os.OpenFile(dir, os.O_RDONLY, 0666)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	var nama string

	for {
		line, _, err := reader.ReadLine()
		count := len(strings.Split(string(line), "="))
		if count > 1 && strings.Split(string(line), "=")[0] == key {
			nama = strings.Split(string(line), "=")[1]
		}

		if err == io.EOF {
			break
		}
	}

	return nama
}

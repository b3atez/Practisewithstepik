package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var files []string

	root := "/Users/nigmatullayev_a/Downloads/task"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	var arr [][]string
	var t []string

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		if strings.Contains(string(data), ",") {
			lst := strings.Split(string(data), ",")

			for _, v := range lst {
				t = append(t, v)

				if len(t) == 10 {

					arr = append(arr, t)
					t = []string{}
				}

			}
		}
	}

	fmt.Println(arr[4][2])

}

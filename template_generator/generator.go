package main

import (
	"fmt"
	"os"
	"path"
)

const (
	TEMPLATE_PATH = "template.go.tmpl"
)

func report_if(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run ./generate.go <dest_directory>")
		return
	}

	aoc_template, err := os.ReadFile(TEMPLATE_PATH)
	report_if(err, fmt.Sprintf("Failed to read %s", TEMPLATE_PATH))

	dest_directory := path.Join("..", args[0])
	mb_error := os.MkdirAll(dest_directory, os.ModePerm)
	report_if(mb_error, fmt.Sprintf("Can't create directory `%s`", dest_directory))

	path_to_generated_file := path.Join(dest_directory, "main.go")
	generated_file, err := os.Create(path_to_generated_file)
	report_if(err, fmt.Sprintf("File `%s` cannot be crated", path_to_generated_file))
	defer generated_file.Close()

	fmt.Printf("Writing content of `%s` to `%s`\n", TEMPLATE_PATH, path_to_generated_file)
	_, err = generated_file.Write(aoc_template)
	report_if(err, fmt.Sprintf("Failed to write `%s` to `%s`", TEMPLATE_PATH, path_to_generated_file))

	input_file := path.Join(dest_directory, "input.txt")
	fmt.Printf("Creating empty file %s", input_file)
	_, err = os.Create(input_file)
	report_if(err, fmt.Sprintf("Failed to create %s", input_file))
}

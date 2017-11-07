package packer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//EnsureDockerfile ..
func EnsureDockerfile() {
	dockerfile, err := getDockerFile()
	if err != nil {
		panic(err)
	}

	if _, pathErr := os.Stat(dockerfile); pathErr == nil {
		return
	}

	fmt.Println("Creating ", dockerfile)
	f, err := os.Create(dockerfile)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "FROM scratch")
	fmt.Fprint(w, "COPY . /package")
	w.Flush()
}

//Pack ...
func Pack(dir string) {
	EnsureDockerfile()
	filename, err := getDockerFile()
	cmd := exec.Command("docker", "build", "-f", filename, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Pack completed.")
}

func getDockerFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	name := "Dockerfile.cpacker"
	filename := filepath.Join(dir, name)
	return filename, nil
}

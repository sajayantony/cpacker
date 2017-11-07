package packer

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

//Pack ...
func Pack(dir string, containerName string) error {
	//EnsureDockerfile()
	//filename, err := getDockerFile()
	//cmd := exec.Command("docker", "build", "-f", filename, dir)
	cmd := exec.Command("docker", "build", "-t", containerName, "-f", "-", dir) //pass docker file through STDIN

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Send the docker file to stdin
	go func() {
		io.WriteString(stdin, "FROM scratch\n")
		io.WriteString(stdin, "COPY . /package\n")
		stdin.Close()
	}()

	if err = cmd.Wait(); err != nil {
		fmt.Println("Pack failed.")
		return err
	}

	fmt.Println("Pack completed.")
	return nil
}

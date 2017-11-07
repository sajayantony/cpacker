package packer

import "os/exec"
import "log"
import "os"

//Unpack ...
func Unpack(image string, unpackLoc string) error {
	const (
		cntName = "cpacker"
		pkgLoc  = "/package"
	)

	//pass in empty entry point to create container
	cntCmd := exec.Command("docker", "create", "--name", cntName, image, "")
	log.Println("Creating container")
	cntCmd.Stdout = os.Stdout
	err := cntCmd.Run()
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer func() {
		cntRmCmd := exec.Command("docker", "rm", cntName)
		cntRmCmd.Run()
	}()

	cpCmd := exec.Command("docker", "cp", cntName+":"+pkgLoc, unpackLoc)
	cpCmd.Stdout = os.Stdout
	log.Println("Copying to" + unpackLoc)
	err = cpCmd.Run()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

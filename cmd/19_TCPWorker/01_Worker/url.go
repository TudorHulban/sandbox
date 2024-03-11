package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/TudorHulban/GolangSandbox/helpers"
)

func runCmd(command, arguments string) error {
	cmd := exec.Command(command, arguments)
	fmt.Println("running: ", cmd.Args)

	if errStart := cmd.Start(); errStart != nil {
		fmt.Println(errStart)

		os.Exit(1)
	}

	fmt.Println("waiting ...", cmd.Args[1])

	return cmd.Wait()
}

func doWork(url string) (string, error) {
	urlHash := helpers.HashSHA256(url)[:16]
	imgPath := ramDisk + urlHash + ".png"
	scriptName := urlHash + ".sh"

	helpers.CreateFile(scriptName, cmd+imgPath+" "+url)

	t1 := time.Now()

	errCmd := runCmd("bash", scriptName)

	elapsed := time.Since(t1)

	if helpers.FileExists(imgPath) {
		return "",
			fmt.Errorf("file: %s not created", imgPath)
	}

	log.Println("seconds: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 1, 64))
	log.Println(imgPath + " exists ?: " + strconv.FormatBool(helpers.FileExists(imgPath)))

	helpers.DeleteFile(scriptName)

	return imgPath, errCmd
}

package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func createFile(path string, content string) error {
	co := []byte(content)

	return os.WriteFile(path, co, 0644)
}

func deleteFile(path string) error {
	return os.Remove(path)
}

func existsFile(path string) bool {
	_, errStat := os.Stat(path)

	return errStat == nil
}

func runCmd(command, args string) error {
	cmd := exec.Command(command, args)
	log.Println("running: ", cmd.Args)

	errStart := cmd.Start()
	if errStart != nil {
		log.Fatal(errStart)
		return errStart
	}

	log.Println("waiting ...", cmd.Args[1])

	return cmd.Wait()
}

func hashURL(url string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:16] // adjust length per traffic
}

func doWork(url string) (string, error) {
	urlHash := hashURL(url)
	imgPath := ramDisk + urlHash + ".png"
	scriptName := urlHash + ".sh"

	errCreateScript := createFile(scriptName, cmd+imgPath+" "+url)
	if errCreateScript != nil {
		return "", errors.New("script not created")
	}

	timeStart := time.Now()
	errCmd := runCmd("bash", scriptName)
	if errCmd != nil || !existsFile(imgPath) {
		return "", errors.New("screenshot not created")
	}

	elapsed := time.Since(timeStart)
	log.Println("seconds: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 1, 64))

	errDelete := deleteFile(scriptName)
	if errDelete != nil {
		log.Println("file: ", scriptName, " could not be deleted")
	}

	return imgPath, errCmd
}

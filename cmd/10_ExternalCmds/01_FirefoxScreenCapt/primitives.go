package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
)

func createFile(path, fileContent string) error {
	return os.WriteFile(path, []byte(fileContent), 0o600)
}

func deleteFile(path string) error {
	return os.Remove(path)
}

func existsFile(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

// runCmd Helper to bash execute a command with given arguments.
// introduced as syscall did not work with firefox.
func runCmd(command, theArguments string) error {
	cmd := exec.Command(command, theArguments)

	fmt.Println("running:", cmd.Args)

	if errStart := cmd.Start(); errStart != nil {
		return errStart
	}

	fmt.Printf("waiting for %s...\n", cmd.Args[1])

	return cmd.Wait()
}

// hashURL Helper to sha the URL in order to have a unique name for the screen capture.
func hashURL(url string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:16] // adjust length per needs
}

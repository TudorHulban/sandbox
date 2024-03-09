package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

const ramDisk = ""
const cmd = "firefox -headless --screenshot "

const (
	u1 = "https://www.olx.ro"
	u2 = "https://www.olx.ro"
)

func init() {
	_ = runCmd("bash", "cleanpng.sh")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	capture(u1, &wg)
	capture(u2, &wg)

	wg.Wait()
}

func capture(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	urlHash := hashURL(url)
	imgPath := ramDisk + urlHash + ".png"

	scriptName := urlHash + ".sh"
	_ = createFile(scriptName, cmd+imgPath+" "+url)

	t := time.Now()

	if errScript := runCmd("bash", scriptName); errScript != nil {
		fmt.Printf(
			"capture url: %s: %s\n",
			url,
			errScript.Error(),
		)

		return
	}

	elapsed := time.Since(t).Milliseconds()
	fmt.Printf("script took %d miliseconds to run.\n", elapsed)

	fmt.Printf("file %s exists? - %s\n", imgPath, strconv.FormatBool(existsFile(imgPath)))

	_ = deleteFile(scriptName)
}

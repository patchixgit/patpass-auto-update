package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var PROGRAM_VERSION = "1.0.0"

func main() {
	fmt.Println("PatPass Auto Update - v" + PROGRAM_VERSION)

	resp, err := http.Get("https://patchix.vip/pw/current-vers")
	if err != nil {
		panic("There was an error getting the current version of the Program. ")
	}

	if resp.StatusCode != 200 {
		fmt.Println("Err: The server did not return an appropriate response code.")
		return
	}

	// read server version properly
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Failed to read server response")
	}
	serverVersion := strings.TrimSpace(string(bodyBytes)) // remove newlines

	// https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
	// if anyone know better way to do it, feel free to make a PR
	currentRunningDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Error getting current running directory: " + err.Error())
		return
	}

	versionTxtFP := filepath.Join(currentRunningDir, "version.txt")

	// https://www.bytesizego.com/blog/reading-file-line-by-line-golang
	// Im a begginer at this stuff, so PRs welcome.
	versionTxt, err := os.Open(versionTxtFP)
	if err != nil {
		log.Fatalf("Failed to open the file : %s", err)
	}
	defer versionTxt.Close()

	scanner := bufio.NewScanner(versionTxt)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// compare local version to server version
		if line != serverVersion {
			fmt.Println("Update available!")
		} else {
			fmt.Println("You are running the latest version.")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading version file: %s", err)
	}

	// TODO: get latest file from server, extract, & save.
}
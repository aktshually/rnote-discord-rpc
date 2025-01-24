package main

import (
	"log"
	"os"
	"strings"
	"time"
)

const (
	CONFIG_FILE_NAME = "rnote_discord_rpc.config.json"
	DISCORD_APP_ID   = "1332132398292271185"
)

func IsRnoteRunning() bool {
	procEntries, err := os.ReadDir("/proc")
	if err != nil {
		log.Println("Could not read directory /proc: ", err)
	}

	for _, procEntry := range procEntries {
		file, err := os.Open("/proc/" + procEntry.Name() + "/status")
		if err != nil {
			continue
		}

		fileContent := make([]byte, 12)
		file.Read(fileContent)

		processName := strings.TrimSpace(strings.Replace(string(fileContent), "Name:", "", 1))
		if processName == "rnote" {
			return true
		}
	}

	return false
}

func main() {
	for {
		select {
		case <-time.After(time.Second * 10):
			if IsRnoteRunning() {
				// Figure out how to integrate with Discord RPC
			}
		}
	}
}

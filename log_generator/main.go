package main

import (
	"fmt"
	"os"
)

func main() {
	log_messages := []string{
		"INFO: Server started successfully",
		"ERROR: Failed to connect to database",
		"INFO: Request processed",
		"ERROR: Timeout occurred while processing request",
	}
	file_count := 5
	message_count := 10

	for i := 1; i <= file_count; i++ {
		filename := fmt.Sprintf("log%d.log", i)
		file, err := os.Create(filename)
		if err != nil {
			continue
		}
		defer file.Close()

		for j := 1; j <= message_count; j++ {
			for _, val := range log_messages {
				file.WriteString(val + "\n")
			}
		}
	}
}

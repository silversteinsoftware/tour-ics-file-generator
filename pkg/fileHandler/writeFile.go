package fileHandler

import (
  "log"
  "os"
)

func WriteFile(filename, events string) {
  err := os.WriteFile(filename, []byte(events), 0644)
  if err != nil {
    log.Fatalf("Error writing ICS event to file: %v", err)
  }
}

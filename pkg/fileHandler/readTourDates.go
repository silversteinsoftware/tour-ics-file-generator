package fileHandler

import (
  "encoding/json"
  "log"
  "os"
  "silversteinsoftware/tour-ics-file-generator/pkg/types"
)

func readFile(filepath string) []byte {
  file, err := os.ReadFile(filepath)
  if err != nil {
    log.Fatalf("Error while reading file: %v", err)
  }

  return file
}

func parseFileContent(file []byte) []types.InputEvent {
  var events []types.InputEvent
  err := json.Unmarshal(file, &events)
  if err != nil {
    log.Fatalf("Error while parsing file content: %v", err)
  }

  return events
}

func ReadTourDates(filepath string) []types.InputEvent {
  file := readFile(filepath)
  content := parseFileContent(file)

  return content
}

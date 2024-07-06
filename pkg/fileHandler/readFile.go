package fileHandler

import (
  "encoding/json"
  "log"
  "os"
)

func parseFileContent[T any](file []byte) T {
  var content T

  err := json.Unmarshal(file, &content)
  if err != nil {
    log.Fatalf("Error while parsing file content: %v", err)
  }

  return content
}

func ReadFile[T any](filename string) T {
  file, err := os.ReadFile(filename)
  if err != nil {
    log.Fatalf("Error while reading file: %v", err)
  }

  parsedContent := parseFileContent[T](file)

  return parsedContent
}

package fileHandler

import (
  "log"
  "os"
  "reflect"
  "silversteinsoftware/tour-ics-file-generator/pkg/types"
  "testing"
)

var testFileContent = []byte(`
[
  {
    "id": 1,
    "showNumber": 1,
    "date": "Jul. 26, 2024",
    "time": "6:00 PM GMT+2",
    "location": "Charlotte, NC",
    "venue": "The Underground"
  }
]
`)

var expectedFileContent = []types.InputEvent{
  {
    ID:         1,
    ShowNumber: 1,
    Date:       "Jul. 26, 2024",
    Time:       "6:00 PM GMT+2",
    Location:   "Charlotte, NC",
    Venue:      "The Underground",
  },
}

func createTempFile(test *testing.T, content []byte) string {
  tmpFile, err := os.CreateTemp("", "events.json")
  if err != nil {
    log.Fatal(err)
  }
  test.Cleanup(func() { os.Remove(tmpFile.Name()) })

  _, err = tmpFile.Write(content)
  if err != nil {
    log.Fatal(err)
  }

  err = tmpFile.Close()
  if err != nil {
    log.Fatal(err)
  }

  return tmpFile.Name()
}

func TestReadFile(test *testing.T) {
  tmpFileName := createTempFile(test, testFileContent)

  actual := readFile(tmpFileName)
  expected := testFileContent
  if !reflect.DeepEqual(actual, expected) {
    test.Errorf("readFile() = %v, want %v", string(actual), string(expected))
  }
}

func TestParseFileContent(test *testing.T) {
  actual := parseFileContent(testFileContent)
  if !reflect.DeepEqual(actual, expectedFileContent) {
    test.Errorf("parseFileContent() = %v, want %v", actual, expectedFileContent)
  }
}

func TestReadTourDates(test *testing.T) {
  tmpFileName := createTempFile(test, testFileContent)

  actual := ReadTourDates(tmpFileName)
  if !reflect.DeepEqual(actual, expectedFileContent) {
    test.Errorf("ReadTourDates() = %v, want %v", actual, expectedFileContent)
  }
}

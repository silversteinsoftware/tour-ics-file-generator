package eventDateParser

import (
  "strings"
  "time"
)

func createTimestamp(timestamp string) string {
  splitTimestamp := strings.Fields(timestamp)

  eventDate := splitTimestamp[0]
  eventTime := splitTimestamp[1]

  parsedDate := strings.Replace(eventDate, "-", "", -1)
  parsedTime := strings.Replace(eventTime, ":", "", -1)

  return parsedDate + "T" + parsedTime + "Z"
}

func ICSDateParser(utcDateTime time.Time) string {
  utcTimestamp := utcDateTime.Format(time.DateTime)
  parsedICSTimestamp := createTimestamp(utcTimestamp)

  return parsedICSTimestamp
}

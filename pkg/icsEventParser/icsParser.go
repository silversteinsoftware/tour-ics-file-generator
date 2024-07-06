package icsEventParser

import (
  "bytes"
  "log"
  "strconv"
  "text/template"
  "time"

  "silversteinsoftware/tour-ics-file-generator/pkg/eventDateParser"
  "silversteinsoftware/tour-ics-file-generator/pkg/icsTemplate"
  "silversteinsoftware/tour-ics-file-generator/pkg/types"
)

func createEndTime(startTime time.Time) time.Time {
  fourHours := 4 * time.Hour
  endTime := startTime.Add(fourHours)

  return endTime
}

func CreateICSEvent(tourDate types.TourEntry, startTime time.Time) string {
  endTime := createEndTime(startTime)

  veventTemplate, err := template.New("vevent").Parse(icsTemplate.Vevent)
  if err != nil {
    log.Fatalf("Error parsing icsTemplate: %v", err)
  }

  parsedDtStamp := eventDateParser.ICSDateParser(time.Now())
  parsedDtStart := eventDateParser.ICSDateParser(startTime)
  parsedDtEnd := eventDateParser.ICSDateParser(endTime)
  showNumber := strconv.Itoa(tourDate.ShowNumber)

  event := &types.ICSEvent{
    UID:         parsedDtStart + "-" + showNumber + "@silverstein.software",
    Summary:     "Silverstein Live",
    Location:    tourDate.Venue + ", " + tourDate.Location,
    Description: "",
    DtStamp:     parsedDtStamp,
    DtStart:     parsedDtStart,
    DtEnd:       parsedDtEnd,
  }

  buf := &bytes.Buffer{}
  if err := veventTemplate.Execute(buf, event); err != nil {
    log.Fatalf("Error executing icsTemplate: %v", err)
  }

  return buf.String()
}

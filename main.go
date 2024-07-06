package main

import (
  "fmt"

  "silversteinsoftware/tour-ics-file-generator/pkg/eventDateParser"
  "silversteinsoftware/tour-ics-file-generator/pkg/fileHandler"
  "silversteinsoftware/tour-ics-file-generator/pkg/icsEventParser"
  "silversteinsoftware/tour-ics-file-generator/pkg/icsTemplate"
  "silversteinsoftware/tour-ics-file-generator/pkg/types"
)

func main() {
  tourDates := fileHandler.ReadFile[[]types.TourEntry]("dates.json")
  timezones := fileHandler.ReadFile[[]types.Timezone]("timezones.json")

  var icsEvent string

  for _, tourDate := range tourDates {
    utcTourDate := eventDateParser.ParseTourDate(tourDate, timezones)
    icsEvent += icsEventParser.CreateICSEvent(tourDate, utcTourDate)
  }

  wrappedICSEvent := fmt.Sprintf(icsTemplate.ICSWrapper, icsEvent)
  fileHandler.WriteFile("events.ics", wrappedICSEvent)
}

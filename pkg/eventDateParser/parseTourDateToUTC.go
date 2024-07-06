package eventDateParser

import (
  "fmt"
  "log"
  "strings"
  "time"

  "silversteinsoftware/tour-ics-file-generator/pkg/types"
)

func getTimezoneAbbreviation(event string) string {
  return strings.Fields(event)[2]
}

func parseDate(eventDate string) string {
  splitDate := strings.Fields(eventDate)

  month := strings.Replace(splitDate[0], ".", "", -1)
  date := strings.Replace(splitDate[1], ",", "", -1)
  year := splitDate[2][2:]

  if len(date) == 1 {
    date = "0" + date
  }

  return date + " " + month + " " + year
}

func parseTime(eventTime string) string {
  splitTime := strings.Fields(eventTime)

  digits := splitTime[0]
  amPm := splitTime[1]
  timezoneAbbreviation := splitTime[2]

  parsedTime, err := time.Parse("3:04 PM", digits+" "+amPm)
  if err != nil {
    log.Fatalf("Error parsing time: %v", err)
  }

  hour := parsedTime.Hour()
  minute := parsedTime.Minute()
  militaryTime := fmt.Sprintf("%02d:%02d", hour, minute)

  return militaryTime + " " + timezoneAbbreviation
}

func parseTourDateToUTC(tzName, parsedDate, parsedTime string) time.Time {
  timezoneLocation, err := time.LoadLocation(tzName)
  if err != nil {
    log.Fatalf("Error loading location: %v", err)
  }

  timeWithLocation, err := time.ParseInLocation(
    time.RFC822, parsedDate+" "+parsedTime,
    timezoneLocation,
  )
  if err != nil {
    log.Fatalf("Error parsing time with location: %v", err)
  }

  return timeWithLocation.UTC()
}

func validateTimezones(
  timezones []types.Timezone,
  eventTimezone,
  parsedDate,
  parsedTime string,
) time.Time {
  for _, timezone := range timezones {
    if eventTimezone == timezone.TzAbbreviation {
      utcEventTime := parseTourDateToUTC(
        timezone.TzIdentifier,
        parsedDate,
        parsedTime,
      )

      return utcEventTime
    }
  }

  return time.Time{}
}

func ParseTourDate(
  event types.TourEntry,
  timezones []types.Timezone,
) time.Time {
  eventTimezone := getTimezoneAbbreviation(event.Time)
  parsedDate := parseDate(event.Date)
  parsedTime := parseTime(event.Time)
  utcEventTime := validateTimezones(
    timezones,
    eventTimezone,
    parsedDate,
    parsedTime,
  )

  return utcEventTime
}

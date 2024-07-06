package types

type TourEntry struct {
  ID         int    `json:"id"`
  ShowNumber int    `json:"showNumber"`
  Date       string `json:"date"`
  Time       string `json:"time"`
  Location   string `json:"location"`
  Venue      string `json:"venue"`
}

type Timezone struct {
  ID             int    `json:"id"`
  TzAbbreviation string `json:"tz_abbreviation"`
  TzName         string `json:"tz_name"`
  TzIdentifier   string `json:"tz_identifier"`
  UTCOffset      string `json:"utc_offset"`
}

type ICSEvent struct {
  UID         string
  Summary     string
  Location    string
  Description string
  DtStamp     string
  DtStart     string
  DtEnd       string
}

package types

type InputEvent struct {
  ID         int    `json:"id"`
  ShowNumber int    `json:"showNumber"`
  Date       string `json:"date"`
  Time       string `json:"time"`
  Location   string `json:"location"`
  Venue      string `json:"venue"`
}

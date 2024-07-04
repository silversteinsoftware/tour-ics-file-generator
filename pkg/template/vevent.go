package template

const vevent = `BEGIN:VCALENDAR
PRODID:-//SILVERSTEIN SOFTWARE//SilversteinTourCalendarFeed//1.0//EN
VERSION:2.0
CALSCALE:GREGORIAN
BEGIN:VEVENT
CLASS:PUBLIC
URL:https://silverstein.software
UID:{{.UID}}
SUMMARY:{{.Summary}}
LOCATION:{{.Location}}
DESCRIPTION:{{.Description}}
DTSTAMP:{{.DtStamp}}
DTSTART:{{.DtStart}}
DTEND:{{.DtEnd}}
END:VEVENT
END:VCALENDAR`

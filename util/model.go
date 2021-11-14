package util

type Date struct {
	Day string
	// ข้างขึ้น (Waxing)
	IsWaxing  bool
	MoonValue int
	Month     int
}

var ThaiNumbers = map[string]string{
	"๐": "0",
	"๑": "1",
	"๒": "2",
	"๓": "3",
	"๔": "4",
	"๕": "5",
	"๖": "6",
	"๗": "7",
	"๘": "8",
	"๙": "9",
}

var ThaiDay = map[string]string{
	"จันทร์":   "monday",
	"อังคาร":   "tuesday",
	"พุธ":      "wednesday",
	"พฤหัสบดี": "thursday",
	"ศุกร์":    "friday",
	"เสาร์":    "saturday",
	"อาทิตย์":  "sunday",
}

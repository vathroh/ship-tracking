package domain

import "time"

type TrackingSummary struct {
	AWB     string
	Courier string
	Service string
	Status  string
	Date    time.Time
	Desc    string
	Amount  string
	Weight  string
}

type TrackingDetail struct {
	Origin      string
	Destination string
	Shipper     string
	Receiver    string
}

type TrackingHistory struct {
	Date     time.Time
	Desc     string
	Location string
}

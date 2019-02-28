package main

import (
	"strconv"
)

var ical []byte

func InitTemplate() {
	ical = []byte("BEGIN:VCALENDAR\nVERSION:2.0\nPRODID:-//hacksw/handcal//NONSGML v1.0//EN")
}

func AppendEvent(index int) {
	ical = append(ical, []byte("\nBEGIN:VEVENT")...)
	ical = append(ical, []byte("\nLOCATION: "+WhereSlice[index])...)
	ical = append(ical, []byte("\nDTSTART:"+DstartSlice[index])...)
	ical = append(ical, []byte("\nDTEND:"+DendSlice[index])...)
	ical = append(ical, []byte("\nSUMMARY: "+WhatSlice[index])...)
	ical = append(ical, []byte("\nRRULE:FREQ=WEEKLY;UNTIL="+strconv.Itoa(semesterfinish)+"T000000")...)
	ical = append(ical, []byte("\nEND:VEVENT")...)
}

func FinitTemplate() {
	ical = append(ical, []byte("\nEND:VCALENDAR")...)
}

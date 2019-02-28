package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//YYYYMMDD
var semesterstart int = 20190304
var semesterfinish int = 20190609
var count int

var IndexSlice []int
var WhenSlice []string
var WhatSlice []string
var WhereSlice []string
var DayIndexSlice []int
var Day [30]int
var DstartSlice []string
var DendSlice []string

func GenerateRawSlices(html []byte) {
	var when string
	var what string
	var where string
	var datapoint int
	htmlraw := string(html)
	for index, element := range htmlraw {
		if string(element) == "c" {
			if htmlraw[index:index+16] == "cssTtableClsSlot" {
				switch {
				case htmlraw[index+16:index+20] == "When":
					substring := htmlraw[index+20:]
					start := strings.IndexAny(substring, ">")
					end := strings.IndexAny(substring, "<")
					when = substring[start+3 : end]
					datapoint++
				case htmlraw[index+16:index+21] == "Where":
					substring := htmlraw[index+21:]
					start := strings.IndexAny(substring, ">")
					end := strings.IndexAny(substring, "<")
					where = substring[start+1 : end]
					datapoint++
				case htmlraw[index+16:index+20] == "What":
					substring := htmlraw[index+20:]
					start := strings.IndexAny(substring, ">")
					end := strings.IndexAny(substring, "<")
					what = substring[start+2 : end]
					datapoint++
				}
				if datapoint%3 == 0 {
					count = count + 1
					IndexSlice = append(IndexSlice, index)
					WhenSlice = append(WhenSlice, when)
					WhereSlice = append(WhereSlice, where)
					WhatSlice = append(WhatSlice, what)
				}
			}
			//Weekdays are preceeded by title
			if htmlraw[index:index+22] == "cssTtbleColHeaderInner" {
				DayIndexSlice = append(DayIndexSlice, index)
			}
		}
	}
}

func FormatDaySlice() {

	for index, elem := range IndexSlice {
		//Monday, discard first element as it does not correspond to a day
		if elem > DayIndexSlice[1] {
			Day[index] = Day[index]
			//Tuesday
			if elem > DayIndexSlice[2] {
				Day[index] = Day[index] + 1
				//Wednesday
				if elem > DayIndexSlice[3] {
					Day[index] = Day[index] + 1
					//Thursday
					if elem > DayIndexSlice[4] {
						Day[index] = Day[index] + 1
						//Friday
						if elem > DayIndexSlice[5] {
							Day[index] = Day[index] + 1
						}
					}
				}
			}
		}
	}
}

func FormatTimeSlices() {

	for i := 0; i < len(WhenSlice); i++ {
		var whenelement []string

		whenelement = strings.Split(WhenSlice[i], "-")
		//whenelement[0] is start, whenelement[1] is end

		start := TimeFormatConvert(whenelement[0])
		end := TimeFormatConvert(whenelement[1])

		date := strconv.Itoa(semesterstart + Day[i])

		DstartSlice = append(DstartSlice, fmt.Sprint(date+start))
		DendSlice = append(DendSlice, fmt.Sprint(date+end))
	}
	return
}

func main() {
	data, err := ioutil.ReadFile("./timetable.html")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Is 'timetable.html' in the same directory as executable?")
	}

	GenerateRawSlices(data)

	fmt.Println("Raw slices built")

	FormatDaySlice()

	fmt.Println("Events set to correct days")

	eventnumber := len(IndexSlice)
	FormatTimeSlices()
	fmt.Println("Times formatted")

	fmt.Println("Initalizing .ical template...")
	InitTemplate()
	for i := 0; i < eventnumber; i++ {
		AppendEvent(i)
	}
	FinitTemplate()

	errd := ioutil.WriteFile("./timetable.ical", ical, 0644)
	if err != nil {
		log.Fatal(errd)
	} else {
		fmt.Println("Success!")
		fmt.Println("timetable.ical written to the local directory")
	}

}

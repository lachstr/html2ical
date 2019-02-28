package main

import (
	"fmt"
	"strconv"
	"strings"
)

func TimeFormatConvert(sliceElement string) (time string) {
	var am bool
	var jam string
	var minutes string
	var hours int

	for _, element := range sliceElement {
		if string(element) == "a" {
			am = true
			break
		} else {
			am = false
		}
	}
	timeslice := strings.Split(sliceElement, ":")
	hours, _ = strconv.Atoi(timeslice[0])
	minutes = timeslice[1]

	if hours != 12 {
		if am == false {
			hours = int(hours) + 12
		}
	}

	if hours < 10 {
		jam = fmt.Sprint("0" + strconv.Itoa(hours))
	} else {
		jam = strconv.Itoa(hours)
	}

	//discard am,pm statments
	minutes = minutes[:2]

	return fmt.Sprint("T" + jam + minutes + "00")
}

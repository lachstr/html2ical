# html2ical [depreciated]
The University of Melbourne uses a new system for timetabling, which supports exporting as .ical. Thus this tool no longer works and will not be updated.

A tool which creates an .ical file from the offical University of Melbourne Student Timetable portal

Compile:
    
    go build ./... in html2ical directory

Use:  Save your entire unimelb 'my timetable' page as "timetable.html" - (The one that opens when you click 'Access my Timetable' from my.unimelb.edu.au)

Place the "timetable.html" in the same folder as "html2ical.exe"

Open "html2ical.exe" and wait (less than 2s), a file "timetable.ical" should be written in the same folder as "html2ical.exe" and "timetable.html"

Import "timetable.ical" into your calendar app!

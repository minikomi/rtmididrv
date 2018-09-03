module github.com/gomidi/rtmididrv

replace github.com/gomidi/rtmididrv/imported/rtmidi => ./imported/rtmidi

require (
	github.com/gomidi/connect v0.10.0
	github.com/gomidi/rtmididrv/imported/rtmidi v0.0.0-20180903191816-feb86b14a13c
)

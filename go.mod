module github.com/gomidi/rtmididrv

replace github.com/gomidi/rtmididrv/imported/rtmidi => ./imported/rtmidi

require (
	github.com/gomidi/connect v0.11.1
	github.com/gomidi/rtmididrv/imported/rtmidi v0.0.0-20180903192224-c212a44b13e6
)

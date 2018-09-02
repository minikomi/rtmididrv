module github.com/gomidi/rtmididrv

replace github.com/gomidi/rtmididrv/imported/rtmidi => ./imported/rtmidi

require (
	github.com/gomidi/connect v0.9.0
	github.com/gomidi/rtmididrv/imported/rtmidi v0.0.0-20180902095240-24e991c33977
)

module github.com/gomidi/rtmididrv

replace github.com/gomidi/rtmididrv/imported/rtmidi => ./imported/rtmidi

require (
	github.com/gomidi/connect v0.11.1
	github.com/gomidi/rtmididrv/imported/rtmidi v0.0.0-20181003214813-394e08a8a616
)

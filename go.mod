module github.com/minikomi/rtmididrv

replace github.com/gomidi/minikomi/imported/rtmidi => ./imported/rtmidi

require (
	github.com/gomidi/connect v0.11.1
	github.com/minikomi/rtmididrv/imported/rtmidi v0.0.0-20190509060538-f5cd780ff5c1
)

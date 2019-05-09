module github.com/minikomi/rtmididrv

replace github.com/gomidi/minikomi/imported/rtmidi => ./imported/rtmidi

require (
	github.com/gomidi/connect v0.11.1
	github.com/gomidi/rtmididrv/imported/rtmidi v0.0.0-20181023173540-4751d32e0b95
)

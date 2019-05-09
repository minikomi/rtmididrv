module github.com/minikomi/rtmididrv

replace github.com/gomidi/minikomi/imported/rtmidi => ./imported/rtmidi

require (
	github.com/gomidi/connect v0.11.1
	github.com/gomidi/minikomi/imported/rtmidi v0.0.0-20181003214813-394e08a8a616
	github.com/metakeule/mutex v0.0.1
)

package rtmididrv

import (
	"fmt"
	"sync"

	"github.com/gomidi/connect"
	"github.com/gomidi/rtmididrv/imported/rtmidi"
)

type driver struct {
	opened []connect.Port
	sync.RWMutex
	destroyed bool
}

func (d *driver) String() string {
	return "rtmididrv"
}

// Close closes all open ports. It must be called at the end of a session.
func (d *driver) Close() (err error) {
	panic("unsupported")
	/*
		d.RLock()
		if d.destroyed {
			d.RUnlock()
			return connect.ErrClosed
		}

		d.Lock()
		d.destroyed = true
		d.Unlock()

		for _, p := range d.opened {
			fmt.Printf("closing %v\n", p)
			err = p.Close()
			if err != nil {
				fmt.Printf("error closing %v: %v\n", p, err)
			}
		}
	*/

	//	not sure about that
	//	out.Destroy()

	// return just the last error to allow closing the other ports.
	// to ensure that all ports have been closed, this function must
	// return nil anyways
	return
}

// New returns a driver based on the default rtmidi in and out
func New() (connect.Driver, error) {
	return &driver{}, nil
}

// Ins returns the available MIDI input ports
func (d *driver) Ins() (ins []connect.In, err error) {
	d.Lock()
	defer d.Unlock()

	if d.destroyed {
		return nil, connect.ErrClosed
	}
	in, err := rtmidi.NewMIDIInDefault()
	if err != nil {
		return nil, fmt.Errorf("can't open default MIDI in: %v", err)
	}

	ports, err := in.PortCount()
	if err != nil {
		return nil, fmt.Errorf("can't get number of in ports: %s", err.Error())
	}

	for i := 0; i < ports; i++ {
		name, err := in.PortName(i)
		if err != nil {
			name = ""
		}
		ins = append(ins, newIn(d, i, name))
	}

	//in.Destroy()
	return
}

// Outs returns the available MIDI output ports
func (d *driver) Outs() (outs []connect.Out, err error) {
	d.Lock()
	defer d.Unlock()
	if d.destroyed {
		return nil, connect.ErrClosed
	}
	out, err := rtmidi.NewMIDIOutDefault()
	if err != nil {
		return nil, fmt.Errorf("can't open default MIDI out: %v", err)
	}

	ports, err := out.PortCount()
	if err != nil {
		return nil, fmt.Errorf("can't get number of out ports: %s", err.Error())
	}

	for i := 0; i < ports; i++ {
		name, err := out.PortName(i)
		if err != nil {
			name = ""
		}
		outs = append(outs, newOut(d, i, name))
	}
	//out.Destroy()
	return
}

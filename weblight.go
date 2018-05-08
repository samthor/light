package light

import (
	"time"
	"github.com/karalabe/gousb/usb"
)

const (
	weblightVendorId     = 0x1209
	weblightProductId    = 0xa800
	weblightMagicRequest = 0x40
	weblightRequestColor = 0x01

	setAfter = time.Second
)

type weblightDevice struct {
	prev    *Color
	context *usb.Context
	device  *usb.Device
	status  int
	lastSet time.Time
}

func (wc *weblightDevice) Set(color Color) error {
	now := time.Now()
	if color.Equal(wc.prev) && now.Sub(wc.lastSet) < setAfter {
		return nil
	}
	var connectedNow bool

	if wc.context == nil {
		c, err := usb.NewContext()
		if err != nil {
			return err
		}
		wc.context = c
		// nb. we should call wc.context.Close() when done, but is that ever?
	}

retry:
	if wc.device == nil {
		wc.prev = nil
		device, err := wc.context.OpenDeviceWithVidPid(weblightVendorId, weblightProductId)
		if err != nil {
			return err
		}
		connectedNow = true
		wc.device = device
	}

	ret, err := wc.device.Control(weblightMagicRequest, weblightRequestColor, 0, 0, color[:])
	if err != nil {
		wc.device.Close()
		wc.device = nil
		if !connectedNow {
			// retry if we didn't already connect during this method
			goto retry
		}
		return err
	}
	wc.lastSet = now
	clone := color
	wc.prev = &clone
	wc.status = ret
	return nil
}

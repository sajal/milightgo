package milight

// Currently only works for White lights cause thats what I got.
// Special thanks to http://www.wifiledlamp.com/service/applamp-api/ for showing me the codes to start with.
// Also http://wifilights.co.nz/blogs/wifi-lights-blog/13851301-open-source-api


import (
	"errors"
	"net"
)

var InvalidZoneErr = errors.New("Invalid zone selected: valid: 0 - 5")
var InvalidBrightnessErr = errors.New("Invalid brightness selected: valid: 1 - 10")

type Controller struct {
	addr string //Address of the wifi controller
}

func NewController(addr string) *Controller {
	return &Controller{addr}
}

//Send command to mi light controler
func (c *Controller) send(data []byte) error {
	conn, err := net.Dial("udp", c.addr)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

//Off all the zones
func (c *Controller) AllOff() error {
	return c.ZoneOff(0)
}

//On all the zones
func (c *Controller) AllOn() error {
	return c.ZoneOn(0)
}

//On specific zone
func (c *Controller) ZoneOn(zone int) error {
	switch zone {
	case 0: //Zone 0 = all zones
		return c.send([]byte{0x35, 0x00, 0x55})
	case 1:
		return c.send([]byte{0x38, 0x00, 0x55})
	case 2:
		return c.send([]byte{0x3D, 0x00, 0x55})
	case 3:
		return c.send([]byte{0x37, 0x00, 0x55})
	case 4:
		return c.send([]byte{0x32, 0x00, 0x55})
	default:
		return InvalidZoneErr
	}
}

//Off specific zone
func (c *Controller) ZoneOff(zone int) error {
	switch zone {
	case 0: //Zone 0 = all zones
		return c.send([]byte{0x39, 0x00, 0x55})
	case 1:
		return c.send([]byte{0x3B, 0x00, 0x55})
	case 2:
		return c.send([]byte{0x33, 0x00, 0x55})
	case 3:
		return c.send([]byte{0x3A, 0x00, 0x55})
	case 4:
		return c.send([]byte{0x36, 0x00, 0x55})
	default:
		return InvalidZoneErr
	}
}

//Increase Brightness
func (c *Controller) ZoneBright(zone int) error {
	//To control individual zone, send on command immidiately followed by bright
	switch zone {
	case 0: //Zone 0 = all zones
		return c.send([]byte{0x3c, 0x00, 0x55})
	case 1:
		c.ZoneOn(zone)
		return c.send([]byte{0x3c, 0x00, 0x55})
	case 2:
		c.ZoneOn(zone)
		return c.send([]byte{0x3c, 0x00, 0x55})
	case 3:
		c.ZoneOn(zone)
		return c.send([]byte{0x3c, 0x00, 0x55})
	case 4:
		c.ZoneOn(zone)
		return c.send([]byte{0x3c, 0x00, 0x55})
	default:
		return InvalidZoneErr
	}
}

//Decrese Brightness
func (c *Controller) ZoneDim(zone int) error {
	//To control individual zone, send on command immidiately followed by bright
	switch zone {
	case 0: //Zone 0 = all zones
		return c.send([]byte{0x34, 0x00, 0x55})
	case 1:
		c.ZoneOn(zone)
		return c.send([]byte{0x34, 0x00, 0x55})
	case 2:
		c.ZoneOn(zone)
		return c.send([]byte{0x34, 0x00, 0x55})
	case 3:
		c.ZoneOn(zone)
		return c.send([]byte{0x34, 0x00, 0x55})
	case 4:
		c.ZoneOn(zone)
		return c.send([]byte{0x34, 0x00, 0x55})
	default:
		return InvalidZoneErr
	}
}

//Set the brightness of a zone. brightness can be 1 - 10.
// This actually dims the bulb all the way down and then brings it up one at a time.
// Have to do this nasty hack because the stupid controler wont tell us the current brightness, nor will it allow us to set it to a specific value.
func (c *Controller) SetBrightness(zone int, brightness int) error {
	if brightness < 1 {
		return InvalidBrightnessErr
	}
	if brightness > 10 {
		return InvalidBrightnessErr
	}
	//Dim it completely
	for i := 0; i < 10; i++ {
		err := c.ZoneDim(zone)
		if err != nil {
			return err
		}
	}
	//Now bring it up one step at a time
	for i := 0; i < brightness; i++ {
		err := c.ZoneBright(zone)
		if err != nil {
			return err
		}
	}
	return nil

}

package v4l2

import (
	"fmt"
)

// PixelFormat returns the pixel format as a FOURCC (4CC) code.
//
// Some examples of what might be returned by PixelFormat are;
//
// • "YUYV",
//
// • "MJPG",
//
// • etc,
func (receiver FormatDescription) PixelFormat() string {
	a := byte(receiver.pixelformat         & 0xff)
	b := byte((receiver.pixelformat >>  8) & 0xff)
	c := byte((receiver.pixelformat >> 16) & 0xff)
	d := byte((receiver.pixelformat >> 24) & 0xff)

	return fmt.Sprintf("%c%c%c%c", a, b, c, d)
}

package main

import (
	cl "colorlib"
)

var (
	windowColor      = cl.RGB{R: 157, G: 174, B: 181}
	doorColor        = cl.RGB{R: 0, G: 0, B: 0}
	wallColor        = cl.RGB{R: 139, G: 69, B: 19}
	defaultTextColor = cl.RGB{R: 0, G: 0, B: 0}
	wall             = cl.Print("   ", wallColor).GetString()
	floor            = cl.Print("   ", cl.RGB{R: 210, G: 180, B: 140}).GetString()
	door             = cl.Print(" D ", doorColor).GetString()
	window           = cl.Print("   ", windowColor).GetString()
	player           = cl.Print(" 𓀠 ", cl.RGB{R: 0, G: 0, B: 0}).GetString()

	tv_on  = cl.RGB{R: 255, G: 228, B: 12}
	tv_off = cl.RGB{R: 0, G: 0, B: 0}
)

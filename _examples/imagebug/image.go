package main

import "github.com/firefly-zero/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

var (
	image      firefly.Image
	subimage   firefly.SubImage
	showImage  bool
	oldButtons firefly.Buttons
)

func boot() {
	// img is 240x160
	image = firefly.LoadFile("img", nil).Image()
	subimage = image.Sub(firefly.P(0, 0), firefly.S(240, 160))
}

func update() {
	buttons := firefly.ReadButtons(firefly.Combined)
	if buttons.JustPressed(oldButtons).Any() {
		showImage = !showImage
	}
	oldButtons = buttons
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)

	if showImage {
		// Full image works:
		image.Draw(firefly.Point{X: 0, Y: 16})
	} else {
		// But drawing sub-image fails:
		subimage.Draw(firefly.Point{X: 0, Y: 16})
	}
}

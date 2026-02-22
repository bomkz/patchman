package guiHandler

import (
	"bytes"
	"image"
	"image/draw"
	"image/gif"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/global/missile"
	"github.com/bomkz/patchman/global/shark"
)

func buildInstaller() {
	go waitForInstall()

	roll := rand.Intn(6) + 1

	img := canvas.NewImageFromImage(nil)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(150, 150))
	installingText := widget.NewLabel("Installing...")
	if roll == 1 {
		playGIF(shark.Shark, img)
	} else {
		playGIF(missile.Missile, img)
	}

	fyne.DoAndWait(func() {
		global.MainWindow.SetContent(container.NewVBox(installingText,
			img))
		global.MainWindow.Resize(fyne.NewSize(250, 100))
	})

}

func playGIF(jif []byte, img *canvas.Image) {

	g, err := gif.DecodeAll(bytes.NewReader(jif))
	if err != nil {
		return
	}

	go func() {
		base := image.NewRGBA(image.Rect(0, 0, g.Config.Width, g.Config.Height))
		var prev *image.RGBA

		for {
			for i, frame := range g.Image {
				// Save state before drawing if next frame needs to restore it
				if i+1 < len(g.Image) && g.Disposal[i+1] == gif.DisposalPrevious {
					prev = image.NewRGBA(base.Bounds())
					draw.Draw(prev, prev.Bounds(), base, image.Point{}, draw.Src)
				}

				// Draw current frame onto base
				draw.Draw(base, frame.Bounds(), frame, frame.Bounds().Min, draw.Over)

				snapshot := image.NewRGBA(base.Bounds())
				draw.Draw(snapshot, snapshot.Bounds(), base, image.Point{}, draw.Src)

				fyne.DoAndWait(func() {
					img.Image = snapshot
					canvas.Refresh(img)
				})

				delay := time.Duration(g.Delay[i]) * 10 * time.Millisecond
				if delay == 0 {
					delay = 100 * time.Millisecond
				}
				time.Sleep(delay)

				// Handle disposal for next frame
				switch g.Disposal[i] {
				case gif.DisposalBackground:
					// Clear the frame area to transparent
					draw.Draw(base, frame.Bounds(), image.Transparent, image.Point{}, draw.Src)
				case gif.DisposalPrevious:
					// Restore to what base looked like before this frame
					if prev != nil {
						draw.Draw(base, base.Bounds(), prev, image.Point{}, draw.Src)
					}
				}
				// gif.DisposalNone / gif.DisposalUnspecified: leave base as-is
			}
		}
	}()
}

var finished = false

func waitForInstall() {

	<-global.Installed
	finished = true
	handleFinish()

}

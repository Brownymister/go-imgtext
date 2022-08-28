package imgtext

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

type ImageData struct {
	BgImgPath string      // path to the image to add the text on
	FontPath  string      // path to a .ttf file containgin the font of the text (always needed)
	FontSize  float64     // size of the text i.e. 20
	Text      string      //
	X         float64     // the X coordinate of the text
	Y         float64     // the Y coorginate of the text
	Color     color.Color // font color
}

// main function to add a text to an Image from an given ImageData structure
func AddTextToImage(imageData ImageData) (image.Image, error) {
	bgImage, err := gg.LoadImage(imageData.BgImgPath)
	if err != nil {
		return nil, err
	}

	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace(imageData.FontPath, imageData.FontSize); err != nil {
		return nil, err
	}

	maxWidth := float64(imgWidth) - 60.0
	dc.SetColor(imageData.Color)
	dc.DrawStringWrapped(imageData.Text, imageData.X, imageData.Y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

// save the given Image at the given path as png
func Save(img image.Image, path string) error {
	if err := gg.SavePNG(path, img); err != nil {
		return err
	}
	return nil
}

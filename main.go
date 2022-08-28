package imgtext

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

type ImageData struct {
	BgImgPath string
	FontPath  string
	FontSize  float64
	Text      string
	X         float64
	Y         float64
	Color     color.Color
}

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

func Save(img image.Image, path string) error {
	if err := gg.SavePNG(path, img); err != nil {
		return err
	}
	return nil
}

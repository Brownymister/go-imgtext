package imgtext

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

type ImageData struct {
	BgImgPath   string      // path to the image to add the text on
	BgImage     image.Image // image data load from file
	FontPath    string      // path to a .ttf file containgin the font of the text (always needed)
	FontSize    float64     // size of the text i.e. 20
	Text        string      //
	X           float64     // the X coordinate of the text
	Y           float64     // the Y coorginate of the text
	Color       color.Color // font color
	OutputImage image.Image // output image with the text
}

type ImageFuncs interface {
	AddTextToImage(Text string, X float64, Y float64, FontSize float64, FontPath string, Color color.Color) error
	loadImage() error

	Save(string) error
}

// new image instance
func NewImage(BgImagePath string) ImageFuncs {

	var img ImageFuncs
	img = &ImageData{BgImgPath: BgImagePath}

	return img
}

// load image from given path into Image
func (img *ImageData) loadImage() error {
	bgImage, err := gg.LoadImage(img.BgImgPath)
	if err != nil {
		return err
	}

	img.BgImage = bgImage
	return nil
}

// main function to add a text to an Image from an given ImageData structure
func (img *ImageData) AddTextToImage(Text string, X float64, Y float64, FontSize float64, FontPath string, Color color.Color) error {
	img.Text = Text
	img.X = X
	img.Y = Y
	img.FontSize = FontSize
	img.FontPath = FontPath
	img.Color = Color

	err := img.loadImage()

	if err != nil {
		return err
	}

	imgWidth := img.BgImage.Bounds().Dx()
	imgHeight := img.BgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(img.BgImage, 0, 0)

	if err = dc.LoadFontFace(img.FontPath, img.FontSize); err != nil {
		return err
	}

	maxWidth := float64(imgWidth) - 60.0
	dc.SetColor(img.Color)
	dc.DrawStringWrapped(img.Text, img.X, img.Y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	img.OutputImage = dc.Image()

	return nil
}

// save the Image at the given path as png
func (img *ImageData) Save(path string) error {
	if err := gg.SavePNG(path, img.OutputImage); err != nil {
		return err
	}
	return nil
}

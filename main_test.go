package imgtext

import (
	"image/color"
	"os"
	"testing"
)

func TestAddTextToImage(t *testing.T) {

	img := NewImage("./assets/image.jpg")
	img.AddTextToImage("Im a Cat!", 112.5, 12, 30, "./assets/DancingScript-VariableFont_wght.ttf", color.Black)
	img.Save("test.png")

	t.Run("Test create Image", func(t *testing.T) {

		if _, err := os.Stat("../test.png"); (err != nil) != true {
			t.Errorf("CreateScoreInDb() error = %v, wantErr %v", err, true)
		}
	})

}

func TestAddTextToImageExpectFail(t *testing.T) {
	type args struct {
		imageData ImageData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test fail expected to load original image",
			args: args{
				imageData: ImageData{
					BgImgPath: "./this/is/a/path/that/dose/not/exsist",
					FontPath:  "./assets/DancingScript-VariableFont_wght.ttf",
					FontSize:  10,
					Text:      "Im a Cat!",
				},
			},
			wantErr: true,
		},

		{
			name: "Test fail expected to load font",
			args: args{
				imageData: ImageData{
					BgImgPath: "./assets/image.jpg",
					FontPath:  "./this/is/a/path/that/dose/not/exsist",
					FontSize:  10,
					Text:      "Im a Cat!",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := NewImage(tt.args.imageData.BgImgPath)
			err := img.AddTextToImage(tt.args.imageData.Text, tt.args.imageData.X, tt.args.imageData.Y, tt.args.imageData.FontSize, tt.args.imageData.FontPath, tt.args.imageData.Color)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTextToImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

package imgtext

import (
	"image/color"
	"os"
	"testing"
)

func TestAddTextToImage(t *testing.T) {
	type args struct {
		imageData ImageData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test create img with text",
			args: args{ImageData{
				BgImgPath: "./assets/image.jpg",
				FontPath:  "./assets/DancingScript-VariableFont_wght.ttf",
				FontSize:  30,
				Text:      "Im a Cat!",
				X:         112.5,
				Y:         12,
				Color:     color.Black,
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := AddTextToImage(tt.args.imageData)
			Save(got, "test.png")
			if _, err := os.Stat("../test.png"); (err != nil) != true {
				t.Errorf("CreateScoreInDb() error = %v, wantErr %v", err, true)
			}
		})
	}
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
			_, err := AddTextToImage(tt.args.imageData)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTextToImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

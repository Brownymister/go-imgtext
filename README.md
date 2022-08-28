# go-imgtext
Simple Go package for adding Text to an Image

# Install

```bash
go get github.com/Brownymister/imgtext
```

# Exmaple

Simple example of creating a new image with an specified text.

```go
imageData := ImageData
    {
        BgImgPath: "/path/to/original/image",
        FontPath:  "/path/to/a/ttf/file",
        FontSize:  30,
        Text:      "Im a Cat!",
        X:         112.5,
        Y:         12,
        Color:     color.Black,
    }

// ceate new image
img := AddTextToImage(imageData)
// save created image with path
Save(img, "path/whete/new/image/has/to/be/saved")

```
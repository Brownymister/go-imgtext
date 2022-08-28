# go-imgtext
Simple Go package for adding Text to an Image

# Install

```bash
go get github.com/Brownymister/imgtext
```

### testing
```bash
go test ./... -cover
```

# Exmaple

Simple example of creating a new image with an specified text.

```go
    // init a new Image instance
    img := NewImage("./assets/image.jpg")
    // add Text to Image
	img.AddTextToImage("Im a Cat!", 112.5, 12, 30, "./assets/DancingScript-VariableFont_wght.ttf", color.Black)
    // save Image on drive
	img.Save("test.png")

```
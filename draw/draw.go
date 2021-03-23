package draw

import (
	"fmt"
	"log"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

// RotateImage image so text is horizontal
func RotateImage(image string) error {
	img, err := imgio.Open(image)
	if err != nil {
		log.Fatal(err)
		return err
	}
	result := transform.Rotate(img, -90, &transform.RotationOptions{
		ResizeBounds: true,
	})
	// result := effect.EdgeDetection(img, 2.5)
	if err := imgio.Save("./sample/horizontal.png", result, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
	}
	return nil
}

// DrawBoxes draws BoundingPolys and returns as new image
func DrawBoxes(image string) {

}

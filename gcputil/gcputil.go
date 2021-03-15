package gcputil

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

// Vertex stores X, Y points
type Vertex struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// BoundingPoly stores array of 4 XY coordinates that make up a polygon
type BoundingPoly struct {
	//using google's generated protobuf.. probably should do this for the other structs if this works
	Vertices []*visionpb.Vertex `json:"vertices"`
}

// Annotation stores Description and BoundingPoly of Description as 'text'
type Annotation struct {
	Description  string       `json:"text"`
	BoundingPoly BoundingPoly `json:"boundingPoly"`
}

// ImageAnnotation stores data returned from Cloud Vision along with image file metadata
type ImageAnnotation struct {
	DirPath       string        `json:"dirPath"`
	FileName      string        `json:"filename"`
	FileExtension string        `json:"fileExtension"`
	Annotations   []*Annotation `json:"annotation"`
}

// NewImageAnnotation creates a new ImageAnnotation from filename and extension
func NewImageAnnotation(filename string) *ImageAnnotation {
	base := filepath.Base(filename)
	file := strings.Split(base, ".")
	dirPath := filepath.Dir(filename)

	// make a new ImageAnnotation and return a pointer to it
	a := ImageAnnotation{
		DirPath:       dirPath,
		FileName:      file[0],
		FileExtension: file[1],
	}
	return &a
}

// DetectText sends image to GCP Cloud Vision for processing and annotating
func DetectText(file string) ([]*visionpb.EntityAnnotation, error) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}

	response, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return nil, err
	}

	if len(response) == 0 {
		fmt.Println("No text detected")
	} else {
		fmt.Println("Text:")
		for _, annotation := range response {
			fmt.Printf(annotation.Description)
		}
	}
	return response, nil
}

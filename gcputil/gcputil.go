package gcputil

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	// Vertices     []Vertex     `json:"vertices"`
}

// ImageAnnotation stores data returned from Cloud Vision along with image file metadata
type ImageAnnotation struct {
	DirPath       string        `json:"dirPath"`
	FileName      string        `json:"filename"`
	FileExtension string        `json:"fileExtension"`
	Annotations   []*Annotation `json:"annotation"`
}

// Creates a new ImageAnnotation from filename and extension
func newImageAnnotation(filename string) *ImageAnnotation {
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

// WriteToFile creates <filename>.json if file doesn't exist, or appends if it does. Example from https://golang.org/pkg/os/#OpenFile
func WriteToFile(filename string, json []byte) {
	// func WriteToFile(annotation *ImageAnnotation) {

	// Create the JSON file
	// file, err := os.OpenFile(fmt.Sprintf("%s.json", filename), os.O_CREATE|os.O_WRONLY, 0644)
	file, err := os.Create(filename + ".json")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := file.Write(json); err != nil {
		file.Close()
		// ignore error; Write error takes precedence
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

// ProcessImage combines all the steps to getting annotations and saving as JSON file
func ProcessImage(file string) error {

	// Create new ImageAnnotation struct with file info and point to it
	imgAnno := newImageAnnotation(file)

	res, err := DetectText(file)
	if err != nil {
		log.Fatal(err)
	}

	// Pull the parts out of EntityAnnotation returned from DetectText (line 76).. kinda redundant for now probably

	for i := 0; i < len(res); i++ {
		annotation := &Annotation{
			Description: res[i].Description,
		}
		// Using google's generated pb thingy
		vertices := res[i].BoundingPoly.GetVertices()
		annotation.BoundingPoly.Vertices = vertices

		imgAnno.Annotations = append(imgAnno.Annotations, annotation)
	}

	jsonOut, err := json.MarshalIndent(imgAnno, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	WriteToFile(fmt.Sprintf("%s/%s", imgAnno.DirPath, imgAnno.FileName), jsonOut)
	return nil
}

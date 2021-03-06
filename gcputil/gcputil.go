package gcputil

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	vision "cloud.google.com/go/vision/apiv1"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

// Vertex stores X, Y points
type Vertex struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Vertices stores 4 XY coordinates that make up a polygon
type Vertices struct {
	Vertices []Vertex `json:"vertices"`
}

// ImageAnnotation stores data returned from Cloud Vision
type ImageAnnotation struct {
	Filename      string     `json:"filename"`
	FileExtension string     `json:"fileExtension"`
	Text          string     `json:"text"`
	BoundingPoly  []Vertices `json:"boundingPoly"`
}

// func newImageAnnotation(filename string) ImageAnnotation {
//         // make a new ImageAnnotation and return a pointer to it
// 	a := ImageAnnotation{
// 		Filename:      filepath.Base(filename),
// 		FileExtension: filepath.Ext(filename),
// 	}
//         return &a
// }

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

// WriteToFile creates filename.json if file doesn't exist, or appends if it does. Example from https://golang.org/pkg/os/#OpenFile
// func WriteToFile(annotation *ImageAnnotation) {
func WriteToFile(filename string, json []byte) {

	// Create the JSON file
	file, err := os.OpenFile(fmt.Sprintf("%s.json", filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	res, err := DetectText(file)
	if err != nil {
		log.Fatal(err)
	}
	var annotations []*ImageAnnotation
	// var vertices []Vertices
	for i := 0; i < len(res); i++ {
		// temp := []interface{}{res[i].BoundingPoly.Vertices}
		// vert, err := json.Marshal(temp)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// vertices = append(vertices, string(vert))
		anno := &ImageAnnotation{
			Filename:      filepath.Base(file),
			FileExtension: filepath.Ext(file),
			Text:          res[i].Description,
			// BoundingPoly:  []Vertices{vertices},
		}
		annotations = append(annotations, anno)
	}
	// jsonOut, err := json.Marshal(annotations)
	// if err != nil {
	//         log.Fatal(err)
	// }
	jsonOut, err := json.MarshalIndent(annotations, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	WriteToFile(filepath.Base(file), jsonOut)
	return nil
}

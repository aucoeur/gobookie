package file

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aucoeur/gobookie/gcputil"
)

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
	imgAnno := gcputil.NewImageAnnotation(file)

	res, err := gcputil.DetectText(file)
	if err != nil {
		log.Fatal(err)
	}

	// Pull the parts out of EntityAnnotation returned from DetectText (line 76).. kinda redundant for now probably

	for i := 0; i < len(res); i++ {
		annotation := &gcputil.Annotation{
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

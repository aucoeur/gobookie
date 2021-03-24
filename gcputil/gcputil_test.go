package gcputil

import (
	"testing"
)

func TestNewImageAnnotation(t *testing.T) {
	sample := ImageAnnotation{
		DirPath:       "testPath",
		FileName:      "testImage",
		FileExtension: ".jpg",
	}

	got := NewImageAnnotation("testPath/testImage.jpg")
	want := sample

	if got.DirPath != want.DirPath {
		t.Errorf("got %s, want %s", got.DirPath, want.DirPath)
	}
}

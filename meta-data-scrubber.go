// meta-data-scrubber
// a tool to scrub metadata from files
// should be able to scrub metadata from jpeg, png, and pdf files
// and also Excel and Word files
// if any private data is found, it should be able to remove it

// scrubber/scrubber.go
package scrubber

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type Metadata struct {
	FileName string                 `json:"file_name"`
	Exif     map[string]interface{} `json:"exif"`
}

func ScrubImageMetadata(filename string) (*Metadata, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

x, err := exif.Decode(file)
if err != nil {
    return nil, err
}

exifData := make(map[string]interface{})
x.Walk(func(name exif.FieldName, val *tiff.Tag) error {
    exifData[string(name)] = val
    return nil
})



metadata := &Metadata{
    FileName: filename,
    Exif:     exifData,
}

// Implement the exif.Walker interface
func (m *Metadata) Walk(name exif.FieldName, tag *tiff.Tag) error {
    m.Exif[string(name)] = tag
    return nil
}

func ScrubImageMetadata(filename string) (*Metadata, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    x, err := exif.Decode(file)
    if err != nil {
        return nil, err
    }

    exifData := make(map[string]interface{})
    x.Walk(func(name exif.FieldName, val *tiff.Tag) error {
        exifData[string(name)] = val
        return nil
    })

    metadata := &Metadata{
        FileName: filename,
        Exif:     exifData,
    }

    x.Walk(metadata)

    return metadata, nil
}

func PrintMetadata(metadata *Metadata) {
    b, err := json.MarshalIndent(metadata, "", "  ")
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Print(string(b))
}

package util

import (
	"encoding/xml"
	"os"
	"path/filepath"
)

// ReadFile reads a file from the specified file path and returns a pointer to the opened file.
// It first obtains the absolute path of the file using filepath.Abs function.
// If an error occurs while obtaining the absolute path, it logs the error and returns it.
// Then, it opens the file using os.Open function.
// If an error occurs while opening the file, it logs the error and returns it.
// Finally, it defers the closing of the file and returns the opened file and nil error if successful.
func readFile(filePath string, closeManually bool) (*os.File, error) {
	fileConfigAbsolutePath, err := filepath.Abs(filePath)
	if err != nil {
		LoggerHandler().Error("Error getting absolute path", "error", err.Error())
		return nil, err
	}

	file, err := os.Open(fileConfigAbsolutePath)
	if err != nil {
		LoggerHandler().Error("Error opening file", "error", err.Error())
		defer file.Close()
		return nil, err
	}
	if !closeManually {
		defer file.Close()
	}

	return file, nil
}

// ReadXmlFile reads an XML file from the given file path and decodes its contents into the provided destination object.
// The filePath parameter specifies the path of the XML file to be read.
// The desfine parameter is a pointer to the destination object where the decoded XML data will be stored.
// It returns an error if there was an issue reading the file or decoding the XML data.
func ReadXmlFile[T any](filePath string, desfine *T) error {
	file, error := readFile(filePath, true)
	if error != nil {
		return error
	}
	if errorDecoding := xml.NewDecoder(file).Decode(desfine); errorDecoding != nil {
		LoggerHandler().Error("Error decoding xml file", "error", errorDecoding.Error())
		return errorDecoding
	}
	defer file.Close()
	return nil
}

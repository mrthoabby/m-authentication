package helpers

import (
	"encoding/xml"
	"os"
	"path/filepath"

	"com.github/mrthoabby/m-authentication/types"
	"com.github/mrthoabby/m-authentication/util"
	"github.com/gin-gonic/gin"
)

// ReadFile reads a file from the specified file path and returns a pointer to the opened file.
// It first obtains the absolute path of the file using filepath.Abs function.
// If an error occurs while obtaining the absolute path, it logs the error and returns it.
// Then, it opens the file using os.Open function.
// If an error occurs while opening the file, it logs the error and returns it.
// Finally, it defers the closing of the file and returns the opened file and nil error if successful.
func readFile(filePath string) (*os.File, error) {
	fileConfigAbsolutePath, err := filepath.Abs(filePath)
	if err != nil {
		util.LoggerHandler().Error("Error getting absolute path", "error", err.Error())
		return nil, err
	}

	file, err := os.Open(fileConfigAbsolutePath)
	if err != nil {
		util.LoggerHandler().Error("Error opening file", "error", err.Error())
		return nil, err
	}
	defer file.Close()

	return file, nil
}

// ReadXmlFile reads an XML file from the given file path and decodes its contents into the provided destination object.
// The filePath parameter specifies the path of the XML file to be read.
// The desfine parameter is a pointer to the destination object where the decoded XML data will be stored.
// It returns an error if there was an issue reading the file or decoding the XML data.
func ReadXmlFile[T any](filePath string, desfine *T) error {
	file, error := readFile(filePath)
	if error != nil {
		return error
	}
	if errorDecoding := xml.NewDecoder(file).Decode(desfine); errorDecoding != nil {
		util.LoggerHandler().Error("Error decoding xml file", "error", errorDecoding.Error())
		return errorDecoding
	}
	return nil
}

// Binder is a function that binds the request body to a given type.
// It takes the content type, the context, and a pointer to the result type as parameters.
// The function returns an error if the binding fails.
func Binder[R any](contentType string, context *gin.Context, result *R) error {

	var binder types.BinderStrategy[R]

	switch contentType {
	case "application/json":
		binder.SetStrategy(&types.JSONBinder[R]{})
	case "application/xml":
		binder.SetStrategy(&types.XMLBinder[R]{})
	case "application/x-www-form-urlencoded":
		binder.SetStrategy(&types.FORMBinder[R]{})
	default:
		return types.NewCustomError("Content type not supported")
	}

	if errorBinding := binder.Bind(context, result); errorBinding != nil {
		return errorBinding
	}
	return nil
}

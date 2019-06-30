package spreadsheet

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// Returns a CSV file as a string based on a spreadsheet path.
// Uses LibreOffice to convert Spreadsheet to CSV.
// path is the path of the spreadsheet file.
func GetCSVFromSpreadsheet(path string) (string, error) {
	csvPath := fmt.Sprintf("%s.csv", strings.Split(path, ".")[0]) // NOTE: does not support paths with dots in path or name
	var err error
	var file []byte

	_, err = exec.Command("soffice", "--headless", "--convert-to", "csv", path).Output()
	fmt.Println("soffice", "--headless", "--convert-to", csvPath, path)
	if err != nil {
		return "", err
	}
	defer exec.Command("rm", csvPath).Output()

	file, err = ioutil.ReadFile(csvPath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

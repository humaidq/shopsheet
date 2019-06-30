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
func GetCSVFromSpreadsheet(path string, dir string) (string, error) {
	csvPath := fmt.Sprintf("%s.csv", strings.Split(path, ".")[0]) // NOTE: does not support paths with dots in path or name
	var err error
	var o, file []byte

	o, err = exec.Command("soffice", "--headless", "--convert-to", "csv", "--outdir", dir, path).Output()
	fmt.Printf("output: %s\n", o)
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

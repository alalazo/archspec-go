package archspec

import (
	"bufio"
	"github.com/markbates/pkger"
)

// ReadJSONAsStringVec reads microarchitectures.json and returns each
// line in a string vector
func ReadJSONAsStringVec() ([]string, error) {
	pkger.Include("/archspec/json")
	filename := "/archspec/json/cpu/microarchitectures.json"
	file, err := pkger.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	return txtlines, nil
}

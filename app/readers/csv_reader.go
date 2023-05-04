package readers

import (
	"encoding/csv"
	"os"
)

type csvReader struct {
	filepath string
}

func NewCsvReader(filepath string) Reader {
	return &csvReader{filepath: filepath}
}

func (c *csvReader) GetData() ([][]string, error) {
	file, err := os.Open(c.filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	return reader.ReadAll()
}

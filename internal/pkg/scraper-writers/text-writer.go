package writers

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// FileWriter - One implementation of the Writer interface. This one will write to txt files, and generate file names like yyyy-mm-dd-exchange-market.txt.
type FileWriter struct{}

// GetWriteFileName - Will generate the file name for you of the format yyyy-mm-dd-exchange-market.txt. New files will be created at midnight because the scrapers are calling this method each time before writing to file.
func (f *FileWriter) GetWriteFileName(exchange string, market string) string {
	now := time.Now()
	return f.clean(fmt.Sprintf("%v-%v-%v-%v-%v.txt", now.Year(), int(now.Month()), now.Day(), exchange, market))
}

// Write - Will write to the filename the line.
func (f *FileWriter) Write(line string, filename string) (int, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return 0, err
	}
	n, err := file.WriteString(line)
	if err != nil {
		return n, err
	}
	return n, nil
}

func (f *FileWriter) clean(s string) string {
	return strings.Replace(s, "/", "_", -1)
}

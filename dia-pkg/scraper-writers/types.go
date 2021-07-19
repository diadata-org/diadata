package writers

// Writer is an interface that is responsible for generating dynamic file names (to create new files at midnight) and writing data to them
type Writer interface {
	GetWriteFileName(exchange string, market string) string // will return the name of the file in which it will write
	Write(line string, filename string) (int, error)        // returns number of bytes written or error
	// rationale for making a line a pointer - is because it can be very large. filename will always be small.
}

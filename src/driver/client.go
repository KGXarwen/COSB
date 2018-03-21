package driver

// Error used for error message

// Client is an interface for different driver
type Client interface {
	SendRequest(data []byte) (bool, error)
}

// Driver interface for different storage backend driver
type Driver interface {
	Get(key string) ([]byte, error)
	Put(fileName string, fileSize int64) (fileKey string, err error) //just need fileSize because the file content should be generated by fakereader
	Delete(key string) error
}
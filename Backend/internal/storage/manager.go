package storage

import (
	"os"

	"github.com/ItsMe-RiiK/LocalWebxMCP/internal/config"
)

// Init creates the necessary storage directory if it doesn't exist
func Init() error {
	return os.MkdirAll(config.StorageDir, os.ModePerm)
}

package multipart

import (
	"path/filepath"
	"mime"
	"strings"
)

var (
	_mimes = map[string]string {
		".txt": "text/plain",
		".csv": "text/csv",
		".xml": "text/xml",
		".json": "application/json",
		".js": "application/javascript",
		".png": "image/png",
		".jpg": "image/jpeg",
		".jpeg": "image/jpeg",
		".gif": "image/gif",
		".svg": "image/svg+xml",
		".mp4": "video/mp4",
		".mpg": "video/mpeg",
		".mpeg": "video/mpeg",
	}
)

const (
	_defMime = "application/octet-stream"
)

func FileContentType(fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))
	if ct, ok := _mimes[ext]; ok {
		return ct
	}
	if ct := mime.TypeByExtension(ext); len(ct) > 0 {
		return ct
	}
	return _defMime
}

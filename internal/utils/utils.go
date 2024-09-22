package utils

import (
	"fmt"
)

func FormatFileSize(fileSize float64) (string) {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	i := 0
	for fileSize >= 1024 && i < len(sizes)-1 {
		fileSize /= 1024
		i++
	}
	return fmt.Sprintf("%.2f%s", fileSize, sizes[i])
}

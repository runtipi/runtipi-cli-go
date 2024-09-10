package utils

import (
	"fmt"
	"strconv"
)

func FormatFileSize(fileSize float64) (string) {
	if fileSize < (1<<10) {
		return fmt.Sprintf("%sB", strconv.Itoa(int(fileSize)))
	} else if fileSize > (1<<10) && fileSize < (1<<20) {
		return fmt.Sprintf("%.2fKB", fileSize/(1<<10))
	} else if fileSize > (1<<20) && fileSize < (1<<30) {
		return fmt.Sprintf("%2.fMB", fileSize/(1<<20))
	} else if fileSize > (1<<30) && fileSize < (1<<40) {
		return fmt.Sprintf("%.2fGB", fileSize/(1<<30))
	} else if fileSize > (1<<40) {
		return fmt.Sprintf("%.2fTB", fileSize/(1<<40))
	}
	return fmt.Sprintf("%sB", strconv.Itoa(int(fileSize)))
}
package libs

import "log"

type ProgressWriter struct {
	Total      int64
	Downloaded int64
	Name       string
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Downloaded += int64(n)
	downloadedMB := float64(pw.Downloaded) / (1024 * 1024)
	totalMB := float64(pw.Total) / (1024 * 1024)
	log.Printf("\r%s [%.2f/%.2f MB (%.2f%%)]", pw.Name, downloadedMB, totalMB, downloadedMB*100/totalMB)
	return n, nil
}

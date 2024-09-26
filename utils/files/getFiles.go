package files

import (
	"os"
	"path/filepath"
	"strings"
)

func GetFilePaths(directory string, extensions ...string) ([]string, error) {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if len(extensions) > 0 {
				for _, ext := range extensions {
					if strings.HasSuffix(info.Name(), ext) {
						files = append(files, path)
						break
					}
				}
			} else {
				files = append(files, path)
			}
		}
		return nil
	})

	return files, err
}

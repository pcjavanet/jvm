//解析使用*的路径

package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //remove *
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".Jar") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}

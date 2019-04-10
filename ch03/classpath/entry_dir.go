package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

func (de *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(de.absDir, className)
	data, err := ioutil.ReadFile(fileName)

	return data, de, err
}

func (de *DirEntry) String() string {
	return de.absDir
}

package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(ze.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	for _, f := range r.File {
		if f.Name == className {
			return func() ([]byte, Entry, error) {
				rc, err := f.Open()
				if err != nil {
					return nil, nil, err
				}

				defer func() {
					if err := rc.Close(); err != nil {
						panic(err)
					}
				}()

				data, err := ioutil.ReadAll(rc)
				if err != nil {
					return nil, nil, err
				}

				return data, ze, nil
			}()
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (ze *ZipEntry) String() string {
	return ze.absPath
}

package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry2 struct {
	absPath string
	zipRC   *zip.ReadCloser
}

func newZipEntry2(path string) *ZipEntry2 {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry2{absPath, nil}
}

func (ze2 *ZipEntry2) readClass(className string) ([]byte, Entry, error) {
	if ze2.zipRC == nil {
		if err := ze2.openJar(); err != nil {
			return nil, nil, err
		}
	}

	classFile := ze2.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)

	return data, ze2, err
}

// todo: close zip
func (ze2 *ZipEntry2) openJar() error {
	r, err := zip.OpenReader(ze2.absPath)
	if err == nil {
		ze2.zipRC = r
	}

	return err
}

func (ze2 *ZipEntry2) findClass(className string) *zip.File {
	for _, f := range ze2.zipRC.File {
		if f.Name == className {
			return f
		}
	}

	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}

	defer rc.Close()
	// read class data
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ze2 *ZipEntry2) String() string {
	return ze2.absPath
}

package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry2 struct {
    absPath string
    zipRC *zip.ReadCloser
}

func newZipEntry2(path string) *ZipEntry2 {
    if absPath, err := filepath.Abs(path); err != nil {
        panic(err)
    }

    return &ZipEntry2{absPath, nil}
}

func (self *ZipEntry2) readClass(className string) ([]byte, Entry, error) {
    if self.zipRC == nil {
        if err := self.openJar(); err != nil {
            return nil, nil, err
        }
    }

    if classFile := self.findClass(className); classFile == nil {
        return nil, nil, errors.New("class not found: " + className
    }

    data, err := readClass(classFile)

    return data, self, err
}

// todo: close zip
func (self *ZipEntry2) openJar() error {
    if r, err := zip.OpenReader(self.absPath); err == nil {
        self.zipRC = r
    }

    return err
}

func (self *ZipEntry2) findClass(className string) *zip.File {
    for _, f := range self.zipRC.file {
        if f.Name == className {
            return f
        }
    }

    return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
    if rc, err := classFile.Open(); err != nil {
        return nil, err
    }

    defer rc.Close()
    // read class data
    if data, err := ioutil.ReadAll(rc); err != nil {
        return nil, err
    }

    return data, nil
}

func (self *ZipEntry2) String() string {
    return self.absPath
}

package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
    absPath string
}

func newZipEntry(path string) *ZipEntry {
    if absPath, err := filepath.Abs(path); err != nil {
        panic(err)
    }

    return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
    if r, err := zip.OpenReader(self.absPath); err != nil {
        return nil, nil, err
    }

    defer r.Close()
    for _, f := range r.File {
        if f.Name == className {
            if rc, err := f.Open(); err != nil {
                return nil, nil, err
            }

            defer rc.Close()
            if data, err := ioutil.ReadAll(rc); err != nil {
                return nil, nil, err
            }

            return data, self, nil
        }
    }

    return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
    return self.absPath
}

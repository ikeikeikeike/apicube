// Package storage gonna be implementation
// that stream io processing for memory performance.
//
package storage

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/dsn"
	"github.com/ikeikeikeike/apicube/base/util/logger"
)

// fileStorage provides implementation file object interface.
type fileStorage struct {
	Env util.Environment `inject:""`
	dsn *dsn.FileDSN
}

// Write will create file into the file systems.
func (adp *fileStorage) Write(filename string, data []byte) error {
	folder, filename := adp.dsn.Folder, adp.dsn.Join(filename)

	fi, err := os.Stat(folder)
	if err != nil {
		_ = os.MkdirAll(folder, 0755)
	} else if !fi.IsDir() {
		return fmt.Errorf("[F] %s should be a directory", folder)
	}

	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	fi, err = file.Stat()
	if err != nil {
		return fmt.Errorf("[F] %s file not exists", filename)
	} else if !fi.Mode().IsRegular() {
		return fmt.Errorf("[F] %s should be a file", filename)
	}

	if gzipPtn.MatchString(filename) {
		adp.gzip(file, data)
	} else {
		adp.plain(file, data)
	}

	return nil
}

// Read returns file data from the file systems.
func (adp *fileStorage) Read(filename string) ([]byte, error) {
	var reader io.ReadCloser

	reader, err := os.Open(adp.dsn.Join(filename))
	if err != nil {
		return nil, errors.Wrap(err, "[F] file read failed")
	}

	defer reader.Close()

	if gzipPtn.MatchString(filename) {
		reader, err = gzip.NewReader(reader)
		if err != nil {
			return nil, errors.Wrap(err, "[F] gzip read failed")
		}
	}

	return ioutil.ReadAll(reader)
}

// Merge will merge file into the file systems.
func (adp *fileStorage) Merge(filename string, data []byte) error {
	head, _ := adp.Read(filename)
	entire := append(head, data...)

	return adp.Write(filename, entire)
}

// Files returns filename list which is traversing with glob from filesystem.
func (adp *fileStorage) Files(ptn string) ([]string, error) {
	matches, err := filepath.Glob(adp.dsn.Join(ptn))
	if err != nil {
		return []string{}, err
	}

	files := []string{}
	for _, file := range matches {
		files = append(files, fmt.Sprintf("file://%s", file))
	}

	if err != nil {
		logger.Printf("Failed to retrieve list files %s", err)
		return []string{}, err
	}

	return files, nil
}

// gzip will create sitemap file as a gzip.
func (adp *fileStorage) gzip(file *os.File, data []byte) {
	gz := gzip.NewWriter(file)
	gz.Write(data)
	defer gz.Close()
}

// plain will create uncompressed file.
func (adp *fileStorage) plain(file *os.File, data []byte) {
	file.Write(data)
	defer file.Close()
}

// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bleve

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"path"
)

// zipArchiver archives a directory into a zip file.
type zipArchiver struct {
	createDirectory bool
}

func (z *zipArchiver) Suffix() string {
	return ".zip"
}

func (z *zipArchiver) Archive(sourceDirectory, destinationFile string) error {
	f, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	archive := zip.NewWriter(f)
	files, err := ioutil.ReadDir(sourceDirectory)
	if err != nil {
		return err
	}
	for _, file := range files {
		w, err := archive.Create(file.Name())
		if err != nil {
			return err
		}
		b, err := ioutil.ReadFile(path.Join(sourceDirectory, file.Name()))
		if err != nil {
			return err
		}
		if n, err := w.Write(b); err != nil || n != len(b) {
			return err
		}
	}
	return archive.Close()
}

func (z *zipArchiver) Unarchive(b []byte, destinationDirectory string) error {
	rd := bytes.NewReader(b)
	archive, err := zip.NewReader(rd, rd.Size())
	if err != nil {
		return err
	}
	if z.createDirectory {
		if err := os.MkdirAll(destinationDirectory, 0755); err != nil {
			return err
		}
	}

	for _, file := range archive.File {
		r, err := file.Open()
		if err != nil {
			return err
		}
		defer r.Close()
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(path.Join(destinationDirectory, file.Name), b, file.Mode()); err != nil {
			return err
		}
	}
	return nil
}

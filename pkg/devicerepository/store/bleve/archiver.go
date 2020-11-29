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

// archiver handles archiving and unarchiving a directory of files.
type archiver interface {
	// Suffix returns the archive suffix.
	Suffix() string
	// Archive archives a directory into destinationFile.
	Archive(sourceDirectory, destinationFile string) error
	// Unarchive extracts all files from a zip file to destinationDirectory.
	Unarchive(b []byte, destinationDirectory string) error
}

type noopArchiver struct{}

func (*noopArchiver) Suffix() string {
	return ""
}

func (*noopArchiver) Archive(sourceDirectory, destinationFile string) error {
	return nil
}

func (*noopArchiver) Unarchive(b []byte, destinationDirectory string) error {
	return nil
}

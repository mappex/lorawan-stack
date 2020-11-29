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
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
)

// Config represents configuration for the store.
type Config struct {
	Store store.Store `name:"-"`

	WorkingDirectory string         `name:"working-directory" description:"Use this directory for the index files"`
	Refresh          string         `name:"refresh" description:"When to refresh the index files (never|startup|timer)"`
	RefreshInterval  *time.Duration `name:"refresh-interval" description:"Interval for updating index files"`

	Static    map[string][]byte `name:"-"`
	Directory string            `name:"directory" description:"Retrieve index files from the filesystem"`
	URL       string            `name:"url" description:"Retrieve index files from a web server"`
}

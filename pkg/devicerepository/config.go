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

package devicerepository

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store/bleve"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
)

// IndexConfig represents configuration for the DeviceRepository index files.
type IndexConfig struct {
	Type  string       `name:"type" description:"Service to use for indexing (bleve)"`
	Bleve bleve.Config `name:"bleve"`
}

// Config represents the DeviceRepository configuration.
type Config struct {
	Store store.Store `name:"-"`

	Static    map[string][]byte `name:"-"`
	Directory string            `name:"directory" description:"Retrieve devices from the filesystem"`
	URL       string            `name:"url" description:"Retrieve devices from a web server"`

	AssetsBaseURL string `name:"assets-base-url" description:"The base URL for assets"`
	RequireAuth   bool   `name:"require-auth" description:"Require authentication for the device repository"`

	Index IndexConfig `name:"index"`
}

var (
	errUnknownIndexType = errors.DefineUnimplemented("unknown_index_type", "unknown index type `{type}`")
)

// NewStore creates a new Store for end devices.
func (c Config) NewStore(ctx context.Context) (store.Store, error) {
	if c.Store != nil {
		return c.Store, nil
	}
	var fetcher fetch.Interface
	switch {
	case c.Static != nil:
		fetcher = fetch.NewMemFetcher(c.Static)
	case c.Directory != "":
		fetcher = fetch.FromFilesystem(c.Directory)
	case c.URL != "":
		var err error
		fetcher, err = fetch.FromHTTP(c.URL, true)
		if err != nil {
			return nil, err
		}
	default:
		return &store.NoopStore{}, nil
	}
	s, err := store.NewRemoteStore(fetcher)
	if err != nil {
		return nil, err
	}
	// Wrap with indexed store if configureds
	switch indexConfig := c.Index; indexConfig.Type {
	case "":
		return s, err
	case "bleve":
		indexConfig.Bleve.Store = s
		return indexConfig.Bleve.NewStore(ctx)
	default:
		return nil, errUnknownIndexType.WithAttributes("type", indexConfig.Type)
	}
}

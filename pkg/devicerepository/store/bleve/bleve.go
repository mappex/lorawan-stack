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
	"context"
	"path"
	"sync"

	"github.com/blevesearch/bleve"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
)

// bleveStore wraps a store.Store adding support for searching/sorting results using a bleve index.
type bleveStore struct {
	ctx context.Context

	store        store.Store
	archiver     archiver
	indexFetcher fetch.Interface

	brandsIndex   bleve.Index
	brandsIndexMu sync.RWMutex

	modelsIndex   bleve.Index
	modelsIndexMu sync.RWMutex

	workingDirectory string
}

func (bl *bleveStore) fetchIndex(which string) error {
	log.FromContext(bl.ctx).WithField("index", which).Debug("Fetching index")
	b, err := bl.indexFetcher.File(which + bl.archiver.Suffix())
	if err != nil {
		return err
	}
	return bl.archiver.Unarchive(b, path.Join(bl.workingDirectory, which))
}

var (
	errNoWorkingDirectory   = errors.DefineInvalidArgument("no_working_directory", "no working directory specified")
	errNoIndexFetcherConfig = errors.DefineInvalidArgument("no_index_fetcher_config", "no index fetcher configuration specified")
	errInvalidRefreshConfig = errors.DefineInvalidArgument("invalid_refresh_config", "unknown refresh configuration `{refresh}`")
)

// NewStore returns a new bleveStore from configuration.
func (c Config) NewStore(ctx context.Context) (store.Store, error) {
	if c.WorkingDirectory == "" {
		return nil, errNoWorkingDirectory.New()
	}

	bl := &bleveStore{
		ctx:      ctx,
		store:    c.Store,
		archiver: &zipArchiver{createDirectory: true},

		workingDirectory: c.WorkingDirectory,
	}
	switch {
	case c.Static != nil:
		bl.indexFetcher = fetch.NewMemFetcher(c.Static)
	case c.Directory != "":
		bl.indexFetcher = fetch.FromFilesystem(c.Directory)
	case c.URL != "":
		var err error
		bl.indexFetcher, err = fetch.FromHTTP(c.URL, true)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errNoIndexFetcherConfig.New()
	}

	switch c.Refresh {
	case "never":
	case "startup":
		if err := bl.fetchIndex(brandsIndexPath); err != nil {
			return nil, err
		}
		if err := bl.fetchIndex(modelsIndexPath); err != nil {
			return nil, err
		}
	case "timer":
		// TODO: goroutine that retrieves the index files periodically
	default:
		return nil, errInvalidRefreshConfig.WithAttributes("refresh", c.Refresh)
	}
	var err error
	bl.brandsIndexMu.Lock()
	bl.brandsIndex, err = bleve.Open(path.Join(bl.workingDirectory, brandsIndexPath))
	if err != nil {
		return nil, err
	}
	bl.brandsIndexMu.Unlock()
	bl.modelsIndexMu.Lock()
	bl.modelsIndex, err = bleve.Open(path.Join(bl.workingDirectory, modelsIndexPath))
	if err != nil {
		return nil, err
	}
	bl.modelsIndexMu.Unlock()

	go func() {
		select {
		case <-bl.ctx.Done():
			bl.modelsIndex.Close()
			bl.brandsIndex.Close()
		}
	}()

	return bl, nil
}

// NewIndexer creates a new indexer from configuration.
func (c Config) NewIndexer(ctx context.Context, store store.Store) (Indexer, error) {
	return &bleveStore{
		ctx: ctx,

		store:    store,
		archiver: &zipArchiver{createDirectory: true},
	}, nil
}

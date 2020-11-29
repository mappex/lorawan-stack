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
	"fmt"
	"path"

	"github.com/blevesearch/bleve"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// Indexer creates an index for end device brands and models.
type Indexer interface {
	IndexBrands(destinationDirectory string) error
	IndexModels(destinationDirectory string) error
}

type indexableBrand struct {
	BrandPB  string // *ttnpb.EndDeviceBrand marshaled into JSON string
	ModelsPB string // []*ttnpb.EndDeviceModel marshaled into JSON string

	// Stored in separate fields to support ordering search results
	BrandID   string
	BrandName string
}

type indexableModel struct {
	BrandPB string // *ttnpb.EndDeviceBrand marshaled into JSON string
	ModelPB string // *ttnpb.EndDeviceModel marshaled into JSON string

	// Stored in separate fields to support ordering search results
	BrandID   string
	BrandName string
	ModelID   string
	ModelName string
}

const (
	brandsIndexPath = "brandsIndex.bleve"
	modelsIndexPath = "modelsIndex.bleve"
)

// IndexBrands creates a new brands index, meant to be used by bleveStore.GetBrands()
func (bl *bleveStore) IndexBrands(destinationDirectory string) error {
	mapping := bleve.NewIndexMapping()
	indexPath := path.Join(destinationDirectory, brandsIndexPath)
	index, err := bleve.New(indexPath, mapping)
	if err != nil {
		return err
	}

	brands, err := bl.store.GetBrands(store.GetBrandsRequest{
		Paths: ttnpb.EndDeviceBrandFieldPathsNested,
	})
	if err != nil {
		return err
	}

	batch := index.NewBatch()
	for _, brand := range brands.Brands {
		log.FromContext(bl.ctx).WithField("brand_id", brand.BrandID).Debug("Indexing")
		models, err := bl.store.GetModels(store.GetModelsRequest{
			Paths:   ttnpb.EndDeviceModelFieldPathsNested,
			BrandID: brand.BrandID,
		})
		if errors.IsNotFound(err) {
			// Skip vendors without any models
			continue
		} else if err != nil {
			return err
		}
		brandPB, err := jsonpb.TTN().Marshal(brand)
		if err != nil {
			return err
		}
		modelsPB, err := jsonpb.TTN().Marshal(models.Models)
		if err != nil {
			return err
		}
		if err := batch.Index(brand.BrandID, indexableBrand{
			BrandPB:   string(brandPB),
			ModelsPB:  string(modelsPB),
			BrandID:   brand.BrandID,
			BrandName: brand.Name,
		}); err != nil {
			return err
		}
	}
	if err := index.Batch(batch); err != nil {
		return err
	}

	return bl.archiver.Archive(indexPath, indexPath+bl.archiver.Suffix())
}

// IndexModels creates a new models index, meant to be used by bleveStore.GetBrands()
func (bl *bleveStore) IndexModels(destinationDirectory string) error {
	mapping := bleve.NewIndexMapping()
	indexPath := path.Join(destinationDirectory, modelsIndexPath)
	index, err := bleve.New(indexPath, mapping)
	if err != nil {
		return err
	}

	brands, err := bl.store.GetBrands(store.GetBrandsRequest{
		Paths: ttnpb.EndDeviceBrandFieldPathsNested,
	})
	if err != nil {
		return err
	}

	batch := index.NewBatch()
	for _, brand := range brands.Brands {
		log.FromContext(bl.ctx).WithField("brand_id", brand.BrandID).Debug("Indexing")
		brandPB, err := jsonpb.TTN().Marshal(brand)
		if err != nil {
			return err
		}
		models, err := bl.store.GetModels(store.GetModelsRequest{
			Paths:   ttnpb.EndDeviceModelFieldPathsNested,
			BrandID: brand.BrandID,
		})
		if errors.IsNotFound(err) {
			// Skip vendors without any models
			continue
		} else if err != nil {
			return err
		}
		for _, model := range models.Models {
			modelPB, err := jsonpb.TTN().Marshal(model)
			if err != nil {
				return err
			}
			if err := batch.Index(fmt.Sprintf("%s/%s", brand.BrandID, model.ModelID), indexableModel{
				BrandPB:   string(brandPB),
				ModelPB:   string(modelPB),
				BrandID:   brand.BrandID,
				BrandName: brand.Name,
				ModelID:   model.ModelID,
				ModelName: model.Name,
			}); err != nil {
				return err
			}
		}
	}
	if err := index.Batch(batch); err != nil {
		return err
	}

	return bl.archiver.Archive(indexPath, indexPath+bl.archiver.Suffix())
}

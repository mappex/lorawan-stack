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

package commands

import (
	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var (
	errInvalidIndexType = errors.DefineInvalidArgument("invalid_index_type", "invalid index type `{type}`")
	drCommand           = &cobra.Command{
		Use:   "dr",
		Short: "Device Repository commands",
	}
	drCreateIndexCommand = &cobra.Command{
		Use:   "create-index",
		Short: "Create a new index for the Device Repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Info("Creating new index...")
			destination, _ := cmd.Flags().GetString("destination")
			if destination == "" {
				return errMissingFlag.WithAttributes("flag", "destination")
			}

			// TODO: this is not too clean
			switch drConfig := config.DR; drConfig.Index.Type {
			case "bleve":
				// FIXME: this can be a bit cleaner.
				// Clean index configuration so that we create the store correctly.
				drConfig.Index.Type = ""
				store, err := drConfig.NewStore(ctx)
				if err != nil {
					return err
				}
				drConfig.Index.Bleve.Store = store
				indexer, err := drConfig.Index.Bleve.NewIndexer(ctx, store)
				if err != nil {
					return nil
				}

				if err := indexer.IndexBrands(destination); err != nil {
					return err
				}
				if err := indexer.IndexModels(destination); err != nil {
					return err
				}
			default:
				return errInvalidIndexType.WithAttributes("type", drConfig.Index.Type)
			}
			return nil
		},
	}
)

func init() {
	Root.AddCommand(drCommand)

	drCreateIndexCommand.Flags().String("destination", "", "Place to create the new index")
	drCommand.AddCommand(drCreateIndexCommand)
}

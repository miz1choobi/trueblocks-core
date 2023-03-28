// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package chunksPkg

import (
	"context"
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/cache"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/index"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (opts *ChunksOptions) HandleIndex(blockNums []uint64) error {
	ctx, cancel := context.WithCancel(context.Background())
	fetchData := func(modelChan chan types.Modeler[types.RawIndex], errorChan chan error) {
		showIndex := func(walker *index.IndexWalker, path string, first bool) (bool, error) {
			if path != cache.ToBloomPath(path) {
				err := fmt.Errorf("should not happen in showFinalizedStats")
				errorChan <- err
				cancel()
				return false, nil
			}

			path = cache.ToIndexPath(path)
			if !file.FileExists(path) {
				// Bloom files exist, but index files don't. It's okay.
				return true, nil
			}

			header, err := index.ReadChunkHeader(path, true)
			if err != nil {
				return false, err
			}

			rng, err := base.RangeFromFilenameE(path)
			if err != nil {
				errorChan <- err
				cancel()
				return false, nil
			}

			s := types.SimpleIndex{
				Range:           rng,
				Magic:           header.Magic,
				Hash:            base.HexToHash(header.Hash.Hex()),
				AddressCount:    header.AddressCount,
				AppearanceCount: header.AppearanceCount,
				Size:            file.FileSize(path),
			}

			modelChan <- &s
			return true, nil
		}

		walker := index.NewIndexWalker(
			opts.Globals.Chain,
			opts.Globals.TestMode,
			100, /* maxTests */
			showIndex,
		)
		if err := walker.WalkBloomFilters(blockNums); err != nil {
			errorChan <- err
		}
	}

	return output.StreamMany(ctx, fetchData, output.OutputOptions{
		Writer:     opts.Globals.Writer,
		Chain:      opts.Globals.Chain,
		TestMode:   opts.Globals.TestMode,
		NoHeader:   opts.Globals.NoHeader,
		ShowRaw:    opts.Globals.ShowRaw,
		Verbose:    opts.Globals.Verbose,
		LogLevel:   opts.Globals.LogLevel,
		Format:     opts.Globals.Format,
		OutputFn:   opts.Globals.OutputFn,
		Append:     opts.Globals.Append,
		JsonIndent: "  ",
	})
}

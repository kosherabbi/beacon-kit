// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/itsdevbear/bolaris/types/consensus"
)

// FinalizeBeaconBlock finalizes a beacon block by processing the logs,
// deposits,
// and voluntary exits. It also updates the finalized and safe eth1 block hashes
// on the beacon state.
func (s *Service) FinalizeBeaconBlock(
	ctx context.Context,
	blk consensus.ReadOnlyBeaconKitBlock,
) error {
	payload, err := blk.ExecutionPayload()
	if err != nil {
		return err
	}

	// TODO: PROCESS LOGS HERE
	// TODO: PROCESS DEPOSITS HERE
	// TODO: PROCESS VOLUNTARY EXITS HERE

	// TEMPORARY, needs to be handled better, this is a hack.
	if err = s.sendFCU(
		ctx, common.Hash(payload.GetBlockHash()), blk.GetSlot()+1,
	); err != nil {
		s.Logger().
			Error("failed to notify forkchoice update in preblocker", "error", err)
	}

	eth1BlockHash := common.Hash(payload.GetBlockHash())
	state := s.BeaconState(ctx)
	state.SetFinalizedEth1BlockHash(eth1BlockHash)
	state.SetSafeEth1BlockHash(eth1BlockHash)
	state.SetLastValidHead(eth1BlockHash)

	return nil
}

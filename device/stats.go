/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2020 WireGuard LLC. All Rights Reserved.
 */

package device

import (
	"sync/atomic"
	"time"
)

// PeerStats are connection statistics for a given Peer.
type PeerStats struct {
	RxBytes                uint64
	TxBytes                uint64
	LastHandshakeInitiated time.Time
}

// PeerStats returns statistics for the peer with public key pk.
// If there is no peer with that key, PeerStats returns nil.
func (device *Device) PeerStats(pk NoisePublicKey) *PeerStats {
	device.peers.RLock()
	peer := device.peers.keyMap[pk]
	device.peers.RUnlock()

	if peer == nil {
		return nil
	}

	peer.RLock()
	defer peer.RUnlock()
	return &PeerStats{
		RxBytes:                atomic.LoadUint64(&peer.stats.rxBytes),
		TxBytes:                atomic.LoadUint64(&peer.stats.txBytes),
		LastHandshakeInitiated: time.Unix(0, atomic.LoadInt64(&peer.stats.lastHandshakeNano)),
	}
}

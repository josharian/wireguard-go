// +build !android,!ios

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2020 WireGuard LLC. All Rights Reserved.
 */

package device

const (
	QueueOutboundSize          = 1024
	QueueInboundSize           = 1024
	QueueHandshakeSize         = 1024
	MaxSegmentSize             = (1 << 16) - 1 // largest possible UDP datagram
	PreallocatedBuffersPerPool = 1024          // Disable and allow for infinite memory growth
)

// TODO: set  prealloc low, so that we  can  detect leaks
//  tricky to find the  right low  number

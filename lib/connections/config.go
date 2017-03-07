// Copyright (C) 2017 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package connections

import (
	"io/ioutil"
	"time"

	"github.com/hashicorp/yamux"
)

const (
	tcpPriority   = 10
	kcpPriority   = 50
	relayPriority = 200

	// KCP filter priorities
	kcpNoFilterPriority           = 100
	kcpConversationFilterPriority = 20
	kcpStunFilterPriority         = 10

	// KCP SetNoDelay options

	// 0 - disabled (default)
	// 1 - enabled
	kcpNoDelay        = 0
	kcpUpdateInterval = 100 // ms (default)
	// 0 - disable (default)
	// 1 - enabled
	kcpFastResend = 0
	// 0 - enabled (default)
	// 1 - disabled
	kcpCongestionControl = 0

	// KCP window sizes
	kcpSendWindowSize    = 128
	kcpReceiveWindowSize = 128
)

var (
	yamuxConfig = &yamux.Config{
		AcceptBacklog:          256,
		EnableKeepAlive:        true,
		KeepAliveInterval:      30 * time.Second,
		ConnectionWriteTimeout: 10 * time.Second,
		MaxStreamWindowSize:    256 * 1024,
		LogOutput:              ioutil.Discard,
	}
)
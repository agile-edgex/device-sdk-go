// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package correlation

import (
	"context"

	"github.com/agile-edgex/go-mod-core-contracts/v3/common"
)

func IdFromContext(ctx context.Context) string {
	hdr, ok := ctx.Value(common.CorrelationHeader).(string)
	if !ok {
		hdr = ""
	}
	return hdr
}

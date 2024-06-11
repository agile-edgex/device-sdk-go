// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package transformer

import (
	"math"

	"github.com/agile-edgex/go-mod-core-contracts/v3/common"
	"github.com/agile-edgex/go-mod-core-contracts/v3/errors"

	"github.com/agile-edgex/device-sdk-go/v3/pkg/models"
)

func isNaN(cv *models.CommandValue) (bool, errors.EdgeX) {
	switch cv.Type {
	case common.ValueTypeFloat32:
		v, err := cv.Float32Value()
		if err != nil {
			return false, errors.NewCommonEdgeXWrapper(err)
		}
		if math.IsNaN(float64(v)) {
			return true, nil
		}
	case common.ValueTypeFloat64:
		v, err := cv.Float64Value()
		if err != nil {
			return false, errors.NewCommonEdgeXWrapper(err)
		}
		if math.IsNaN(v) {
			return true, nil
		}
	}
	return false, nil
}

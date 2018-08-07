// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package docker

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
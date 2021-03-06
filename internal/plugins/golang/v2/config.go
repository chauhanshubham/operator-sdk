// Copyright 2020 The Operator-SDK Authors
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

package v2

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
)

// Config configures this plugin, and is saved in the project config file.
// Deprecated: use scorecard and manifests plugin configs directly.
type Config struct{}

// hasPluginConfig returns true if cfg.Layout contains an exact match for this plugin's key.
func hasPluginConfig(cfg config.Config) bool {
	info := Config{}
	return cfg.DecodePluginConfig(pluginConfigKey, &info) == nil
}

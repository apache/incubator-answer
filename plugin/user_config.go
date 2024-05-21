/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package plugin

type UserConfig interface {
	Base

	// UserConfigFields returns the list of config fields
	UserConfigFields() []ConfigField
	// UserConfigReceiver receives the config data, it calls when the config is saved or initialized.
	// We recommend to unmarshal the data to a struct, and then use the struct to do something.
	// The config is encoded in JSON format.
	// It depends on the definition of ConfigFields.
	UserConfigReceiver(userID string, config []byte) error
}

var (
	// CallUserConfig is a function that calls all registered config plugins
	CallUserConfig,
	registerUserConfig = MakePlugin[UserConfig](false)
	getPluginUserConfigFn func(userID, pluginSlugName string) []byte
)

// GetPluginUserConfig returns the user config of the given user id
func GetPluginUserConfig(userID, pluginSlugName string) []byte {
	if getPluginUserConfigFn != nil {
		return getPluginUserConfigFn(userID, pluginSlugName)
	}
	return nil
}

// RegisterGetPluginUserConfigFunc registers a function to get the user config of the given user id
func RegisterGetPluginUserConfigFunc(fn func(userID, pluginSlugName string) []byte) {
	getPluginUserConfigFn = fn
}

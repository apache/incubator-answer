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

package entity

// PluginConfig plugin config
type PluginConfig struct {
	ID             int    `xorm:"not null pk autoincr INT(11) id"`
	PluginSlugName string `xorm:"unique VARCHAR(128) plugin_slug_name"`
	Value          string `xorm:"TEXT value"`
}

// TableName config table name
func (PluginConfig) TableName() string {
	return "plugin_config"
}

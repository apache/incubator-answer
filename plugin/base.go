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

// Info presents the plugin information
type Info struct {
	Name        Translator
	SlugName    string
	Description Translator
	Author      string
	Version     string
	Link        string
}

// Base is the base plugin
type Base interface {
	// Info returns the plugin information
	Info() Info
}

var (
	// CallBase is a function that calls all registered base plugins
	CallBase,
	registerBase = MakePlugin[Base](true)
)

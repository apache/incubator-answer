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

package converter

func ArrayNotInArray(original []string, search []string) []string {
	var result []string
	originalMap := make(map[string]bool)
	for _, v := range original {
		originalMap[v] = true
	}
	for _, v := range search {
		if _, ok := originalMap[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

func UniqueArray[T comparable](input []T) []T {
	result := make([]T, 0, len(input))
	seen := make(map[T]bool, len(input))
	for _, element := range input {
		if !seen[element] {
			result = append(result, element)
			seen[element] = true
		}
	}
	return result
}

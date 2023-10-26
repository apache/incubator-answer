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

package checker

import (
	"golang.org/x/image/webp"
	"image"
	_ "image/gif" // use init to support decode jpeg,jpg,png,gif
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strings"
)

// IsSupportedImageFile currently answers support image type is
// `image/jpeg, image/jpg, image/png, image/gif, image/webp`
func IsSupportedImageFile(file io.Reader, ext string) bool {
	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	var err error
	switch ext {
	case "jpg", "jpeg", "png", "gif": // only allow for `image/jpeg,image/jpg,image/png, image/gif`
		_, _, err = image.Decode(file)
	case "ico":
		// TODO: There is currently no good Golang library to parse whether the image is in ico format.
		return true
	case "webp":
		_, err = webp.Decode(file)
	default:
		return false
	}
	return err == nil
}

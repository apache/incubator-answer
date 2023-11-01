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

import (
	"github.com/gin-gonic/gin"
)

type Agent interface {
	Base
	RegisterUnAuthRouter(r *gin.RouterGroup)
	RegisterAuthUserRouter(r *gin.RouterGroup)
	RegisterAuthAdminRouter(r *gin.RouterGroup)
}

var (
	CallAgent,
	registerAgent = MakePlugin[Agent](true)
	siteURLFn func() string
)

// SiteURL The site url is the domain address of the current site. e.g. http://localhost:8080
// When some Agent plugins want to redirect to the origin site, it can use this function to get the site url.
func SiteURL() string {
	if siteURLFn != nil {
		return siteURLFn()
	}
	return ""
}

// RegisterGetSiteURLFunc Register a function to get the site url.
func RegisterGetSiteURLFunc(fn func() string) {
	siteURLFn = fn
}

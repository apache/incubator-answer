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

package router

import (
	"github.com/apache/incubator-answer/internal/service/service_config"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
)

// StaticRouter static api router
type StaticRouter struct {
	serviceConfig *service_config.ServiceConfig
}

// NewStaticRouter new static api router
func NewStaticRouter(serviceConfig *service_config.ServiceConfig) *StaticRouter {
	return &StaticRouter{
		serviceConfig: serviceConfig,
	}
}

// RegisterStaticRouter register static api router
func (a *StaticRouter) RegisterStaticRouter(r *gin.RouterGroup) {
	r.Static("/uploads", a.serviceConfig.UploadPath)

	r.GET("/download/*filepath", func(c *gin.Context) {
		// The filePath such as /download/hash/123.png
		filePath := c.Param("filepath")
		// The download filename is 123.png
		downloadFilename := filepath.Base(filePath)

		// After trimming, the downloadLink is /uploads/hash
		downloadLink := strings.TrimSuffix(filePath, "/"+downloadFilename)
		// After add the extension, the downloadLink is /uploads/hash.png
		downloadLink += filepath.Ext(downloadFilename)

		downloadLink = filepath.Join(a.serviceConfig.UploadPath, downloadLink)
		c.FileAttachment(downloadLink, downloadFilename)
	})
}

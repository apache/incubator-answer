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

package controller

import (
	"github.com/apache/incubator-answer/internal/base/handler"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/plugin"
	"github.com/gin-gonic/gin"
)

type EmbedController struct {
}

func NewEmbedController() *EmbedController {
	return &EmbedController{}
}

// GetEmbedConfig godoc
// @Summary GetEmbedConfig
// @Description GetEmbedConfig
// @Tags PluginEmbed
// @Accept json
// @Produce json
// @Router /answer/api/v1/embed/config [get]
// @Success 200 {object} handler.RespBody{data=[]schema.GetEmbedOptionResp}
func (c *EmbedController) GetEmbedConfig(ctx *gin.Context) {
	resp := make([]*schema.GetEmbedOptionResp, 0)
	var slugName string

	_ = plugin.CallEmbed(func(base plugin.Embed) error {
		slugName = base.Info().SlugName
		return nil
	})

	_ = plugin.CallConfig(func(fn plugin.Config) error {
		if fn.Info().SlugName == slugName {
			for _, field := range fn.ConfigFields() {
				resp = append(resp, &schema.GetEmbedOptionResp{
					Platform: field.Name,
					Enable:   field.Value.(bool),
				})
			}
			return nil
		}
		return nil
	})
	handler.HandleResponse(ctx, nil, resp)
}

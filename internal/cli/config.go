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

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/incubator-answer/internal/base/constant"
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/entity"
	"xorm.io/xorm"
)

type ConfigField struct {
	AllowPasswordLogin bool `json:"allow_password_login"`
}

// SetDefaultConfig set default config
func SetDefaultConfig(dbConf *data.Database, cacheConf *data.CacheConf, field *ConfigField) error {
	db, err := data.NewDB(false, dbConf)
	if err != nil {
		return err
	}
	defer db.Close()

	cache, cacheCleanup, err := data.NewCache(cacheConf)
	if err != nil {
		fmt.Println("new cache failed")
	}
	defer func() {
		if cache != nil {
			cache.Flush(context.Background())
			cacheCleanup()
		}
	}()

	if field.AllowPasswordLogin {
		return defaultLoginConfig(db)
	}

	return nil
}

func defaultLoginConfig(x *xorm.Engine) (err error) {
	fmt.Println("set default login config")

	loginSiteInfo := &entity.SiteInfo{
		Type: constant.SiteTypeLogin,
	}
	exist, err := x.Get(loginSiteInfo)
	if err != nil {
		return fmt.Errorf("get config failed: %w", err)
	}
	if exist {
		var content map[string]any
		_ = json.Unmarshal([]byte(loginSiteInfo.Content), &content)
		content["allow_password_login"] = true
		dataByte, _ := json.Marshal(content)
		loginSiteInfo.Content = string(dataByte)
		_, err = x.ID(loginSiteInfo.ID).Cols("content").Update(loginSiteInfo)
		if err != nil {
			return fmt.Errorf("update site info failed: %w", err)
		}
	}
	return nil
}

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

package schema

import "time"

var AppStartTime time.Time

const (
	DashboardCacheKey  = "answer:dashboard"
	DashboardCacheTime = 60 * time.Minute
)

type DashboardInfo struct {
	QuestionCount         int64                `json:"question_count"`
	ResolvedCount         int64                `json:"resolved_count"`
	ResolvedRate          string               `json:"resolved_rate"`
	UnansweredCount       int64                `json:"unanswered_count"`
	UnansweredRate        string               `json:"unanswered_rate"`
	AnswerCount           int64                `json:"answer_count"`
	CommentCount          int64                `json:"comment_count"`
	VoteCount             int64                `json:"vote_count"`
	UserCount             int64                `json:"user_count"`
	ReportCount           int64                `json:"report_count"`
	UploadingFiles        bool                 `json:"uploading_files"`
	SMTP                  string               `json:"smtp"`
	HTTPS                 bool                 `json:"https"`
	TimeZone              string               `json:"time_zone"`
	OccupyingStorageSpace string               `json:"occupying_storage_space"`
	AppStartTime          string               `json:"app_start_time"`
	VersionInfo           DashboardInfoVersion `json:"version_info"`
	LoginRequired         bool                 `json:"login_required"`
	GoVersion             string               `json:"go_version"`
	DatabaseVersion       string               `json:"database_version"`
	DatabaseSize          string               `json:"database_size"`
}

type DashboardInfoVersion struct {
	Version       string `json:"version"`
	Revision      string `json:"revision"`
	RemoteVersion string `json:"remote_version"`
}

type RemoteVersion struct {
	Release struct {
		Version string `json:"version"`
		URL     string `json:"url"`
	} `json:"release"`
}

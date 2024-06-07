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

type Paginator struct {
	Pages      []int
	Totalpages int
	Prevpage   int
	Nextpage   int
	Currpage   int
}

type QAPageJsonLD struct {
	Context    string `json:"@context"`
	Type       string `json:"@type"`
	MainEntity struct {
		Type        string    `json:"@type"`
		Name        string    `json:"name"`
		Text        string    `json:"text"`
		AnswerCount int       `json:"answerCount"`
		UpvoteCount int       `json:"upvoteCount"`
		DateCreated time.Time `json:"dateCreated"`
		Author      struct {
			URL  string `json:"url"`
			Type string `json:"@type"`
			Name string `json:"name"`
		} `json:"author"`
		AcceptedAnswer  *AcceptedAnswerItem    `json:"acceptedAnswer,omitempty"`
		SuggestedAnswer []*SuggestedAnswerItem `json:"suggestedAnswer"`
	} `json:"mainEntity"`
}

type AcceptedAnswerItem struct {
	Type        string    `json:"@type"`
	Text        string    `json:"text"`
	DateCreated time.Time `json:"dateCreated"`
	UpvoteCount int       `json:"upvoteCount"`
	URL         string    `json:"url"`
	Author      struct {
		URL  string `json:"url"`
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"author"`
}

type SuggestedAnswerItem struct {
	Type        string    `json:"@type"`
	Text        string    `json:"text"`
	DateCreated time.Time `json:"dateCreated"`
	UpvoteCount int       `json:"upvoteCount"`
	URL         string    `json:"url"`
	Author      struct {
		URL  string `json:"url"`
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"author"`
}

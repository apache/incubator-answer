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

type UploadSource string

const (
	UserAvatar         UploadSource = "user_avatar"
	UserPost           UploadSource = "user_post"
	UserPostAttachment UploadSource = "user_post_attachment"
	AdminBranding      UploadSource = "admin_branding"
)

var (
	DefaultFileTypeCheckMapping = map[UploadSource]map[string]bool{
		UserAvatar: {
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".webp": true,
		},
		UserPost: {
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
			".webp": true,
		},
		AdminBranding: {
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".ico":  true,
		},
	}
)

type UploadFileCondition struct {
	// Source is the source of the file
	Source UploadSource
	// MaxImageSize is the maximum size of the image in MB
	MaxImageSize int
	// MaxAttachmentSize is the maximum size of the attachment in MB
	MaxAttachmentSize int
	// MaxImageMegapixel is the maximum megapixel of the image
	MaxImageMegapixel int
	// AuthorizedImageExtensions is the list of authorized image extensions
	AuthorizedImageExtensions []string
	// AuthorizedAttachmentExtensions is the list of authorized attachment extensions
	AuthorizedAttachmentExtensions []string
}

type UploadFileResponse struct {
	// FullURL is the URL that can be used to access the file
	FullURL string
	// OriginalError is the error returned by the storage plugin. It is used for debugging.
	OriginalError error
	// DisplayErrorMsg is the error message that will be displayed to the user.
	DisplayErrorMsg Translator
}

type Storage interface {
	Base

	// UploadFile uploads a file to storage.
	// The file is in the Form of the ctx and the key is "file"
	UploadFile(ctx *GinContext, condition UploadFileCondition) UploadFileResponse
}

var (
	// CallStorage is a function that calls all registered storage
	CallStorage,
	registerStorage = MakePlugin[Storage](false)
)

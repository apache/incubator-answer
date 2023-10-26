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

package uid

import (
	"strconv"

	"github.com/segmentfault/pacman/utils"
)

const salt = int64(100)

// NumToShortID num to string
func NumToShortID(id int64) string {
	sid := strconv.FormatInt(id, 10)
	if len(sid) < 17 {
		return ""
	}
	sTypeCode := sid[1:4]
	sid = sid[4:int32(len(sid))]
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return ""
	}
	typeCode, err := strconv.ParseInt(sTypeCode, 10, 64)
	if err != nil {
		return ""
	}
	code := utils.EnShortID(id, salt)
	tcode := utils.EnShortID(typeCode, salt)
	return tcode + code
}

// ShortIDToNum string to num
func ShortIDToNum(code string) int64 {
	if len(code) < 2 {
		return 0
	}
	scodeType := code[0:2]
	code = code[2:int32(len(code))]

	id := utils.DeShortID(code, salt)
	codeType := utils.DeShortID(scodeType, salt)
	return 10000000000000000 + codeType*10000000000000 + id
}

func EnShortID(id string) string {
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return id
	}
	return NumToShortID(num)
}

func DeShortID(sid string) string {
	num, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return strconv.FormatInt(ShortIDToNum(sid), 10)
	}
	if num < 10000000000000000 {
		return strconv.FormatInt(ShortIDToNum(sid), 10)
	}
	return sid
}

func IsShortID(id string) bool {
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return true
	}
	if num < 10000000000000000 {
		return true
	}
	return false
}

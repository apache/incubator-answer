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

import { DRAFT_TIMESIGH_STORAGE_KEY as timeSign } from '@/common/constants';

const store = {
  storage: localStorage || window.localStorage,
  set(key: string, value, time?: number): void {
    const t = time || Date.now() + 1000 * 60 * 60 * 24 * 7; // default 7 days
    try {
      this.storage.setItem(key, `${t}${timeSign}${JSON.stringify(value)}`);
    } catch {
      // ignore
      console.error('set storage error: the key is', key);
    }
  },
  get(key: string): any {
    const timeSignLen = timeSign.length;
    let index = 0;
    let time = 0;
    let res;
    try {
      res = this.storage.getItem(key);
    } catch {
      console.error('get storage error: the key is', key);
    }
    if (res) {
      index = res.indexOf(timeSign);
      time = +res.slice(0, index);
      if (time > new Date().getTime()) {
        res = res.slice(index + timeSignLen);
        try {
          res = JSON.parse(res);
        } catch {
          // ignore
        }
      } else {
        // timeout remove storage
        res = null;
        this.storage.removeItem(key);
      }
      return res;
    }

    return res;
  },
  remove(key: string, callback?: () => void): void {
    this.storage.removeItem(key);
    callback?.();
  },
};

export default store;

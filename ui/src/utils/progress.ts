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

class Progress {
  private timestamp;

  private duration: number;

  private progress: number;

  private delta: number;

  private isLoop: boolean;

  constructor(
    param = {
      timestamp: null,
      duration: 1000,
      delta: 0,
      isLoop: true,
    },
  ) {
    this.timestamp = null;
    this.duration = param.duration || Progress.CONST.DURATION;
    this.progress = 0;
    this.delta = 0;
    this.isLoop = !!param.isLoop;

    this.reset();
  }

  static get CONST() {
    return {
      DURATION: 1000,
    };
  }

  reset() {
    this.timestamp = null;
  }

  start(now) {
    this.timestamp = now;
  }

  tick(now) {
    if (this.timestamp) {
      this.delta = now - this.timestamp;
      this.progress = Math.min(this.delta / this.duration, 1);

      if (this.progress >= 1 && this.isLoop) {
        this.start(now);
      }

      return this.progress;
    }
    return 0;
  }
}

export default Progress;

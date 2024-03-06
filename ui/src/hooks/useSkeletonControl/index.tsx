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

import { useCallback, useRef, useState } from 'react';

import { SKELETON_SHOW_MIN_TIME } from '@/common/constants';

interface IRef {
  startTime?: number;
  timer?: NodeJS.Timeout;
  isSkeletonShow?: boolean;
}

const useSkeletonControl = () => {
  const [isSkeletonShow, setIsSkeletonShow] = useState(false);
  const controlRef = useRef<IRef>({});
  const openSkeleton = () => {
    if (!controlRef.current.timer) {
      controlRef.current.timer = setTimeout(() => {
        setIsSkeletonShow(true);
        controlRef.current.startTime = Date.now();
      }, 1000);
    }
  };

  const closeSkeleton = useCallback(() => {
    clearTimeout(controlRef.current.timer);
    controlRef.current.timer = undefined;
    if (isSkeletonShow && controlRef.current.startTime) {
      const delayTime =
        Date.now() - controlRef.current.startTime + SKELETON_SHOW_MIN_TIME;
      console.log(delayTime);
      setTimeout(
        () => {
          setIsSkeletonShow(false);
        },
        delayTime <= 0 ? 0 : delayTime,
      );
    }
  }, [isSkeletonShow]);

  return { isSkeletonShow, openSkeleton, closeSkeleton };
};

export default useSkeletonControl;

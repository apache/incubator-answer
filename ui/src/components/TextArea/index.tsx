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

import { FC, useRef, useEffect, memo } from 'react';
import { FormControl, FormControlProps } from 'react-bootstrap';

const TextArea: FC<
  FormControlProps & { rows?: number; autoFocus?: boolean }
> = ({
  value,
  onChange,
  size,
  rows = 1,
  autoFocus = true,
  isInvalid,
  ...rest
}) => {
  const ref = useRef<HTMLTextAreaElement>(null);

  const autoGrow = () => {
    if (ref.current) {
      ref.current.style.height = 'auto';
      ref.current.style.height = `${ref.current.scrollHeight}px`;
    }
  };

  useEffect(() => {
    if (ref.current && value) {
      autoGrow();
    }
  }, [ref, value]);

  return (
    <FormControl
      as="textarea"
      className="resize-none font-monospace"
      rows={rows}
      size={size}
      value={value}
      onChange={onChange}
      autoFocus={autoFocus}
      ref={ref}
      onInput={autoGrow}
      isInvalid={isInvalid}
      {...rest}
    />
  );
};
export default memo(TextArea);

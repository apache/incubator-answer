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

/* eslint-disable @typescript-eslint/no-use-before-define */
import * as React from 'react';

import ReactDOM from 'react-dom/client';

import Modal from './Modal';
import type { Props } from './Modal';

const div = document.createElement('div');

const root = ReactDOM.createRoot(div);

export interface Config extends Props {
  content: string;
}

const Index = ({
  title = '',
  confirmText = '',
  content,
  onCancel: onClose,
  onConfirm,
  cancelBtnVariant = 'link',
  confirmBtnVariant = 'primary',
  ...props
}: Config) => {
  const onCancel = () => {
    if (typeof onClose === 'function') {
      onClose();
    }
    render({ visible: false });
    div.remove();
  };
  const onOk = (e) => {
    if (typeof onConfirm === 'function') {
      onConfirm(e);
    }
    onCancel();
  };
  function render({ visible }: { visible: boolean }) {
    root.render(
      <Modal
        visible={visible}
        title={title}
        centered={false}
        onCancel={onCancel}
        onConfirm={onOk}
        confirmText={confirmText}
        cancelBtnVariant={cancelBtnVariant}
        confirmBtnVariant={confirmBtnVariant}
        {...props}>
        <p dangerouslySetInnerHTML={{ __html: content }} />
      </Modal>,
    );
  }
  render({ visible: true });
};

export default Index;

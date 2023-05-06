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

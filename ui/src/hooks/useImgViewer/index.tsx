import { useLayoutEffect, useState, MouseEvent } from 'react';
import { Modal } from 'react-bootstrap';

import ReactDOM from 'react-dom/client';

const div = document.createElement('div');
const root = ReactDOM.createRoot(div);

const useImgViewer = () => {
  const [visible, setVisible] = useState(false);
  const [imgSrc, setImgSrc] = useState('');
  const onClose = () => {
    setVisible(false);
    setImgSrc('');
  };

  const checkIfInLink = (target) => {
    let ret = false;
    let el = target.parentElement;
    while (el) {
      if (el.nodeName.toLowerCase() === 'a') {
        ret = true;
        break;
      }
      el = el.parentElement;
    }
    return ret;
  };

  const checkClickForImgView = (evt: MouseEvent<HTMLElement>) => {
    const { target } = evt;
    // @ts-ignore
    if (target.nodeName.toLowerCase() !== 'img') {
      return;
    }
    const img = target as HTMLImageElement;
    if (!img.naturalWidth || !img.naturalHeight) {
      return;
    }
    const src = img.currentSrc || img.src;
    if (src && checkIfInLink(img) === false) {
      setImgSrc(src);
      setVisible(true);
    }
  };

  useLayoutEffect(() => {
    root.render(
      <Modal
        show={visible}
        fullscreen
        centered
        scrollable
        contentClassName="bg-transparent"
        onHide={onClose}>
        <Modal.Body>
          {/* eslint-disable-next-line jsx-a11y/click-events-have-key-events,jsx-a11y/no-noninteractive-element-interactions */}
          <img
            className="cursor-zoom-out img-fluid position-absolute top-50 start-50 translate-middle"
            src={imgSrc}
            alt={imgSrc}
            onClick={onClose}
          />
        </Modal.Body>
      </Modal>,
    );
  });
  return {
    onClose,
    checkClickForImgView,
  };
};

export default useImgViewer;

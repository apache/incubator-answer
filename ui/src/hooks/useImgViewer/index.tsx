import { useLayoutEffect, useState, MouseEvent, useEffect } from 'react';
import { Modal } from 'react-bootstrap';
import { useLocation } from 'react-router-dom';

import ReactDOM from 'react-dom/client';

const div = document.createElement('div');
const root = ReactDOM.createRoot(div);

const useImgViewer = () => {
  const location = useLocation();
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
      img.classList.add('broken');
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
        <Modal.Body onClick={onClose}>
          <img
            className="cursor-zoom-out img-fluid position-absolute top-50 start-50 translate-middle"
            src={imgSrc}
            alt={imgSrc}
          />
        </Modal.Body>
      </Modal>,
    );
  });
  useEffect(() => {
    onClose();
  }, [location]);
  return {
    onClose,
    checkClickForImgView,
  };
};

export default useImgViewer;

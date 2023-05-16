import { FC, MouseEvent, ReactNode, useEffect, useState } from 'react';
import { Modal } from 'react-bootstrap';

import './index.css';
import classnames from 'classnames';

const Index: FC<{
  children: ReactNode;
  className?: classnames.Argument;
}> = ({ children, className }) => {
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

  useEffect(() => {
    return () => {
      onClose();
    };
  }, []);

  return (
    // eslint-disable-next-line jsx-a11y/click-events-have-key-events
    <div
      className={classnames('img-viewer', className)}
      onClick={checkClickForImgView}>
      {children}
      <Modal
        show={visible}
        fullscreen
        centered
        scrollable
        contentClassName="bg-transparent"
        onHide={onClose}>
        <Modal.Body onClick={onClose} className="img-viewer p-0 d-flex">
          {/* eslint-disable-next-line jsx-a11y/click-events-have-key-events,jsx-a11y/no-noninteractive-element-interactions */}
          <img
            className="cursor-zoom-out img-fluid m-auto"
            onClick={(evt) => {
              evt.stopPropagation();
              onClose();
            }}
            src={imgSrc}
            alt={imgSrc}
          />
        </Modal.Body>
      </Modal>
    </div>
  );
};

export default Index;

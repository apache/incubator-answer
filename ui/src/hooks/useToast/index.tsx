import { useLayoutEffect, useState } from 'react';
import { Toast } from 'react-bootstrap';

import ReactDOM from 'react-dom/client';

const div = document.createElement('div');
div.style.position = 'fixed';
div.style.top = '90px';
div.style.left = '0';
div.style.right = '0';
div.style.margin = 'auto';
div.style.zIndex = '5';

const root = ReactDOM.createRoot(div);

interface Params {
  /** main content */
  msg: string;
  /** theme color */
  variant?: 'warning' | 'success' | 'danger';
}

const useToast = () => {
  const [show, setShow] = useState(false);
  const [data, setData] = useState<Params>({
    msg: '',
    variant: 'warning',
  });

  const onClose = () => {
    setShow(false);
  };

  const onShow = (t: Params) => {
    setData(t);
    setShow(true);
  };
  useLayoutEffect(() => {
    const parent = document.querySelector('.page-wrap');
    parent?.appendChild(div);

    root.render(
      <div className="d-flex justify-content-center">
        <Toast
          className="align-items-center border-0"
          delay={5000}
          bg={data.variant || 'warning'}
          show={show}
          autohide
          onClose={onClose}>
          <div className="d-flex">
            <Toast.Body
              dangerouslySetInnerHTML={{ __html: data.msg }}
              className={`${data.variant !== 'warning' ? 'text-white' : ''}`}
            />
            <button
              className={`btn-close me-2 m-auto ${
                data.variant !== 'warning' ? 'btn-close-white' : ''
              }`}
              onClick={onClose}
              data-bs-dismiss="toast"
              aria-label="Close"
            />
          </div>
        </Toast>
      </div>,
    );
  }, [show, data]);
  return {
    onShow,
  };
};

export default useToast;

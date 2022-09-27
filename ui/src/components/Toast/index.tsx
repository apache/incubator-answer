import { FC, memo } from 'react';
import { Toast } from 'react-bootstrap';

interface IProps {
  /** main content */
  msg: string;
  /** theme color */
  variant?: 'warning' | 'success' | 'danger';
  /** callback click close */
  onClose: () => void;
}

const Index: FC<IProps> = ({ msg, variant = 'warning', onClose }) => {
  return (
    <div
      className="d-flex justify-content-center"
      style={{
        position: 'fixed',
        top: '90px',
        left: 0,
        right: 0,
        margin: 'auto',
        zIndex: 5,
      }}>
      <Toast
        className="align-items-center border-0"
        delay={5000}
        bg={variant}
        show={Boolean(msg)}
        autohide
        onClose={onClose}>
        <div className="d-flex">
          <Toast.Body
            dangerouslySetInnerHTML={{ __html: msg }}
            className={`${variant !== 'warning' ? 'text-white' : ''}`}
          />
          <button
            className={`btn-close me-2 m-auto ${
              variant !== 'warning' ? 'btn-close-white' : ''
            }`}
            onClick={onClose}
            data-bs-dismiss="toast"
            aria-label="Close"
          />
        </div>
      </Toast>
    </div>
  );
};

export default memo(Index);

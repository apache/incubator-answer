import React, { FC } from 'react';
import { Button, Modal } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

export interface Props {
  /** header title */
  title?: string;
  children?: React.ReactNode;
  /** visible */
  visible?: boolean;
  centered?: boolean;
  onCancel?: () => void;
  onConfirm?: (event: any) => void;
  cancelText?: string;
  showCancel?: boolean;
  cancelBtnVariant?: string;
  confirmText?: string;
  showConfirm?: boolean;
  confirmBtnDisabled?: boolean;
  confirmBtnVariant?: string;
  /** body style */
  bodyClass?: string;
  scrollable?: boolean;
  className?: string;
}
const Index: FC<Props> = ({
  title = 'title',
  visible = false,
  centered = true,
  onCancel,
  children,
  onConfirm,
  cancelText = '',
  showCancel = true,
  cancelBtnVariant = 'primary',
  confirmText = '',
  showConfirm = true,
  confirmBtnVariant = 'link',
  confirmBtnDisabled = false,
  bodyClass = '',
  scrollable = false,
  className = '',
}) => {
  const { t } = useTranslation();
  return (
    <Modal
      className={className}
      scrollable={scrollable}
      show={visible}
      onHide={onCancel}
      centered={centered}
      fullscreen="sm-down">
      <Modal.Header closeButton>
        <Modal.Title as="h5">
          {title || t('title', { keyPrefix: 'modal_confirm' })}
        </Modal.Title>
      </Modal.Header>
      <Modal.Body className={bodyClass}>{children}</Modal.Body>
      {(showCancel || showConfirm) && (
        <Modal.Footer>
          {showCancel && (
            <Button variant={cancelBtnVariant} onClick={onCancel}>
              {cancelText || t('btns.cancel')}
            </Button>
          )}
          {showConfirm && (
            <Button
              variant={confirmBtnVariant}
              onClick={(event) => {
                onConfirm?.(event);
              }}
              id="ok_button"
              disabled={confirmBtnDisabled}>
              {confirmText || t('btns.confirm')}
            </Button>
          )}
        </Modal.Footer>
      )}
    </Modal>
  );
};

export default React.memo(Index);

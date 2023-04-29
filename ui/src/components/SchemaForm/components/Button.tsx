import React, { FC, useLayoutEffect, useState } from 'react';
import { Button, ButtonProps, Spinner } from 'react-bootstrap';

import { request } from '@/utils';
import type * as Type from '@/common/interface';
import type { UIAction } from '../types';
import { useToast } from '@/hooks';

interface Props {
  fieldName: string;
  text: string;
  action: UIAction | undefined;
  formData: Type.FormDataType;
  readOnly: boolean;
  variant?: ButtonProps['variant'];
  size?: ButtonProps['size'];
}
const Index: FC<Props> = ({
  fieldName,
  action,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  formData,
  text = '',
  readOnly = false,
  variant = 'primary',
  size,
}) => {
  const Toast = useToast();
  const [isLoading, setLoading] = useState(false);
  const handleNotify = (msg, type: 'success' | 'danger' = 'success') => {
    const tm = action?.toastMessage;
    if (tm === false || !msg) {
      return;
    }
    Toast.onShow({
      msg,
      variant: type,
    });
  };
  const handleAction = () => {
    if (!action) {
      return;
    }
    setLoading(true);
    request
      .request({
        method: action.method,
        url: action.url,
        timeout: 0,
      })
      .then((resp) => {
        if ('message' in resp) {
          handleNotify(resp.message, 'success');
        }
      })
      .catch((ex) => {
        if (ex && 'msg' in ex) {
          handleNotify(ex.msg, 'danger');
        }
      })
      .finally(() => {
        setLoading(false);
      });
  };
  useLayoutEffect(() => {
    if (action?.loading?.state === 'pending') {
      setLoading(true);
    }
  }, []);
  const loadingText = action?.loading?.text || text;
  const disabled = isLoading || readOnly;

  return (
    <div className="d-flex">
      <Button
        name={fieldName}
        onClick={handleAction}
        disabled={disabled}
        size={size}
        variant={variant}>
        {isLoading ? (
          <>
            <Spinner
              className="align-middle me-2"
              animation="border"
              size="sm"
              variant={variant}
            />
            {loadingText}
          </>
        ) : (
          text
        )}
      </Button>
    </div>
  );
};

export default Index;

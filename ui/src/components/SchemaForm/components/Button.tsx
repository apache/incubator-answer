import React, { FC, useState } from 'react';
import { Button } from 'react-bootstrap';

import { request } from '@/utils';
import type * as Type from '@/common/interface';
import type { UIAction } from '../index.d';

interface Props {
  fieldName: string;
  text: string;
  action: UIAction | undefined;
  formData: Type.FormDataType;
  readOnly: boolean;
}
const Index: FC<Props> = ({
  fieldName,
  action,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  formData,
  readOnly = false,
  text = '',
}) => {
  const [isLoading, setLoading] = useState(false);
  const handleAction = async () => {
    if (!action) {
      return;
    }
    setLoading(true);
    const method = action.method || 'get';
    await request[method](action.url);
    setLoading(false);
  };
  const disabled = isLoading || readOnly;
  return (
    <div className="d-flex">
      <Button name={fieldName} onClick={handleAction} disabled={disabled}>
        {text || fieldName}
        {isLoading ? '...' : ''}
      </Button>
    </div>
  );
};

export default Index;

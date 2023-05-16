import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  type: string | undefined;
  placeholder: string | undefined;
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  formData: Type.FormDataType;
  readOnly: boolean;
}
const Index: FC<Props> = ({
  type = 'text',
  placeholder = '',
  fieldName,
  onChange,
  formData,
  readOnly = false,
}) => {
  const fieldObject = formData[fieldName];
  const handleChange = (evt: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = evt.currentTarget;
    const state = {
      ...formData,
      [name]: {
        ...formData[name],
        value,
        isInvalid: false,
      },
    };
    if (typeof onChange === 'function') {
      onChange(state);
    }
  };

  return (
    <Form.Control
      name={fieldName}
      placeholder={placeholder}
      type={type}
      value={fieldObject?.value || ''}
      onChange={handleChange}
      disabled={readOnly}
      isInvalid={fieldObject?.isInvalid}
      style={type === 'color' ? { width: '6rem' } : {}}
    />
  );
};

export default Index;

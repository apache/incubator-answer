import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  type: string | undefined;
  placeholder: string | undefined;
  fieldName: string;
  onChange: (evt: React.ChangeEvent<HTMLInputElement>, ...rest) => void;
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({
  type = 'text',
  placeholder = '',
  fieldName,
  onChange,
  formData,
}) => {
  const fieldObject = formData[fieldName];
  return (
    <Form.Control
      name={fieldName}
      placeholder={placeholder}
      type={type}
      value={fieldObject?.value || ''}
      onChange={onChange}
      style={type === 'color' ? { width: '6rem' } : {}}
      isInvalid={fieldObject?.isInvalid}
    />
  );
};

export default Index;

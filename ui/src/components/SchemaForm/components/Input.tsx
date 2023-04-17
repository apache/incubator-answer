import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  title: string;
  desc: string | undefined;
  type: string | undefined;
  placeholder: string | undefined;
  fieldName: string;
  onChange: (evt: React.ChangeEvent<HTMLInputElement>, ...rest) => void;
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({
  title,
  type = 'text',
  desc,
  placeholder = '',
  fieldName,
  onChange,
  formData,
}) => {
  const fieldObject = formData[fieldName];
  return (
    <>
      <Form.Label>{title}</Form.Label>
      <Form.Control
        name={fieldName}
        placeholder={placeholder}
        type={type}
        value={fieldObject?.value || ''}
        onChange={onChange}
        style={type === 'color' ? { width: '6rem' } : {}}
        isInvalid={fieldObject?.isInvalid}
      />
      <Form.Control.Feedback type="invalid">
        {fieldObject?.errorMsg}
      </Form.Control.Feedback>
      {desc ? <Form.Text className="text-muted">{desc}</Form.Text> : null}
    </>
  );
};

export default Index;

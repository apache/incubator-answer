import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  title: string;
  label: string | undefined;
  fieldName: string;
  onChange: (evt: React.ChangeEvent<HTMLInputElement>, ...rest) => void;
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({ title, fieldName, onChange, label, formData }) => {
  const fieldObject = formData[fieldName];
  return (
    <Form.Check
      required
      id={`switch-${title}`}
      name={fieldName}
      type="switch"
      label={label}
      checked={fieldObject?.value || ''}
      feedback={fieldObject?.errorMsg}
      feedbackType="invalid"
      isInvalid={fieldObject.isInvalid}
      onChange={onChange}
    />
  );
};

export default Index;

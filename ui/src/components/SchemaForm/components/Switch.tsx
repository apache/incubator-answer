import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  title: string;
  label: string | undefined;
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  formData: Type.FormDataType;
  readOnly?: boolean;
}
const Index: FC<Props> = ({
  title,
  fieldName,
  onChange,
  label,
  formData,
  readOnly = false,
}) => {
  const fieldObject = formData[fieldName];
  const handleChange = (evt: React.ChangeEvent<HTMLInputElement>) => {
    const { name, checked } = evt.currentTarget;
    const state = {
      ...formData,
      [name]: {
        ...formData[name],
        value: checked,
        isInvalid: false,
      },
    };
    if (typeof onChange === 'function') {
      onChange(state);
    }
  };
  return (
    <Form.Check
      id={`switch-${title}`}
      name={fieldName}
      type="switch"
      label={label}
      checked={fieldObject?.value || ''}
      feedback={fieldObject?.errorMsg}
      feedbackType="invalid"
      isInvalid={fieldObject.isInvalid}
      disabled={readOnly}
      onChange={handleChange}
    />
  );
};

export default Index;

import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import classnames from 'classnames';

import type * as Type from '@/common/interface';

interface Props {
  placeholder: string | undefined;
  rows: number | undefined;
  className: classnames.Argument;
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  formData: Type.FormDataType;
  readOnly: boolean;
}
const Index: FC<Props> = ({
  placeholder = '',
  rows = 3,
  className,
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
      as="textarea"
      name={fieldName}
      placeholder={placeholder}
      value={fieldObject?.value || ''}
      onChange={handleChange}
      isInvalid={fieldObject?.isInvalid}
      rows={rows}
      disabled={readOnly}
      className={classnames(className)}
    />
  );
};

export default Index;

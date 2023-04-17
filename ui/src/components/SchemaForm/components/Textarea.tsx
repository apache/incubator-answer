import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import classnames from 'classnames';

import type * as Type from '@/common/interface';

interface Props {
  placeholder: string | undefined;
  rows: number | undefined;
  className: classnames.Argument;
  fieldName: string;
  onChange: (evt: React.ChangeEvent<HTMLInputElement>, ...rest) => void;
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({
  placeholder = '',
  rows = 3,
  className,
  fieldName,
  onChange,
  formData,
}) => {
  const fieldObject = formData[fieldName];
  return (
    <Form.Control
      as="textarea"
      name={fieldName}
      placeholder={placeholder}
      value={fieldObject?.value || ''}
      onChange={onChange}
      isInvalid={fieldObject?.isInvalid}
      rows={rows}
      className={classnames(className)}
    />
  );
};

export default Index;

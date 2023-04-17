import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import classnames from 'classnames';

import type * as Type from '@/common/interface';

interface Props {
  title: string;
  desc: string | undefined;
  placeholder: string | undefined;
  rows: number | undefined;
  className: classnames.Argument;
  fieldName: string;
  onChange: (evt: React.ChangeEvent<HTMLInputElement>, ...rest) => void;
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({
  title,
  desc,
  placeholder = '',
  rows = 3,
  className,
  fieldName,
  onChange,
  formData,
}) => {
  const fieldObject = formData[fieldName];
  return (
    <>
      <Form.Label>{title}</Form.Label>
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
      <Form.Control.Feedback type="invalid">
        {fieldObject?.errorMsg}
      </Form.Control.Feedback>
      {desc ? <Form.Text className="text-muted">{desc}</Form.Text> : null}
    </>
  );
};

export default Index;

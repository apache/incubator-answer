import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  title: string;
  desc: string | undefined;
  fieldName: string;
  onChange: (evt: React.ChangeEvent<HTMLSelectElement>) => void;
  enumValues: (string | boolean | number)[];
  enumNames: string[];
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({
  title,
  desc,
  fieldName,
  onChange,
  enumValues,
  enumNames,
  formData,
}) => {
  const fieldObject = formData[fieldName];
  return (
    <>
      <Form.Label>{title}</Form.Label>
      <Form.Select
        aria-label={desc}
        name={fieldName}
        value={fieldObject?.value || ''}
        onChange={onChange}
        isInvalid={fieldObject?.isInvalid}>
        {enumValues?.map((item, index) => {
          return (
            <option value={String(item)} key={String(item)}>
              {enumNames?.[index]}
            </option>
          );
        })}
      </Form.Select>
      <Form.Control.Feedback type="invalid">
        {fieldObject?.errorMsg}
      </Form.Control.Feedback>
      {desc ? <Form.Text className="text-muted">{desc}</Form.Text> : null}
    </>
  );
};

export default Index;

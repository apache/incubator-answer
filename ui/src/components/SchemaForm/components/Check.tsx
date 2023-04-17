import React, { FC } from 'react';
import { Form, Stack } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  type: 'radio' | 'checkbox';
  title: string;
  desc: string | undefined;
  fieldName: string;
  onChange: (evt: React.ChangeEvent<HTMLInputElement>, ...rest) => void;
  enumValues: (string | boolean | number)[];
  enumNames: string[];
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({
  type = 'radio',
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
      <Stack direction="horizontal">
        {enumValues?.map((item, index) => {
          return (
            <Form.Check
              key={String(item)}
              inline
              required
              type={type}
              name={fieldName}
              id={`form-${String(item)}`}
              label={enumNames?.[index]}
              checked={(fieldObject?.value || '') === item}
              feedback={fieldObject?.errorMsg}
              feedbackType="invalid"
              isInvalid={fieldObject?.isInvalid}
              onChange={(evt) => onChange(evt, index)}
            />
          );
        })}
      </Stack>
      <Form.Control.Feedback type="invalid">
        {fieldObject?.errorMsg}
      </Form.Control.Feedback>
      {desc ? <Form.Text className="text-muted">{desc}</Form.Text> : null}
    </>
  );
};

export default Index;

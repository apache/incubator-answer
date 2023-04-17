import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';
import TimeZonePicker from '@/components/TimeZonePicker';

interface Props {
  title: string;
  desc: string | undefined;
  fieldName: string;
  onChange: (evt: React.ChangeEvent<HTMLSelectElement>, ...rest) => void;
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({ title, desc, fieldName, onChange, formData }) => {
  const fieldObject = formData[fieldName];
  return (
    <>
      <Form.Label>{title}</Form.Label>
      <TimeZonePicker
        value={fieldObject?.value || ''}
        isInvalid={fieldObject?.isInvalid}
        name={fieldName}
        onChange={onChange}
      />
      <Form.Control.Feedback type="invalid">
        {fieldObject?.errorMsg}
      </Form.Control.Feedback>
      {desc ? <Form.Text className="text-muted">{desc}</Form.Text> : null}
    </>
  );
};

export default Index;

import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';
import BrandUpload from '@/components/BrandUpload';

interface Props {
  title: string;
  type: Type.UploadType | undefined;
  acceptType: string | undefined;
  desc: string | undefined;
  fieldName: string;
  onChange: (key, val) => void;
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({
  title,
  type = 'avatar',
  acceptType = '',
  desc,
  fieldName,
  onChange,
  formData,
}) => {
  const fieldObject = formData[fieldName];
  return (
    <>
      <Form.Label>{title}</Form.Label>
      <BrandUpload
        type={type}
        acceptType={acceptType}
        value={fieldObject?.value}
        onChange={(value) => onChange(fieldName, value)}
      />
      <Form.Control
        name={fieldName}
        className="d-none"
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

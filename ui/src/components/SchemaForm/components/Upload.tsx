import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';
import BrandUpload from '@/components/BrandUpload';

interface Props {
  type: Type.UploadType | undefined;
  acceptType: string | undefined;
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({
  type = 'avatar',
  acceptType = '',
  fieldName,
  onChange,
  formData,
}) => {
  const fieldObject = formData[fieldName];
  const handleChange = (name: string, value: string) => {
    console.log('upload: ', name, value);
    const state = {
      ...formData,
      [name]: {
        ...formData[name],
        value,
      },
    };
    if (typeof onChange === 'function') {
      onChange(state);
    }
  };
  return (
    <>
      <BrandUpload
        type={type}
        acceptType={acceptType}
        value={fieldObject?.value}
        onChange={(value) => handleChange(fieldName, value)}
      />
      <Form.Control
        name={fieldName}
        className="d-none"
        isInvalid={fieldObject?.isInvalid}
      />
    </>
  );
};

export default Index;

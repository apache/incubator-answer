import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import classNames from 'classnames';

import type * as Type from '@/common/interface';
import BrandUpload from '@/components/BrandUpload';

interface Props {
  type: Type.UploadType | undefined;
  acceptType: string | undefined;
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  formData: Type.FormDataType;
  readOnly?: boolean;
  imgClassNames?: classNames.Argument;
}
const Index: FC<Props> = ({
  type = 'avatar',
  acceptType = '',
  fieldName,
  onChange,
  formData,
  readOnly = false,
  imgClassNames = '',
}) => {
  const fieldObject = formData[fieldName];
  const handleChange = (name: string, value: string) => {
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
        readOnly={readOnly}
        onChange={(value) => handleChange(fieldName, value)}
        imgClassNames={imgClassNames}
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

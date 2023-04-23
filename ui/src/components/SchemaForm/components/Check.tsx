import React, { FC } from 'react';
import { Form, Stack } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  type: 'radio' | 'checkbox';
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  enumValues: (string | boolean | number)[];
  enumNames: string[];
  formData: Type.FormDataType;
  readOnly?: boolean;
}
const Index: FC<Props> = ({
  type = 'radio',
  fieldName,
  onChange,
  enumValues,
  enumNames,
  formData,
  readOnly = false,
}) => {
  const fieldObject = formData[fieldName];
  const handleCheck = (
    evt: React.ChangeEvent<HTMLInputElement>,
    index: number,
  ) => {
    const { name, checked } = evt.currentTarget;
    const freshVal = checked ? enumValues?.[index] : '';
    const state = {
      ...formData,
      [name]: {
        ...formData[name],
        value: freshVal,
        isInvalid: false,
      },
    };
    if (typeof onChange === 'function') {
      onChange(state);
    }
  };
  return (
    <Stack direction="horizontal">
      {enumValues?.map((item, index) => {
        return (
          <Form.Check
            key={String(item)}
            inline
            type={type}
            name={fieldName}
            id={`form-${String(item)}`}
            label={enumNames?.[index]}
            checked={(fieldObject?.value || '') === item}
            feedback={fieldObject?.errorMsg}
            feedbackType="invalid"
            isInvalid={fieldObject?.isInvalid}
            disabled={readOnly}
            onChange={(evt) => handleCheck(evt, index)}
          />
        );
      })}
    </Stack>
  );
};

export default Index;

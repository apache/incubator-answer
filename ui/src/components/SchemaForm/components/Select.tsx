import React, { FC } from 'react';
import { Form } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  desc: string | undefined;
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  enumValues: (string | boolean | number)[];
  enumNames: string[];
  formData: Type.FormDataType;
  readOnly: boolean;
}
const Index: FC<Props> = ({
  desc,
  fieldName,
  onChange,
  enumValues,
  enumNames,
  formData,
  readOnly = false,
}) => {
  const fieldObject = formData[fieldName];
  const handleChange = (evt: React.ChangeEvent<HTMLSelectElement>) => {
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
    <Form.Select
      aria-label={desc}
      name={fieldName}
      value={fieldObject?.value || ''}
      onChange={handleChange}
      disabled={readOnly}
      isInvalid={fieldObject?.isInvalid}>
      {enumValues?.map((item, index) => {
        return (
          <option value={String(item)} key={String(item)}>
            {enumNames?.[index]}
          </option>
        );
      })}
    </Form.Select>
  );
};

export default Index;

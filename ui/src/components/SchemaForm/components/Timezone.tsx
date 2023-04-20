import React, { FC } from 'react';

import type * as Type from '@/common/interface';
import TimeZonePicker from '@/components/TimeZonePicker';

interface Props {
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  formData: Type.FormDataType;
  readOnly?: boolean;
}
const Index: FC<Props> = ({
  fieldName,
  onChange,
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
    <TimeZonePicker
      value={fieldObject?.value || ''}
      isInvalid={fieldObject?.isInvalid}
      name={fieldName}
      disabled={readOnly}
      onChange={handleChange}
    />
  );
};

export default Index;

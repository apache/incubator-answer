import React, { FC } from 'react';

import type * as Type from '@/common/interface';
import TimeZonePicker from '@/components/TimeZonePicker';

interface Props {
  fieldName: string;
  onChange: (evt: React.ChangeEvent<HTMLSelectElement>, ...rest) => void;
  formData: Type.FormDataType;
}
const Index: FC<Props> = ({ fieldName, onChange, formData }) => {
  const fieldObject = formData[fieldName];
  return (
    <TimeZonePicker
      value={fieldObject?.value || ''}
      isInvalid={fieldObject?.isInvalid}
      name={fieldName}
      onChange={onChange}
    />
  );
};

export default Index;

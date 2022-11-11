import { FC, memo, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { JSONSchema, SchemaForm, UISchema } from '@/components';
import { FormDataType } from '@/common/interface';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.branding',
  });

  const [formData, setFormData] = useState<FormDataType>({
    logo: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    mobile_logo: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    square_icon: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    favicon: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });

  // const onChange = (fieldName, fieldValue) => {
  //   if (!formData[fieldName]) {
  //     return;
  //   }
  //   const fieldData: FormDataType = {
  //     [fieldName]: {
  //       value: fieldValue,
  //       isInvalid: false,
  //       errorMsg: '',
  //     },
  //   };
  //   setFormData({ ...formData, ...fieldData });
  // };

  // const [img, setImg] = useState(
  //   'https://image-static.segmentfault.com/405/057/4050570037-636c7b0609a49',
  // );

  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      logo: {
        type: 'string',
        title: t('logo.label'),
        description: t('logo.text'),
      },
      mobile_logo: {
        type: 'string',
        title: t('mobile_logo.label'),
        description: t('mobile_logo.text'),
      },
      square_icon: {
        type: 'string',
        title: t('square_icon.label'),
        description: t('square_icon.text'),
      },
      favicon: {
        type: 'string',
        title: t('favicon.label'),
        description: t('favicon.text'),
      },
    },
  };

  const uiSchema: UISchema = {
    logo: {
      'ui:widget': 'upload',
      'ui:options': {
        imageType: 'logo',
      },
    },
    mobile_logo: {
      'ui:widget': 'upload',
      'ui:options': {
        imageType: 'mobile_logo',
      },
    },
    square_icon: {
      'ui:widget': 'upload',
      'ui:options': {
        imageType: 'square_icon',
      },
    },
    favicon: {
      'ui:widget': 'upload',
      'ui:options': {
        imageType: 'favicon',
      },
    },
  };

  const handleOnChange = (data) => {
    setFormData(data);
  };

  const onSubmit = () => {
    // undo
  };

  return (
    <div>
      <h3 className="mb-4">{t('page_title')}</h3>
      <SchemaForm
        schema={schema}
        uiSchema={uiSchema}
        formData={formData}
        onSubmit={onSubmit}
        onChange={handleOnChange}
      />
    </div>
  );
};

export default memo(Index);

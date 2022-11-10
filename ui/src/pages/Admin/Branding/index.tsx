import { FC, memo, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { BrandUpload } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.branding',
  });

  const [img, setImg] = useState(
    'https://image-static.segmentfault.com/405/057/4050570037-636c7b0609a49',
  );

  return (
    <div>
      <h3 className="mb-4">{t('page_title')}</h3>
      <BrandUpload type="logo" value={img} onChange={setImg} />
    </div>
  );
};

export default memo(Index);

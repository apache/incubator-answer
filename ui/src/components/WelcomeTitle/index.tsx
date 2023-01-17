import React, { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import { siteInfoStore } from '@/stores';

interface Props {
  as?: React.ElementType;
  className?: string;
}

const Index: FC<Props> = ({ as: Component = 'h3', className = 'mb-5' }) => {
  const { t } = useTranslation();
  const { name: siteName } = siteInfoStore((_) => _.siteInfo);
  return (
    <Component className={classnames('text-center', className)}>
      {t('website_welcome', { site_name: siteName })}
    </Component>
  );
};

export default memo(Index);

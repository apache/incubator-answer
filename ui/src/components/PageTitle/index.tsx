import { FC } from 'react';

import { siteInfoStore } from '@/stores';

interface IProp {
  title?: string;
  suffix?: string;
}
const setPageTitle = (title) => {
  if (document) {
    document.title = title;
  }
  return null;
};
// TODO: use Helmet for static response
const PageTitle: FC<IProp> = ({ title = '', suffix = '' }) => {
  const { siteInfo } = siteInfoStore();
  if (!suffix) {
    suffix = `${siteInfo.name}`;
  }
  title = title ? `${title} - ${suffix}` : suffix;
  return <>{setPageTitle(title)}</>;
};

export default PageTitle;

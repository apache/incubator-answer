import { FC } from 'react';
import { useTranslation } from 'react-i18next';

const Dashboard: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.dashboard' });
  return (
    <>
      <h3 className="text-capitalize">{t('title')}</h3>
      <p className="mt-4">{t('welcome')}</p>
      {process.env.REACT_APP_VERSION && (
        <p className="mt-4">
          {`${t('version')} `}
          {process.env.REACT_APP_VERSION}
        </p>
      )}
    </>
  );
};
export default Dashboard;

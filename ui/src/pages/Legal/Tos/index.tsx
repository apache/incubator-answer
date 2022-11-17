import { FC } from 'react';
import { useTranslation } from 'react-i18next';

import { useLegalTos } from '@/services';
import { PageTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'nav_menus' });
  const { data: tos } = useLegalTos();
  const contentText = tos?.terms_of_service_original_text;
  let matchUrl: URL | undefined;
  try {
    if (contentText) {
      matchUrl = new URL(contentText);
    }
    // eslint-disable-next-line no-empty
  } catch (ex) {}
  if (matchUrl) {
    window.location.replace(matchUrl.toString());
    return null;
  }
  return (
    <>
      <PageTitle title={t('tos')} />
      <h3>{t('tos')}</h3>
      <div
        className="fmt"
        dangerouslySetInnerHTML={{
          __html: tos?.terms_of_service_parsed_text || '',
        }}
      />
    </>
  );
};

export default Index;

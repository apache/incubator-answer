import { FC } from 'react';
import { useTranslation } from 'react-i18next';

import { useLegalPrivacy } from '@/services';
import { PageTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'nav_menus' });
  const { data: privacy } = useLegalPrivacy();
  const contentText = privacy?.privacy_policy_original_text;
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
      <PageTitle title={t('privacy')} />
      <h3>{t('privacy')}</h3>
      <div
        className="fmt"
        dangerouslySetInnerHTML={{
          __html: privacy?.privacy_policy_parsed_text || '',
        }}
      />
    </>
  );
};

export default Index;

import { FC, useEffect } from 'react';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { useLegalPrivacy } from '@/services';
import { htmlRender } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'nav_menus' });
  usePageTags({
    title: t('privacy'),
  });
  const { data: privacy } = useLegalPrivacy();
  const contentText = privacy?.privacy_policy_original_text;
  let matchUrl: URL | undefined;

  useEffect(() => {
    const fmt = document.querySelector('.fmt') as HTMLElement;
    if (!fmt) {
      return;
    }
    htmlRender(fmt);
  }, [privacy?.privacy_policy_parsed_text]);

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
      <h3 className="mb-4">{t('privacy')}</h3>
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

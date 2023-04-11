import { FC, useEffect } from 'react';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { useLegalTos } from '@/services';
import { htmlRender } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'nav_menus' });
  usePageTags({
    title: t('tos'),
  });
  const { data: tos } = useLegalTos();
  const contentText = tos?.terms_of_service_original_text;
  let matchUrl: URL | undefined;

  useEffect(() => {
    const fmt = document.querySelector('.fmt') as HTMLElement;
    if (!fmt) {
      return;
    }
    htmlRender(fmt);
  }, [tos?.terms_of_service_parsed_text]);

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
    <div>
      <h3 className="mb-4">{t('tos')}</h3>
      <div
        className="fmt"
        dangerouslySetInnerHTML={{
          __html: tos?.terms_of_service_parsed_text || '',
        }}
      />
    </div>
  );
};

export default Index;

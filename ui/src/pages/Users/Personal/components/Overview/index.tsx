import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import TopList from '../TopList';

interface Props {
  visible: boolean;
  introduction: string;
  data;
}
const Index: FC<Props> = ({ visible, introduction, data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });
  if (!visible) {
    return null;
  }
  return (
    <div>
      <h5 className="mb-3">{t('about_me')}</h5>
      {introduction ? (
        <div
          className="mb-4 text-break"
          dangerouslySetInnerHTML={{ __html: introduction }}
        />
      ) : (
        <div className="text-center py-5 mb-4">{t('about_me_empty')}</div>
      )}

      {data?.answer?.length > 0 && (
        <>
          <h5 className="mb-3">{t('top_answers')}</h5>
          <TopList data={data?.answer} type="answer" />
        </>
      )}

      {data?.question?.length > 0 && (
        <>
          <h5 className="mb-3">{t('top_questions')}</h5>
          <TopList data={data?.question} type="question" />
        </>
      )}
    </div>
  );
};

export default memo(Index);

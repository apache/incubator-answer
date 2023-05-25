import { memo, FC } from 'react';
import { Card } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';

const Index: FC = () => {
  const { t } = useTranslation();
  return (
    <Card>
      <Card.Header>{t('search.tips.title')}</Card.Header>
      <Card.Body className="small ext-secondary">
        <div className="mb-1">
          <Trans i18nKey="search.tips.tag" components={{ 1: <code /> }} />
        </div>
        <div className="mb-1">
          <Trans i18nKey="search.tips.user" components={{ 1: <code /> }} />
        </div>
        <div className="mb-1">
          <Trans i18nKey="search.tips.answer" components={{ 1: <code /> }} />
        </div>
        <div className="mb-1">
          <Trans i18nKey="search.tips.score" components={{ 1: <code /> }} />
        </div>
        <div className="mb-1">
          <Trans i18nKey="search.tips.question" components={{ 1: <code /> }} />
        </div>
        <div>
          <Trans i18nKey="search.tips.is_answer" components={{ 1: <code /> }} />
        </div>
      </Card.Body>
    </Card>
  );
};

export default memo(Index);

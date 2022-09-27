import { FC, memo } from 'react';
import { Container } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useMatch } from 'react-router-dom';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.admin_header',
  });
  const adminPathMatch = useMatch('/admin/*');
  if (!adminPathMatch) {
    return null;
  }
  return (
    <div className="bg-light py-2">
      <Container className="py-1">
        <h6 className="mb-0 fw-bold lh-base">{t('title')}</h6>
      </Container>
    </div>
  );
};

export default memo(Index);

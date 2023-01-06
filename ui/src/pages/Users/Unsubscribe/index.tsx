import { FC, memo, useEffect } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Link, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { unsubscribe } from '@/services';
import { usePageTags } from '@/hooks';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'unsubscribe' });
  usePageTags({
    title: t('page_title'),
  });
  const [searchParams] = useSearchParams();
  const code = searchParams.get('code');
  useEffect(() => {
    if (code) {
      unsubscribe(code);
    }
  }, [code]);
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col lg={6}>
          <h3 className="text-center mt-3 mb-5">{t('success_title')}</h3>
          <p className="text-center">{t('success_desc')}</p>
          <div className="text-center">
            <Link to="/users/settings/notify">{t('link')}</Link>
          </div>
        </Col>
      </Row>
    </Container>
  );
};

export default memo(Index);

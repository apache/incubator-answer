import { FC, memo } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Link, useLocation } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { WelcomeTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'account_result' });
  const location = useLocation();
  usePageTags({
    title: t('account_activation', { keyPrefix: 'page_title' }),
  });
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col lg={6}>
          <WelcomeTitle className="mt-3 mb-5" />
          {location.pathname?.includes('success') && (
            <>
              <p className="text-center">{t('success')}</p>
              <div className="text-center">
                <Link to="/">{t('link')}</Link>
              </div>
            </>
          )}

          {location.pathname?.includes('failed') && (
            <p className="text-center">{t('invalid')}</p>
          )}
        </Col>
      </Row>
    </Container>
  );
};

export default memo(Index);

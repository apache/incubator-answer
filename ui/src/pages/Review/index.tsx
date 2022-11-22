import { FC } from 'react';
import {
  Container,
  Row,
  Col,
  Alert,
  Badge,
  Stack,
  Button,
} from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { BaseUserCard, FormatTime, Empty } from '@/components';
import { loggedUserInfoStore } from '@/stores';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_review' });

  const { user } = loggedUserInfoStore.getState();
  return (
    <Container className="pt-2 mt-4 mb-5">
      <Row>
        <Col lg={{ span: 7, offset: 1 }}>
          <h3 className="mb-4">{t('review')}</h3>
          <Alert variant="secondary">
            <Stack className="align-items-start">
              <Badge bg="secondary" className="mb-2">
                {t('question_edit')}
              </Badge>
              <Link to="/review">
                How do I test whether variable against multiple
              </Link>
              <p className="mb-0">
                {t('edit_summary')}: Editing part of the code and correcting the
                grammar.
              </p>
            </Stack>
            <Stack
              direction="horizontal"
              gap={1}
              className="align-items-baseline mt-2">
              <BaseUserCard data={user} avatarSize="24" />
              <FormatTime
                time={Date.now()}
                className="small text-secondary"
                preFix={t('proposed')}
              />
            </Stack>
          </Alert>
        </Col>
        <Col lg={{ span: 7, offset: 1 }}>Content</Col>
        <Col lg={{ span: 7, offset: 1 }}>
          <Stack direction="horizontal" gap={2}>
            <Button variant="outline-primary">
              {t('approve', { keyPrefix: 'btns' })}
            </Button>
            <Button variant="outline-primary">
              {t('reject', { keyPrefix: 'btns' })}
            </Button>
            <Button variant="outline-primary">
              {t('skip', { keyPrefix: 'btns' })}
            </Button>
          </Stack>
        </Col>
        <Col lg={{ span: 7, offset: 1 }}>
          <Empty>{t('empty')}</Empty>
        </Col>
      </Row>
    </Container>
  );
};

export default Index;

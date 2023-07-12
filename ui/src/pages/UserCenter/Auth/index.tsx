import { memo } from 'react';
import { Container, Col } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { userCenterStore } from '@/stores';
import { USER_AGENT_NAMES } from '@/common/constants';
import { usePageTags } from '@/hooks';

import WeComAuth from './components/WeCom';

const Index = () => {
  const { t } = useTranslation('translation');
  const [searchParam] = useSearchParams();
  const { agent: ucAgent } = userCenterStore();
  let agentName = ucAgent?.agent_info?.name || '';
  if (searchParam.get('agent_name')) {
    agentName = searchParam.get('agent_name') || '';
  }
  usePageTags({
    title: t('login', { keyPrefix: 'page_title' }),
  });
  return (
    <Container style={{ paddingTop: '5rem', paddingBottom: '5rem' }}>
      <Col md={4} className="mx-auto">
        {USER_AGENT_NAMES.WeCom.toLowerCase() === agentName.toLowerCase() ? (
          <WeComAuth />
        ) : null}
      </Col>
    </Container>
  );
};

export default memo(Index);

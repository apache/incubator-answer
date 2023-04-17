import { memo } from 'react';
import { Container } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';

import { userCenterStore } from '@/stores';

import WeCom from './components/WeCom';

const Index = () => {
  const [searchParam] = useSearchParams();
  const { agent: ucAgent } = userCenterStore();
  let agentName = ucAgent?.agent_info.name || '';
  if (searchParam.get('agent_name')) {
    agentName = searchParam.get('agent_name') || '';
  }
  return <Container>{/^WeCom/i.test(agentName) ? <WeCom /> : null}</Container>;
};

export default memo(Index);

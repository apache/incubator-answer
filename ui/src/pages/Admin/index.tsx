import { FC, useEffect } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Outlet, useMatch } from 'react-router-dom';

import { cloneDeep } from 'lodash';

import { usePageTags } from '@/hooks';
import { AccordionNav } from '@/components';
import { ADMIN_NAV_MENUS } from '@/common/constants';
import { useQueryPlugins } from '@/services';
import { interfaceStore } from '@/stores';

import './index.scss';

const g10Paths = [
  'dashboard',
  'questions',
  'answers',
  'users',
  'flags',
  'installed-plugins',
];
const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const pathMatch = useMatch('/admin/:path');
  const curPath = pathMatch?.params.path || 'dashboard';

  const interfaceLang = interfaceStore((_) => _.interface.language);
  const { data: configurablePlugins, mutate: updateConfigurablePlugins } =
    useQueryPlugins({
      status: 'active',
      have_config: true,
    });

  const menus = cloneDeep(ADMIN_NAV_MENUS);
  if (configurablePlugins && configurablePlugins.length > 0) {
    menus.forEach((item) => {
      if (item.name === 'plugins' && item.children) {
        item.children = [
          ...item.children,
          ...configurablePlugins.map((plugin) => ({
            name: plugin.slug_name,
            displayName: plugin.name,
          })),
        ];
      }
    });
  }

  const observePlugins = (evt) => {
    if (evt.data.msgType === 'refreshConfigurablePlugins') {
      updateConfigurablePlugins();
    }
  };
  useEffect(() => {
    window.addEventListener('message', observePlugins);
    return () => {
      window.removeEventListener('message', observePlugins);
    };
  }, []);
  useEffect(() => {
    updateConfigurablePlugins();
  }, [interfaceLang]);

  usePageTags({
    title: t('admin'),
  });
  return (
    <>
      <div className="bg-light py-2">
        <Container className="py-1">
          <h6 className="mb-0 fw-bold lh-base">
            {t('title', { keyPrefix: 'admin.admin_header' })}
          </h6>
        </Container>
      </div>
      <Container className="admin-container">
        <Row>
          <Col lg={2}>
            <AccordionNav menus={menus} path="/admin/" />
          </Col>
          <Col lg={g10Paths.find((v) => curPath === v) ? 10 : 6}>
            <Outlet />
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default Index;

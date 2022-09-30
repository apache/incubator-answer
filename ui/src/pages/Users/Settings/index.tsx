import React, { useState, useEffect } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Outlet } from 'react-router-dom';

import { getUserInfo } from '@answer/api';
import type { FormDataType } from '@answer/common/interface';

import Nav from './components/Nav';

import { PageTitle } from '@/components';

const Index: React.FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.profile',
  });

  const [formData, setFormData] = useState<FormDataType>({
    display_name: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    avatar: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    bio: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    website: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    location: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });
  const getProfile = () => {
    getUserInfo().then((res) => {
      formData.display_name.value = res.display_name;
      formData.bio.value = res.bio;
      formData.avatar.value = res.avatar;
      formData.location.value = res.location;
      formData.website.value = res.website;
      setFormData({ ...formData });
    });
  };

  useEffect(() => {
    getProfile();
  }, []);
  return (
    <>
      <PageTitle title={t('settings', { keyPrefix: 'page_title' })} />
      <Container className="mt-4 mb-5 pb-5">
        <Row className="justify-content-center">
          <Col xxl={10} md={12}>
            <h3 className="mb-4">
              {t('page_title', { keyPrefix: 'settings' })}
            </h3>
          </Col>
        </Row>

        <Row>
          <Col xxl={1} />
          <Col md={3} lg={2} className="mb-3">
            <Nav />
          </Col>
          <Col md={9} lg={6}>
            <Outlet />
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default React.memo(Index);

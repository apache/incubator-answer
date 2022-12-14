import { Container, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import { usePageTags } from '@/hooks';
import { useQueryContributeUsers } from '@/services';
import { Avatar } from '@/components';

const Users = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'users' });

  const { data: users } = useQueryContributeUsers();

  usePageTags({
    title: t('users', { keyPrefix: 'page_title' }),
  });

  if (!users) {
    return null;
  }

  const keys = Object.keys(users);
  return (
    <Container className="py-3 my-3">
      <Row className="mb-4 d-flex justify-content-center">
        <Col xxl={10} sm={12}>
          <h3 className="mb-4">{t('title')}</h3>
        </Col>

        <Col xxl={10} sm={12}>
          {keys.map((key, index) => {
            if (users[key]?.length === 0) {
              return null;
            }
            return (
              <>
                <Row className="mb-4">
                  <Col>
                    <h6 className="mb-0">{t(key)}</h6>
                  </Col>
                </Row>
                <Row className={index === keys.length - 1 ? '' : 'mb-4'}>
                  {users[key]?.map((user) => (
                    <Col
                      key={user.username}
                      xs={12}
                      lg={3}
                      md={4}
                      sm={6}
                      className="mb-4">
                      <div className="d-flex">
                        <Avatar size="48px" avatar={user?.avatar} />

                        <div className="ms-2">
                          <Link to={`/users/${user.username}`}>
                            {user.display_name}
                          </Link>
                          <div className="text-secondary fs-14">
                            {key === 'users_with_the_most_vote'
                              ? `${user.vote_count} ${t('votes')}`
                              : `${user.rank} ${t('reputation')}`}
                          </div>
                        </div>
                      </div>
                    </Col>
                  ))}
                </Row>
              </>
            );
          })}
        </Col>
      </Row>
    </Container>
  );
};

export default Users;

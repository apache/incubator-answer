import { FC } from 'react';
import { Card, Badge } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import { BaseUserCard, Tag, FormatTime, Avatar } from '@/components';

const tag = [
  {
    display_name: 'bug',
    slug_name: 'bug',
    original_text: '111',
    recommend: true,
  },
  {
    display_name: 'react',
    slug_name: 'react',
    original_text: '222',
    reserved: true,
  },
  {
    display_name: 'test',
    slug_name: 'test',
    original_text: '111',
    recommend: false,
    reserved: false,
  },
];

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_review' });
  const objectType = 'question';
  return (
    <Card>
      <Card.Header>{t('flag_type', { type: 'post' })}</Card.Header>
      <Card.Body className="p-0">
        <div className="p-3">
          <h5 className="mb-3">
            How do I test weather variable against multiple
          </h5>
          {objectType === 'question' && (
            <div className="mb-4">
              {tag?.map((item) => {
                return (
                  <Tag key={item.slug_name} className="me-1" data={item} />
                );
              })}
            </div>
          )}
          <div className="small font-monospace">
            Python is a multi-paradigm, dynamically typed, multi-purpose
            programming language. It is designed to be quick to learn,
            understand, and use, and enforces a clean and uniform syntax. Please
            note that Python 2 is officially out of support as of 2020-01-01.
            For version-specific Python questions, add the [python-2.7] or
            [python-3.x] tag. When using a Python variant library (e.g. Pandas,
            NumPy), please include it in the tags.
          </div>
          <div className="d-flex align-items-center justify-content-between mt-4">
            <Badge bg="success">normal</Badge>
            <div className="d-flex align-items-center small">
              <BaseUserCard
                data={{
                  username: 'username',
                  display_name: 'username',
                  avatar: '',
                  reputation: 100,
                }}
                avatarSize="24"
              />
              <FormatTime
                time={1688107033}
                className="text-secondary ms-1 flex-shrink-0"
                preFix="answered"
              />
            </div>
          </div>
        </div>

        <div className="p-3 d-flex">
          <Avatar
            avatar=""
            size="40"
            searchStr="s=48"
            alt=""
            className="me-2"
          />
          <div className="small">
            <Link to="/test">
              111 <span className="text-secondary">@111</span>
            </Link>
            <div className="mt-1">
              I'm a web developer with in-depth experience in UI/UX design.
            </div>
            <div className="text-secondary mt-1">280 {t('reputation')}</div>
          </div>
        </div>
      </Card.Body>
    </Card>
  );
};

export default Index;

import { FC } from 'react';
import { Container, Row, Col, Form, Table } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { loggedUserInfoStore } from '@/stores';
import { useTimelineData } from '@/services';

import HistoryItem from './components/Item';

// const list = [
//   {
//     activity_id: 1,
//     revision_id: 1,
//     created_at: 1669084579,
//     activity_type: 'deleted',
//     username: 'John Doe',
//     user_display_name: 'John Doe',
//     comment: '啊撒旦法师打发房管局挥洒过短发合计干哈就撒刚发几哈',
//     object_id: '1',
//     object_type: 'question',
//     cancelled: false,
//     cancelled_at: null,
//   },
//   {
//     activity_id: 2,
//     revision_id: 2,
//     created_at: 1669084579,
//     activity_type: 'undeleted',
//     username: 'John Doe2',
//     user_display_name: 'John Doe2',
//     comment: '啊撒旦法师打发房管局挥洒过短发合计干哈就撒刚发几哈',
//     object_id: '2',
//     object_type: 'question',
//     cancelled: false,
//     cancelled_at: null,
//   },
//   {
//     activity_id: 3,
//     revision_id: 3,
//     created_at: 1669084579,
//     activity_type: 'downvote',
//     username: 'johndoe3',
//     user_display_name: 'John Doe3',
//     comment: '啊撒旦法师打发房管局挥洒过短发合计干哈就撒刚发几哈',
//     object_id: '3',
//     object_type: 'question',
//     cancelled: true,
//     cancelled_at: 1637021579,
//   },
//   {
//     activity_id: 4,
//     revision_id: 4,
//     created_at: 1669084579,
//     activity_type: 'rollback',
//     username: 'johndoe4',
//     user_display_name: 'John Doe4',
//     comment: '啊撒旦法师打发房管局挥洒过短发合计干哈就撒刚发几哈',
//     object_id: '4',
//     object_type: 'question',
//     cancelled: false,
//     cancelled_at: null,
//   },
//   {
//     activity_id: 5,
//     revision_id: 5,
//     created_at: 1669084579,
//     activity_type: 'edited',
//     username: 'johndoe4',
//     user_display_name: 'John Doe4',
//     object_id: '5',
//     object_type: 'question',
//     comment: '啊撒旦法师打发房管局挥洒过短发合计干哈就撒刚发几哈',
//     cancelled: false,
//     cancelled_at: null,
//   },
// ];

// const object_info = {
//   title: '问题标题，当回答时也是问题标题，当为 tag 时是 slug_name',
//   object_type: 'question', // question/answer/tag
//   question_id: 'xxxxxxxxxxxxxxxxxxx',
//   answer_id: 'xxxxxxxxxxxxxxxx',
// };

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'timeline' });
  const { is_admin } = loggedUserInfoStore((state) => state.user);

  const { data: timelineData } = useTimelineData({
    object_id: '10010000000000001',
    object_type: 'question',
    show_vote: false,
  });

  console.log('timelineData=', timelineData);

  return (
    <Container className="py-3">
      <Row className="py-3 justify-content-center">
        <Col xxl={10}>
          <h5 className="mb-4">
            {t('title')} <Link to="/">{timelineData?.object_info?.title}</Link>
          </h5>
          <Form.Check
            className="mb-4"
            type="switch"
            id="custom-switch"
            label={t('show_votes')}
          />
          <Table hover>
            <thead>
              <tr>
                <th style={{ width: '20%' }}>Datetime</th>
                <th style={{ width: '15%' }}>Type</th>
                <th style={{ width: '19%' }}>By</th>
                <th>Comment</th>
              </tr>
            </thead>
            <tbody>
              {timelineData?.timeline?.map((item) => {
                return (
                  <HistoryItem
                    data={item}
                    objectInfo={timelineData?.object_info}
                    key={item.revision_id}
                    isAdmin={is_admin}
                    source="question"
                  />
                );
              })}
            </tbody>
          </Table>
        </Col>
      </Row>
    </Container>
  );
};

export default Index;

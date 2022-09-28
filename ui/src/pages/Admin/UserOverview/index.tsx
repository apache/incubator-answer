import { FC } from 'react';
import {
  Container,
  Row,
  Col,
  Button,
  Table,
  Figure,
  Stack,
} from 'react-bootstrap';

import { AccordionNav } from '@answer/components';

import { ADMIN_NAV_MENUS } from '@answer/common/constants';
import '../index.scss';

const UserOverview: FC = () => {
  return (
    <Container className="admin-container">
      <Row>
        <Col lg={2}>
          <AccordionNav menus={ADMIN_NAV_MENUS} />
        </Col>
        <Col lg={10}>
          <Button variant="outline-secondary" size="sm">
            ← Back
          </Button>
          <h5 className="mb-3 mt-4">Profile</h5>
          <Table className="mb-5">
            <tbody className="align-middle">
              <tr>
                <td>ID</td>
                <td>1030000000091295</td>
                <td />
              </tr>
              <tr>
                <td>Display name</td>
                <td>Jim Green</td>
                <td />
              </tr>
              <tr>
                <td>username</td>
                <td>jimgreen</td>
                <td>
                  <Button variant="link" size="sm">
                    Edit
                  </Button>
                </td>
              </tr>
              <tr>
                <td>Profile image</td>
                <td>
                  <Figure.Image
                    width={48}
                    height={48}
                    className="rounded-1 m-0"
                    src="https://gw.alicdn.com/bao/uploaded/i4/1607723262/O1CN01JJCGVD1Zy2jryOhDc_!!1607723262.jpg"
                  />
                </td>
                <td />
              </tr>
            </tbody>
          </Table>
          <h5 className="mb-3 mt-4">Permissions</h5>
          <Table className="mb-5">
            <tbody className="align-middle">
              <tr>
                <td>Activated</td>
                <td>No</td>
                <td>
                  <Button size="sm" variant="link">
                    Activate
                  </Button>
                </td>
              </tr>
              <tr>
                <td>Admin?</td>
                <td>No</td>
                <td>
                  <Button size="sm" variant="link">
                    Grant
                  </Button>
                </td>
              </tr>
              <tr>
                <td>Suspended?</td>
                <td>No</td>
                <td>
                  <Stack direction="horizontal" gap={1}>
                    <Button size="sm" variant="link" className="text-danger">
                      Suspend
                    </Button>
                    <div className="text-secondary text-nowrap">
                      A suspended user can't log in
                    </div>
                  </Stack>
                </td>
              </tr>
            </tbody>
          </Table>
          <h5 className="mb-3 mt-4">Activity</h5>
          <Table className="mb-5">
            <tbody className="align-middle">
              <tr>
                <td>Reputation</td>
                <td>1805</td>
              </tr>
              <tr>
                <td>Answers</td>
                <td>30</td>
              </tr>
              <tr>
                <td>Questions</td>
                <td>10</td>
              </tr>
              <tr>
                <td>Created</td>
                <td>Sep 1, 2022 at 16:00</td>
              </tr>
              <tr>
                <td>Registration IP address</td>
                <td>11.22.33.44</td>
              </tr>
              <tr>
                <td>Seen</td>
                <td>Sep 6, 2022 at 09:35</td>
              </tr>
              <tr>
                <td>Last IP address</td>
                <td>11.22.33.44</td>
              </tr>
            </tbody>
          </Table>
        </Col>
      </Row>
    </Container>
  );
};

export default UserOverview;

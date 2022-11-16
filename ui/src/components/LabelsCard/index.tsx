import { Card, Dropdown, Form } from 'react-bootstrap';

import Label from '../Label';

const Labels = ({ className }) => {
  return (
    <Card className={className}>
      <Card.Header className="d-flex justify-content-between align-items-center">
        <Card.Title className="mb-0">Labels</Card.Title>

        <Dropdown align="end">
          <Dropdown.Toggle variant="link" className="no-toggle">
            Edit
          </Dropdown.Toggle>

          <Dropdown.Menu className="p-3">
            <Form.Check className="mb-2" type="checkbox" label="featured" />
            <Form.Check className="mb-2" type="checkbox" label="featured" />
            <Form.Check className="mb-2" type="checkbox" label="featured" />
            <Form.Check type="checkbox" label="featured" />
          </Dropdown.Menu>
        </Dropdown>
      </Card.Header>
      <Card.Body>
        <Label className="badge-label" color="#DC3545">
          featured
        </Label>
      </Card.Body>
    </Card>
  );
};

export default Labels;

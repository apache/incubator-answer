import { FC } from 'react';
import { Form } from 'react-bootstrap';

interface Props {
  title: string;
}
const Index: FC<Props> = ({ title }) => {
  return <Form.Label>{title}</Form.Label>;
};

export default Index;

import { memo } from 'react';
import { Container, Card, Col, Carousel } from 'react-bootstrap';

const data = [
  {
    id: 1,
    url: require('@/assets/images/carousel-wecom-1.jpg'),
  },
  {
    id: 2,
    url: require('@/assets/images/carousel-wecom-2.jpg'),
  },
  {
    id: 3,
    url: require('@/assets/images/carousel-wecom-3.jpg'),
  },
  {
    id: 4,
    url: require('@/assets/images/carousel-wecom-4.jpg'),
  },
  {
    id: 5,
    url: require('@/assets/images/carousel-wecom-5.jpg'),
  },
];

const Index = () => {
  return (
    <Container>
      <Col lg={4} className="mx-auto mt-3 py-5">
        <Card>
          <Card.Body>
            <h3 className="text-center pt-3 mb-3">WeCome Login</h3>
            <p className="text-danger text-center">
              Login failed, please allow this app to access your email
              information before try again.
            </p>

            <Carousel controls={false}>
              {data.map((item) => (
                <Carousel.Item key={item.id}>
                  <img
                    className="d-block w-100"
                    src={item.url}
                    alt="First slide"
                  />
                </Carousel.Item>
              ))}
            </Carousel>
          </Card.Body>
        </Card>
      </Col>
    </Container>
  );
};

export default memo(Index);

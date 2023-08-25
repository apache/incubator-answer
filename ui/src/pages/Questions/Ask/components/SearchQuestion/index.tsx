import { memo } from 'react';
import { Accordion, ListGroup } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { Icon } from '@/components';
import { pathFactory } from '@/router/pathFactory';

import './index.scss';

const SearchQuestion = ({ similarQuestions }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'ask' });
  // set max similar number
  if (similarQuestions && similarQuestions.length > 5) {
    similarQuestions.length = 5;
  }
  return (
    <Accordion defaultActiveKey="0" className="search-question-wrap mt-3">
      <Accordion.Item eventKey="0" className="overflow-hidden">
        <Accordion.Button className="px-3 py-2 bg-light text-body">
          {t('similar_questions')}
        </Accordion.Button>

        <Accordion.Body className="p-0">
          <ListGroup variant="flush">
            {similarQuestions.map((item) => {
              return (
                <ListGroup.Item
                  action
                  as="a"
                  className="link-dark"
                  key={item.id}
                  href={pathFactory.questionLanding(item.id, item.url_title)}
                  target="_blank">
                  <span className="text-wrap text-break">
                    {item.title}
                    {item.status === 'closed'
                      ? ` [${t('closed', { keyPrefix: 'question' })}]`
                      : null}
                  </span>
                  {item.accepted_answer ? (
                    <span className="ms-3 text-success">
                      <Icon type="bi" name="check-circle-fill" />
                      <span className="ms-1">
                        {t('x_answers', {
                          keyPrefix: 'question',
                          count: item.answer_count,
                        })}
                      </span>
                    </span>
                  ) : (
                    item.answer_count > 0 && (
                      <span className="ms-3 text-secondary">
                        <Icon type="bi" name="chat-square-text-fill" />
                        <span className="ms-1">
                          {t('x_answers', {
                            keyPrefix: 'question',
                            count: item.answer_count,
                          })}
                        </span>
                      </span>
                    )
                  )}
                </ListGroup.Item>
              );
            })}
          </ListGroup>
        </Accordion.Body>
      </Accordion.Item>
    </Accordion>
  );
};

export default memo(SearchQuestion);

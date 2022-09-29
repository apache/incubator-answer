import React, { useState, useEffect, useRef } from 'react';
import { Container, Row, Col, Form, Button, Card } from 'react-bootstrap';
import { useParams, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import dayjs from 'dayjs';
import classNames from 'classnames';

import { Editor, EditorRef, Icon, PageTitle } from '@answer/components';
import {
  useQueryAnswerInfo,
  modifyAnswer,
  useQueryRevisions,
} from '@answer/api';
import type * as Type from '@/services/types';

import './index.scss';

interface FormDataItem {
  answer: {
    value: string;
    isInvalid: boolean;
    errorMsg: string;
  };
  description: {
    value: string;
    isInvalid: boolean;
    errorMsg: string;
  };
}
const initFormData = {
  answer: {
    value: '',
    isInvalid: false,
    errorMsg: '',
  },
  description: {
    value: '',
    isInvalid: false,
    errorMsg: '',
  },
};
const Ask = () => {
  const [formData, setFormData] = useState<FormDataItem>(initFormData);
  const { aid = '', qid = '' } = useParams();
  const [focusType, setForceType] = useState('');

  const { t } = useTranslation('translation', { keyPrefix: 'edit_answer' });
  const { t: t2 } = useTranslation('translation', { keyPrefix: 'dates' });
  const navigate = useNavigate();

  const { data } = useQueryAnswerInfo(aid);
  const { data: revisions = [] } = useQueryRevisions(aid);

  const editorRef = useRef<EditorRef>({
    getHtml: () => '',
  });

  const questionContentRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!data) {
      return;
    }
    formData.answer.value = data.info.content;
    setFormData({ ...formData });
  }, [data]);

  const handleAnswerChange = (value: string) =>
    setFormData({
      ...formData,
      answer: { ...formData.answer, value },
    });

  const checkValidated = (): boolean => {
    let bol = true;
    const { answer } = formData;

    if (!answer.value) {
      bol = false;
      formData.answer = {
        value: '',
        isInvalid: true,
        errorMsg: '标题不能为空',
      };
    } else {
      formData.answer = {
        value: answer.value,
        isInvalid: false,
        errorMsg: '',
      };
    }

    setFormData({
      ...formData,
    });
    return bol;
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();
    if (!checkValidated()) {
      return;
    }

    const params: Type.AnswerParams = {
      content: formData.answer.value,
      html: editorRef.current.getHtml(),
      question_id: qid,
      id: aid,
    };
    modifyAnswer(params).then(() => {
      window.location.href = `/questions/${qid}/${aid}`;
    });
  };

  const handleSelectedRevision = (e) => {
    const index = e.target.value;
    const revision = revisions[index];
    formData.answer.value = revision.content.content;
    setFormData({ ...formData });
  };

  const backPage = () => {
    navigate(-1);
  };

  return (
    <>
      <PageTitle title={t('edit_answer', { keyPrefix: 'page_title' })} />
      <Container className="pt-4 mt-2 mb-5 edit-answer-wrap">
        <Row className="justify-content-center">
          <Col sm={12} md={10}>
            <h3 className="mb-4">{t('title')}</h3>
          </Col>
        </Row>
        <Row className="justify-content-center">
          <Col sm={12} md={7} className="mb-4 mb-md-0">
            <a href={`/questions/${qid}`} target="_blank" rel="noreferrer">
              <h5 className="mb-3">{data?.question.title}</h5>
            </a>

            <div className="content-wrap">
              <div
                ref={questionContentRef}
                className="content position-absolute top-0 w-100"
                dangerouslySetInnerHTML={{ __html: data?.question.html }}
              />
              <div
                className="resize-bottom"
                style={{ maxHeight: questionContentRef?.current?.scrollHeight }}
              />
              <div className="line bg-light  d-flex justify-content-center align-items-center">
                <Icon name="three-dots" />
              </div>
            </div>

            <Form noValidate onSubmit={handleSubmit}>
              <Form.Group controlId="revision" className="mb-3">
                <Form.Label>{t('form.fields.revision.label')}</Form.Label>
                <Form.Select onChange={handleSelectedRevision}>
                  {revisions.map(({ create_at, reason }, index) => {
                    const date = dayjs(create_at * 1000).format(
                      t2('long_date_with_time'),
                    );
                    return (
                      <option key={`${create_at}`} value={index}>
                        {`${date} - robin - ${reason || t('default_reason')}`}
                      </option>
                    );
                  })}
                </Form.Select>
              </Form.Group>

              <Form.Group controlId="answer" className="mt-4">
                <Form.Label>{t('form.fields.answer.label')}</Form.Label>
                <Editor
                  value={formData.answer.value}
                  onChange={handleAnswerChange}
                  className={classNames(
                    'form-control p-0',
                    focusType === 'answer' && 'focus',
                  )}
                  onFocus={() => {
                    setForceType('answer');
                  }}
                  onBlur={() => {
                    setForceType('');
                  }}
                  ref={editorRef}
                />
                <Form.Control
                  value={formData.answer.value}
                  type="text"
                  isInvalid={formData.answer.isInvalid}
                  readOnly
                  hidden
                />
                <Form.Control.Feedback type="invalid">
                  {formData.answer.errorMsg}
                </Form.Control.Feedback>
              </Form.Group>
              <Form.Group controlId="edit_summary" className="my-3">
                <Form.Label>{t('form.fields.edit_summary.label')}</Form.Label>
                <Form.Control
                  type="text"
                  defaultValue={formData.description.value}
                  isInvalid={formData.description.isInvalid}
                  placeholder={t('form.fields.edit_summary.placeholder')}
                />
                <Form.Control.Feedback type="invalid">
                  {formData.description.errorMsg}
                </Form.Control.Feedback>
              </Form.Group>

              <div className="mt-3">
                <Button type="submit" className="me-2">
                  {t('btn_save_edits')}
                </Button>
                <Button variant="link" onClick={backPage}>
                  {t('btn_cancel')}
                </Button>
              </div>
            </Form>
          </Col>
          <Col sm={12} md={3}>
            <Card className="mb-4">
              <Card.Header>{t('how_to_ask.title')}</Card.Header>
              <Card.Body>
                <Card.Text>{t('how_to_ask.description')}</Card.Text>
              </Card.Body>
            </Card>
            <Card className="mb-4">
              <Card.Header>{t('how_to_format.title')}</Card.Header>
              <Card.Body>
                <Card.Text>{t('how_to_format.description')}</Card.Text>
              </Card.Body>
            </Card>
            <Card>
              <Card.Header>{t('how_to_tag.title')}</Card.Header>
              <Card.Body>
                <Card.Text>{t('how_to_tag.description')}</Card.Text>
                <ul className="mb-0">
                  {Array.from(
                    t('how_to_tag.tips', { returnObjects: true }) as string[],
                  ).map((item) => {
                    return <li>{item}</li>;
                  })}
                </ul>
              </Card.Body>
            </Card>
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default Ask;

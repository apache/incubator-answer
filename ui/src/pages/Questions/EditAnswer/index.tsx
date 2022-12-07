import React, { useState, useEffect, useRef } from 'react';
import { Container, Row, Col, Form, Button, Card } from 'react-bootstrap';
import { useParams, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import dayjs from 'dayjs';
import classNames from 'classnames';

import { usePageTags } from '@/hooks';
import { pathFactory } from '@/router/pathFactory';
import { Editor, EditorRef, Icon } from '@/components';
import type * as Type from '@/common/interface';
import {
  useQueryAnswerInfo,
  modifyAnswer,
  useQueryRevisions,
} from '@/services';

import './index.scss';

interface FormDataItem {
  answer: Type.FormValue<string>;
  description: Type.FormValue<string>;
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
  const handleSummaryChange = (evt) => {
    const v = evt.currentTarget.value;
    setFormData({
      ...formData,
      description: { ...formData.description, value: v },
    });
  };

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
      edit_summary: formData.description.value,
    };
    modifyAnswer(params).then((res) => {
      navigate(
        pathFactory.answerLanding({
          questionId: qid,
          questionTitle: data?.question?.title,
          answerId: aid,
        }),
        {
          state: { isReview: res?.wait_for_review },
        },
      );
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
  usePageTags({
    title: t('edit_answer', { keyPrefix: 'page_title' }),
  });
  return (
    <Container className="pt-4 mt-2 mb-5 edit-answer-wrap">
      <Row className="justify-content-center">
        <Col xxl={10} md={12}>
          <h3 className="mb-4">{t('title')}</h3>
        </Col>
      </Row>
      <Row className="justify-content-center">
        <Col xxl={7} lg={8} sm={12} className="mb-4 mb-md-0">
          <a
            href={pathFactory.questionLanding(qid, data?.question.title)}
            target="_blank"
            rel="noreferrer">
            <h5 className="mb-3">{data?.question.title}</h5>
          </a>

          <div className="question-content-wrap">
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
                {revisions.map(({ create_at, reason, user_info }, index) => {
                  const date = dayjs(create_at * 1000)
                    .tz()
                    .format(t('long_date_with_time', { keyPrefix: 'dates' }));
                  return (
                    <option key={`${create_at}`} value={index}>
                      {`${date} - ${user_info.display_name} - ${
                        reason || t('default_reason')
                      }`}
                    </option>
                  );
                })}
              </Form.Select>
            </Form.Group>

            <Form.Group controlId="answer" className="mt-3">
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
                onChange={handleSummaryChange}
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
        <Col xxl={3} lg={4} sm={12} className="mt-5 mt-lg-0">
          <Card>
            <Card.Header>
              {t('title', { keyPrefix: 'how_to_format' })}
            </Card.Header>
            <Card.Body
              className="fmt small"
              dangerouslySetInnerHTML={{
                __html: t('description', { keyPrefix: 'how_to_format' }),
              }}
            />
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Ask;

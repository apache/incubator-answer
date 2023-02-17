import React, { useState, useEffect, useRef } from 'react';
import { Container, Row, Col, Form, Button, Card } from 'react-bootstrap';
import { useParams, useNavigate, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import dayjs from 'dayjs';
import classNames from 'classnames';
import { isEqual } from 'lodash';

import { usePageTags, usePromptWithUnload } from '@/hooks';
import { Editor, EditorRef, TagSelector } from '@/components';
import type * as Type from '@/common/interface';
import {
  saveQuestion,
  questionDetail,
  modifyQuestion,
  useQueryRevisions,
  postAnswer,
  useQueryQuestionByTitle,
  getTagsBySlugName,
} from '@/services';
import { handleFormError } from '@/utils';
import { pathFactory } from '@/router/pathFactory';

import SearchQuestion from './components/SearchQuestion';

interface FormDataItem {
  title: Type.FormValue<string>;
  tags: Type.FormValue<Type.Tag[]>;
  content: Type.FormValue<string>;
  answer: Type.FormValue<string>;
  edit_summary: Type.FormValue<string>;
}

const Ask = () => {
  const initFormData = {
    title: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    tags: {
      value: [],
      isInvalid: false,
      errorMsg: '',
    },
    content: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    answer: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    edit_summary: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  };
  const { t } = useTranslation('translation', { keyPrefix: 'ask' });
  const [formData, setFormData] = useState<FormDataItem>(initFormData);
  const [immData, setImmData] = useState<FormDataItem>(initFormData);
  const [checked, setCheckState] = useState(false);
  const [contentChanged, setContentChanged] = useState(false);
  const [focusType, setForceType] = useState('');
  const resetForm = () => {
    setFormData(initFormData);
    setCheckState(false);
    setForceType('');
  };

  const editorRef = useRef<EditorRef>({
    getHtml: () => '',
  });
  const editorRef2 = useRef<EditorRef>({
    getHtml: () => '',
  });

  const { qid } = useParams();
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const initQueryTags = () => {
    const queryTags = searchParams.get('tags');
    if (!queryTags) {
      return;
    }
    getTagsBySlugName(queryTags).then((tags) => {
      // eslint-disable-next-line @typescript-eslint/no-use-before-define
      handleTagsChange(tags);
    });
  };

  const isEdit = qid !== undefined;
  const { data: similarQuestions = { list: [] } } = useQueryQuestionByTitle(
    isEdit ? '' : formData.title.value,
  );

  useEffect(() => {
    const { title, tags, content, answer } = formData;
    const { title: editTitle, tags: editTags, content: editContent } = immData;

    // edited
    if (qid) {
      if (
        editTitle.value !== title.value ||
        editContent.value !== content.value ||
        !isEqual(
          editTags.value.map((v) => v.slug_name),
          tags.value.map((v) => v.slug_name),
        )
      ) {
        setContentChanged(true);
      } else {
        setContentChanged(false);
      }
      return;
    }

    // write
    if (title.value || tags.value.length > 0 || content.value || answer.value) {
      setContentChanged(true);
    } else {
      setContentChanged(false);
    }
  }, [formData]);

  usePromptWithUnload({
    when: contentChanged,
  });

  useEffect(() => {
    if (!isEdit) {
      resetForm();
      initQueryTags();
    }
  }, [isEdit]);
  const { data: revisions = [] } = useQueryRevisions(qid);

  useEffect(() => {
    if (!isEdit) {
      return;
    }
    questionDetail(qid).then((res) => {
      formData.title.value = res.title;
      formData.content.value = res.content;
      formData.tags.value = res.tags.map((item) => {
        return {
          ...item,
          parsed_text: '',
          original_text: '',
        };
      });
      setImmData({ ...formData });
      setFormData({ ...formData });
    });
  }, [qid]);

  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      title: { ...formData.title, value: e.currentTarget.value, errorMsg: '' },
    });
  };
  const handleContentChange = (value: string) => {
    setFormData({
      ...formData,
      content: { ...formData.content, value, errorMsg: '' },
    });
  };
  const handleTagsChange = (value) =>
    setFormData({
      ...formData,
      tags: { ...formData.tags, value, errorMsg: '' },
    });

  const handleAnswerChange = (value: string) =>
    setFormData({
      ...formData,
      answer: { ...formData.answer, value, errorMsg: '' },
    });

  const handleSummaryChange = (evt: React.ChangeEvent<HTMLInputElement>) =>
    setFormData({
      ...formData,
      edit_summary: {
        ...formData.edit_summary,
        value: evt.currentTarget.value,
      },
    });

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    setContentChanged(false);
    event.preventDefault();
    event.stopPropagation();

    const params: Type.QuestionParams = {
      title: formData.title.value,
      content: formData.content.value,
      tags: formData.tags.value,
    };
    if (isEdit) {
      modifyQuestion({
        ...params,
        id: qid,
        edit_summary: formData.edit_summary.value,
      })
        .then((res) => {
          navigate(pathFactory.questionLanding(qid, params.url_title), {
            state: { isReview: res?.wait_for_review },
          });
        })
        .catch((err) => {
          if (err.isError) {
            const data = handleFormError(err, formData);
            setFormData({ ...data });
          }
        });
    } else {
      const res = await saveQuestion(params).catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          setFormData({ ...data });
        }
      });

      const id = res?.id;
      if (id) {
        if (checked) {
          postAnswer({
            question_id: id,
            content: formData.answer.value,
          })
            .then(() => {
              navigate(pathFactory.questionLanding(id, params.url_title));
            })
            .catch((err) => {
              if (err.isError) {
                const data = handleFormError(err, formData, [
                  { from: 'content', to: 'answer' },
                ]);
                setFormData({ ...data });
              }
            });
        } else {
          navigate(pathFactory.questionLanding(id));
        }
      }
    }
  };
  const backPage = () => {
    navigate(-1);
  };

  const handleSelectedRevision = (e) => {
    const index = e.target.value;
    const revision = revisions[index];
    formData.content.value = revision.content.content;
    setImmData({ ...formData });
    setFormData({ ...formData });
  };
  const bool = similarQuestions.length > 0 && !isEdit;
  let pageTitle = t('ask_a_question', { keyPrefix: 'page_title' });
  if (isEdit) {
    pageTitle = t('edit_question', { keyPrefix: 'page_title' });
  }
  usePageTags({
    title: pageTitle,
  });
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col xxl={10} md={12}>
          <h3 className="mb-4">{isEdit ? t('edit_title') : t('title')}</h3>
        </Col>
      </Row>
      <Row className="justify-content-center">
        <Col xxl={7} lg={8} sm={12} className="mb-4 mb-md-0">
          <Form noValidate onSubmit={handleSubmit}>
            {isEdit && (
              <Form.Group controlId="revision" className="mb-3">
                <Form.Label>{t('form.fields.revision.label')}</Form.Label>
                <Form.Select onChange={handleSelectedRevision}>
                  {revisions.map(({ reason, create_at, user_info }, index) => {
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
            )}

            <Form.Group controlId="title" className="mb-3">
              <Form.Label>{t('form.fields.title.label')}</Form.Label>
              <Form.Control
                value={formData.title.value}
                isInvalid={formData.title.isInvalid}
                onChange={handleTitleChange}
                placeholder={t('form.fields.title.placeholder')}
                autoFocus
              />

              <Form.Control.Feedback type="invalid">
                {formData.title.errorMsg}
              </Form.Control.Feedback>
              {bool && <SearchQuestion similarQuestions={similarQuestions} />}
            </Form.Group>
            <Form.Group controlId="body">
              <Form.Label>{t('form.fields.body.label')}</Form.Label>
              <Form.Control
                defaultValue={formData.content.value}
                isInvalid={formData.content.isInvalid}
                hidden
              />
              <Editor
                value={formData.content.value}
                onChange={handleContentChange}
                className={classNames(
                  'form-control p-0',
                  focusType === 'content' && 'focus',
                )}
                onFocus={() => {
                  setForceType('content');
                }}
                onBlur={() => {
                  setForceType('');
                }}
                ref={editorRef}
              />
              <Form.Control.Feedback type="invalid">
                {formData.content.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>
            <Form.Group controlId="tags" className="my-3">
              <Form.Label>{t('form.fields.tags.label')}</Form.Label>
              <Form.Control
                defaultValue={JSON.stringify(formData.tags.value)}
                isInvalid={formData.tags.isInvalid}
                hidden
              />
              <TagSelector
                value={formData.tags.value}
                onChange={handleTagsChange}
                showRequiredTagText
              />
              <Form.Control.Feedback type="invalid">
                {formData.tags.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>
            {isEdit && (
              <Form.Group controlId="edit_summary" className="my-3">
                <Form.Label>{t('form.fields.edit_summary.label')}</Form.Label>
                <Form.Control
                  type="text"
                  defaultValue={formData.edit_summary.value}
                  isInvalid={formData.edit_summary.isInvalid}
                  placeholder={t('form.fields.edit_summary.placeholder')}
                  onChange={handleSummaryChange}
                />
                <Form.Control.Feedback type="invalid">
                  {formData.edit_summary.errorMsg}
                </Form.Control.Feedback>
              </Form.Group>
            )}
            {!checked && (
              <div className="mt-3">
                <Button type="submit" className="me-2">
                  {isEdit ? t('btn_save_edits') : t('btn_post_question')}
                </Button>

                <Button variant="link" onClick={backPage}>
                  {t('cancel', { keyPrefix: 'btns' })}
                </Button>
              </div>
            )}
            {!isEdit && (
              <>
                <Form.Check
                  className="mt-5"
                  checked={checked}
                  type="checkbox"
                  label={t('answer_question')}
                  onChange={(e) => setCheckState(e.target.checked)}
                  id="radio-answer"
                />
                {checked && (
                  <Form.Group controlId="answer" className="mt-4">
                    <Form.Label>{t('form.fields.answer.label')}</Form.Label>
                    <Editor
                      value={formData.answer.value}
                      onChange={handleAnswerChange}
                      ref={editorRef2}
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
                    />
                    <Form.Control
                      value={formData.answer.value}
                      type="text"
                      isInvalid={formData.answer.isInvalid}
                      hidden
                    />
                    <Form.Control.Feedback type="invalid">
                      {formData.answer.errorMsg}
                    </Form.Control.Feedback>
                  </Form.Group>
                )}
              </>
            )}
            {checked && (
              <Button type="submit" className="mt-3">
                {t('post_question&answer')}
              </Button>
            )}
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
                __html: t('desc', { keyPrefix: 'how_to_format' }),
              }}
            />
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Ask;

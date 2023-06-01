import React, { useState, useRef, useEffect } from 'react';
import { Row, Col, Form, Button, Card } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { usePageTags, usePromptWithUnload } from '@/hooks';
import { Editor, EditorRef } from '@/components';
import { loggedUserInfoStore } from '@/stores';
import type * as Type from '@/common/interface';
import { createTag } from '@/services';
import { handleFormError } from '@/utils';

interface FormDataItem {
  displayName: Type.FormValue<string>;
  slugName: Type.FormValue<string>;
  description: Type.FormValue<string>;
}

const Index = () => {
  const initFormData = {
    displayName: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    slugName: {
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
  const { role_id = 1 } = loggedUserInfoStore((state) => state.user);
  const navigate = useNavigate();
  const { t } = useTranslation('translation', { keyPrefix: 'tag_modal' });
  const [focusType, setForceType] = useState('');

  const [formData, setFormData] = useState<FormDataItem>(initFormData);
  const [immData] = useState(initFormData);
  const [contentChanged, setContentChanged] = useState(false);

  const editorRef = useRef<EditorRef>({
    getHtml: () => '',
  });

  usePromptWithUnload({
    when: contentChanged,
  });

  useEffect(() => {
    const { displayName, slugName, description } = formData;
    const {
      displayName: display_name,
      slugName: slug_name,
      description: original_text,
    } = immData;
    if (!display_name || !slug_name || !original_text) {
      return;
    }

    if (
      display_name.value !== displayName.value ||
      slug_name.value !== slugName.value ||
      original_text.value !== description.value
    ) {
      setContentChanged(true);
    } else {
      setContentChanged(false);
    }
  }, [
    formData.displayName.value,
    formData.slugName.value,
    formData.description.value,
  ]);

  const handleDescriptionChange = (value: string) =>
    setFormData({
      ...formData,
      description: { ...formData.description, value, isInvalid: false },
    });

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();
    setContentChanged(false);
    const params = {
      display_name: formData.displayName.value,
      slug_name: formData.slugName.value,
      original_text: formData.description.value,
    };
    createTag(params)
      .then((res) => {
        navigate(`/tags/${res.slug_name}/info`, {
          replace: true,
        });
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData, [
            { from: 'display_name', to: 'displayName' },
            { from: 'slug_name', to: 'slugName' },
            { from: 'original_text', to: 'description' },
          ]);
          setFormData({ ...data });
        }
      });
  };

  const handleDisplayNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      displayName: {
        ...formData.displayName,
        value: e.currentTarget.value,
        isInvalid: false,
      },
    });
  };

  const handleSlugNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      slugName: {
        ...formData.slugName,
        value: e.currentTarget.value,
        isInvalid: false,
      },
    });
  };

  usePageTags({
    title: t('create_tag', { keyPrefix: 'page_title' }),
  });

  return (
    <div className="pt-4 mb-5">
      <h3 className="mb-4">{t('title')}</h3>
      <Row>
        <Col className="page-main flex-auto">
          <Form noValidate onSubmit={handleSubmit}>
            <Form.Group controlId="display_name" className="mb-3">
              <Form.Label>{t('form.fields.display_name.label')}</Form.Label>
              <Form.Control
                value={formData.displayName.value}
                isInvalid={formData.displayName.isInvalid}
                disabled={role_id !== 2 && role_id !== 3}
                onChange={handleDisplayNameChange}
              />

              <Form.Control.Feedback type="invalid">
                {formData.displayName.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>
            <Form.Group controlId="slug_name" className="mb-3">
              <Form.Label>{t('form.fields.slug_name.label')}</Form.Label>
              <Form.Control
                value={formData.slugName.value}
                isInvalid={formData.slugName.isInvalid}
                disabled={role_id !== 2 && role_id !== 3}
                onChange={handleSlugNameChange}
              />
              <Form.Text as="div">{t('form.fields.slug_name.desc')}</Form.Text>
              <Form.Control.Feedback type="invalid">
                {formData.slugName.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>

            <Form.Group controlId="description" className="mt-4">
              <Form.Label>{t('form.fields.desc.label')}</Form.Label>
              <Editor
                value={formData.description.value}
                onChange={handleDescriptionChange}
                className={classNames(
                  'form-control p-0',
                  focusType === 'description' && 'focus',
                )}
                onFocus={() => {
                  setForceType('description');
                }}
                onBlur={() => {
                  setForceType('');
                }}
                ref={editorRef}
              />
              <Form.Control
                value={formData.description.value}
                type="text"
                isInvalid={formData.description.isInvalid}
                readOnly
                hidden
              />
              <Form.Control.Feedback type="invalid">
                {formData.description.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>
            <div className="mt-3">
              <Button type="submit">{t('btn_post')}</Button>
            </div>
          </Form>
        </Col>
        <Col className="page-right-side mt-4 mt-xl-0">
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
    </div>
  );
};

export default Index;

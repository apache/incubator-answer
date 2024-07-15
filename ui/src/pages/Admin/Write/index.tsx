/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Form, Button } from 'react-bootstrap';

import { TagSelector } from '@/components';
import type * as Type from '@/common/interface';
import { useToast } from '@/hooks';
import {
  getRequireAndReservedTag,
  postRequireAndReservedTag,
} from '@/services';
import { handleFormError, scrollToElementTop } from '@/utils';
import { writeSettingStore } from '@/stores';

const initFormData = {
  reserved_tags: {
    value: [] as Type.Tag[], // Replace `Type.Tag` with the correct type for `reserved_tags.value`
    errorMsg: '',
    isInvalid: false,
  },
  recommend_tags: {
    value: [] as Type.Tag[],
    errorMsg: '',
    isInvalid: false,
  },
  required_tag: {
    value: false,
    errorMsg: '',
    isInvalid: false,
  },
  restrict_answer: {
    value: false,
    errorMsg: '',
    isInvalid: false,
  },
};

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.write',
  });
  const Toast = useToast();

  const [formData, setFormData] = useState(initFormData);

  const handleValueChange = (value) => {
    setFormData({
      ...formData,
      ...value,
    });
  };

  const checkValidated = (): boolean => {
    let bol = true;
    const { recommend_tags, reserved_tags } = formData;
    // 找出 recommend_tags 和 reserved_tags 中是否有重复的标签
    // 通过标签中的 slug_name 来去重
    const repeatTag = recommend_tags.value.filter((tag) =>
      reserved_tags.value.some((rTag) => rTag?.slug_name === tag?.slug_name),
    );
    if (repeatTag.length > 0) {
      handleValueChange({
        recommend_tags: {
          ...recommend_tags,
          errorMsg: t('recommend_tags.msg.contain_reserved'),
          isInvalid: true,
        },
      });
      bol = false;
      const ele = document.getElementById('recommend_tags');
      scrollToElementTop(ele);
    } else {
      handleValueChange({
        recommend_tags: {
          ...recommend_tags,
          errorMsg: '',
          isInvalid: false,
        },
      });
    }
    return bol;
  };

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();
    if (!checkValidated()) {
      return;
    }
    const reqParams: Type.AdminSettingsWrite = {
      recommend_tags: formData.recommend_tags.value,
      reserved_tags: formData.reserved_tags.value,
      required_tag: formData.required_tag.value,
      restrict_answer: formData.restrict_answer.value,
    };
    postRequireAndReservedTag(reqParams)
      .then(() => {
        Toast.onShow({
          msg: t('update', { keyPrefix: 'toast' }),
          variant: 'success',
        });
        writeSettingStore
          .getState()
          .update({ restrict_answer: reqParams.restrict_answer });
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          setFormData({ ...data });
          const ele = document.getElementById(err.list[0].error_field);
          scrollToElementTop(ele);
        }
      });
  };

  const initData = () => {
    getRequireAndReservedTag().then((res) => {
      if (Array.isArray(res.recommend_tags)) {
        formData.recommend_tags.value = res.recommend_tags;
      }
      formData.required_tag.value = res.required_tag;
      formData.restrict_answer.value = res.restrict_answer;
      if (Array.isArray(res.reserved_tags)) {
        formData.reserved_tags.value = res.reserved_tags;
      }
      setFormData({ ...formData });
    });
  };

  useEffect(() => {
    initData();
  }, []);

  // const handleOnChange = (data) => {
  //   setFormData(data);
  // };

  return (
    <>
      <h3 className="mb-4">{t('page_title')}</h3>
      <Form noValidate onSubmit={onSubmit}>
        <Form.Group className="mb-3" controlId="reserved_tags">
          <Form.Label>{t('reserved_tags.label')}</Form.Label>
          <TagSelector
            value={formData.reserved_tags.value}
            onChange={(val) => {
              handleValueChange({
                reserved_tags: {
                  value: val,
                  errorMsg: '',
                  isInvalid: false,
                },
              });
            }}
            showRequiredTag={false}
            maxTagLength={0}
            tagStyleMode="simple"
            formText={t('reserved_tags.text')}
            isInvalid={formData.reserved_tags.isInvalid}
            errMsg={formData.reserved_tags.errorMsg}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="recommend_tags">
          <Form.Label>{t('recommend_tags.label')}</Form.Label>
          <TagSelector
            value={formData.recommend_tags.value}
            onChange={(val) => {
              handleValueChange({
                recommend_tags: {
                  value: val,
                  errorMsg: '',
                  isInvalid: false,
                },
              });
            }}
            showRequiredTag={false}
            tagStyleMode="simple"
            formText={t('recommend_tags.text')}
            isInvalid={formData.recommend_tags.isInvalid}
            errMsg={formData.recommend_tags.errorMsg}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="required_tag">
          <Form.Label>{t('required_tag.title')}</Form.Label>
          <Form.Switch
            label={t('required_tag.label')}
            checked={formData.required_tag.value}
            onChange={(evt) => {
              handleValueChange({
                required_tag: {
                  value: evt.target.checked,
                  errorMsg: '',
                  isInvalid: false,
                },
              });
            }}
          />
          <Form.Text>{t('required_tag.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.required_tag.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>

        <Form.Group className="mb-3" controlId="restrict_answer">
          <Form.Label>{t('restrict_answer.title')}</Form.Label>
          <Form.Switch
            label={t('restrict_answer.label')}
            checked={formData.restrict_answer.value}
            onChange={(evt) => {
              handleValueChange({
                restrict_answer: {
                  value: evt.target.checked,
                  errorMsg: '',
                  isInvalid: false,
                },
              });
            }}
          />
          <Form.Text>{t('restrict_answer.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.restrict_answer.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>

        <Form.Group className="mb-3">
          <Button type="submit">{t('save', { keyPrefix: 'btns' })}</Button>
        </Form.Group>
      </Form>
    </>
  );
};

export default Index;

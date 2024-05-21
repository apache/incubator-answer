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

import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import type * as Type from '@/common/interface';
import { useToast } from '@/hooks';
import {
  getRequireAndReservedTag,
  postRequireAndReservedTag,
} from '@/services';
import { handleFormError, scrollToElementTop } from '@/utils';
import { writeSettingStore } from '@/stores';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.write',
  });
  const Toast = useToast();

  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      restrict_answer: {
        type: 'boolean',
        title: t('restrict_answer.title'),
        description: t('restrict_answer.text'),
        default: true,
      },
      recommend_tags: {
        type: 'string',
        title: t('recommend_tags.label'),
        description: t('recommend_tags.text'),
      },
      required_tag: {
        type: 'boolean',
        title: t('required_tag.title'),
        description: t('required_tag.text'),
      },
      reserved_tags: {
        type: 'string',
        title: t('reserved_tags.label'),
        description: t('reserved_tags.text'),
      },
    },
  };
  const uiSchema: UISchema = {
    restrict_answer: {
      'ui:widget': 'switch',
      'ui:options': {
        label: t('restrict_answer.label'),
      },
    },
    recommend_tags: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 10,
      },
    },
    required_tag: {
      'ui:widget': 'switch',
      'ui:options': {
        label: t('required_tag.label'),
      },
    },
    reserved_tags: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 10,
      },
    },
  };
  const [formData, setFormData] = useState(initFormData(schema));

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();
    let recommend_tags = [];
    if (formData.recommend_tags.value?.trim()) {
      recommend_tags = formData.recommend_tags.value.trim().split('\n');
    }
    let reserved_tags = [];
    if (formData.reserved_tags.value?.trim()) {
      reserved_tags = formData.reserved_tags.value.trim().split('\n');
    }
    const reqParams: Type.AdminSettingsWrite = {
      recommend_tags,
      reserved_tags,
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
        formData.recommend_tags.value = res.recommend_tags.join('\n');
      }
      formData.required_tag.value = res.required_tag;
      formData.restrict_answer.value = res.restrict_answer;
      if (Array.isArray(res.reserved_tags)) {
        formData.reserved_tags.value = res.reserved_tags.join('\n');
      }
      setFormData({ ...formData });
    });
  };

  useEffect(() => {
    initData();
  }, []);

  const handleOnChange = (data) => {
    setFormData(data);
  };

  return (
    <>
      <h3 className="mb-4">{t('page_title')}</h3>
      <SchemaForm
        schema={schema}
        formData={formData}
        onSubmit={onSubmit}
        uiSchema={uiSchema}
        onChange={handleOnChange}
      />
    </>
  );
};

export default Index;

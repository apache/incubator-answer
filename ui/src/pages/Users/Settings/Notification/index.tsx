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

import React, { useState, FormEvent, useEffect } from 'react';
import { useTranslation } from 'react-i18next';

import type { FormDataType, NotificationConfig } from '@/common/interface';
import { useToast } from '@/hooks';
import { useGetNotificationConfig, putNotificationConfig } from '@/services';
import { SchemaForm, JSONSchema, UISchema, initFormData } from '@/components';

const Index = () => {
  const toast = useToast();
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.notification',
  });
  const { data: configData } = useGetNotificationConfig();

  const schema: JSONSchema = {
    title: t('heading'),
    properties: {
      inbox: {
        type: 'boolean',
        title: t('inbox.label'),
        description: t('inbox.description'),
        enum: configData?.inbox?.map((v) => v.enable),
        default: configData?.inbox?.map((v) => v.enable),
        enumNames: configData?.inbox?.map((v) => t(v.key)),
      },
      all_new_question: {
        type: 'boolean',
        title: t('all_new_question.label'),
        description: t('all_new_question.description'),
        enum: configData?.all_new_question?.map((v) => v.enable),
        default: configData?.all_new_question?.map((v) => v.enable),
        enumNames: configData?.all_new_question?.map((v) => t(v.key)),
      },
      all_new_question_for_following_tags: {
        type: 'boolean',
        title: t('all_new_question_for_following_tags.label'),
        description: t('all_new_question_for_following_tags.description'),
        enum: configData?.all_new_question_for_following_tags?.map(
          (v) => v.enable,
        ),
        default: configData?.all_new_question_for_following_tags?.map(
          (v) => v.enable,
        ),
        enumNames: configData?.all_new_question_for_following_tags?.map((v) =>
          t(v.key),
        ),
      },
    },
  };
  const uiSchema: UISchema = {
    inbox: {
      'ui:widget': 'checkbox',
      'ui:options': {
        label: t('email'),
      },
    },
    all_new_question: {
      'ui:widget': 'checkbox',
      'ui:options': {
        label: t('email'),
      },
    },
    all_new_question_for_following_tags: {
      'ui:widget': 'checkbox',
      'ui:options': {
        label: t('email'),
        text: t('all_new_question_for_following_tags.description'),
      },
    },
  };
  const [formData, setFormData] = useState<FormDataType>(initFormData(schema));

  useEffect(() => {
    setFormData(initFormData(schema));
  }, [configData]);

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    const params = {
      inbox: configData?.inbox.map((v, index) => {
        return { enable: formData.inbox.value[index], key: v.key };
      }),
      all_new_question: configData?.all_new_question.map((v, index) => {
        return { enable: formData.all_new_question.value[index], key: v.key };
      }),
      all_new_question_for_following_tags:
        configData?.all_new_question_for_following_tags.map((v, index) => {
          return {
            enable: formData.all_new_question_for_following_tags.value[index],
            key: v.key,
          };
        }),
    } as NotificationConfig;

    putNotificationConfig(params).then(() => {
      toast.onShow({
        msg: t('update', { keyPrefix: 'toast' }),
        variant: 'success',
      });
    });
  };

  const handleChange = (ud) => {
    setFormData(ud);
  };
  return (
    <>
      <h3 className="mb-4">{t('heading')}</h3>
      <SchemaForm
        schema={schema}
        uiSchema={uiSchema}
        formData={formData}
        onChange={handleChange}
        onSubmit={handleSubmit}
      />
    </>
  );
};

export default React.memo(Index);

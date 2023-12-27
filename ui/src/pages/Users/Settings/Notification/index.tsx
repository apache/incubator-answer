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
        default: configData?.inbox.enable,
      },
      all_new_question: {
        type: 'boolean',
        title: t('all_new_question.label'),
        description: t('all_new_question.description'),
        default: configData?.all_new_question.enable,
      },
      all_new_question_for_following_tags: {
        type: 'boolean',
        title: t('all_new_question_for_following_tags.label'),
        description: t('all_new_question_for_following_tags.description'),
        default: configData?.all_new_question_for_following_tags.enable,
      },
    },
  };
  const uiSchema: UISchema = {
    inbox: {
      'ui:widget': 'switch',
      'ui:options': {
        label: t('turn_on'),
      },
    },
    all_new_question: {
      'ui:widget': 'switch',
      'ui:options': {
        label: t('turn_on'),
      },
    },
    all_new_question_for_following_tags: {
      'ui:widget': 'switch',
      'ui:options': {
        label: t('turn_on'),
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
      inbox: {
        enable: formData.inbox.value,
        key: configData?.inbox.key,
      },
      all_new_question: {
        enable: formData.all_new_question.value,
        key: configData?.all_new_question.key,
      },
      all_new_question_for_following_tags: {
        enable: formData.all_new_question_for_following_tags.value,
        key: configData?.all_new_question_for_following_tags.key,
      },
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

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

import { FC, memo, useState } from 'react';
import { Card, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { TagSelector, Tag } from '@/components';
import { tryLoggedAndActivated } from '@/utils/guard';
import { useFollowingTags, followTags } from '@/services';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'question' });
  const [isEdit, setEditState] = useState(false);
  const { data: followingTags, mutate } = useFollowingTags();

  const newTags: any = followingTags?.map((item) => {
    if (item.slug_name) {
      return item.slug_name;
    }
    return item;
  });

  const handleFollowTags = () => {
    followTags({
      slug_name_list: newTags,
    });
    setEditState(false);
  };

  const handleTagsChange = (value) => {
    mutate([...value], {
      revalidate: false,
    });
  };

  if (!tryLoggedAndActivated().ok) {
    return null;
  }
  return isEdit ? (
    <Card className="mb-4">
      <Card.Header className="text-nowrap d-flex justify-content-between">
        {t('following_tags')}
        <Button
          variant="link"
          className="p-0 m-0 btn-no-border"
          onClick={handleFollowTags}>
          {t('save')}
        </Button>
      </Card.Header>
      <Card.Body>
        <TagSelector
          value={followingTags}
          onChange={handleTagsChange}
          hiddenDescription
          hiddenCreateBtn
          autoFocus
        />
      </Card.Body>
    </Card>
  ) : (
    <Card className="mb-4">
      <Card.Header className="text-nowrap d-flex justify-content-between text-capitalize">
        {t('following_tags')}
        <Button
          variant="link"
          className="p-0 btn-no-border text-capitalize"
          onClick={() => setEditState(true)}>
          {t('edit')}
        </Button>
      </Card.Header>
      <Card.Body>
        {followingTags?.length ? (
          <div className="m-n1">
            {followingTags.map((item) => {
              const slugName = item?.slug_name;
              return <Tag key={slugName} className="m-1" data={item} />;
            })}
          </div>
        ) : (
          <>
            <div className="text-muted">{t('follow_tag_tip')}</div>
            <Button
              size="sm"
              className="mt-3"
              variant="outline-primary"
              onClick={() => setEditState(true)}>
              {t('follow_a_tag')}
            </Button>
          </>
        )}
      </Card.Body>
    </Card>
  );
};

export default memo(Index);

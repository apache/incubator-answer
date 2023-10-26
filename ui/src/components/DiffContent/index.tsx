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

import { FC, memo } from 'react';

import classnames from 'classnames';

import { Tag } from '@/components';
import { diffText } from '@/utils';

interface Props {
  objectType: string | 'question' | 'answer' | 'tag';
  newData: Record<string, any>;
  oldData?: Record<string, any>;
  className?: string;
  opts?: Partial<{
    showTitle: boolean;
    showTagUrlSlug: boolean;
  }>;
}

const Index: FC<Props> = ({
  objectType,
  newData,
  oldData,
  className = '',
  opts = {
    showTitle: true,
    showTagUrlSlug: true,
  },
}) => {
  if (!newData) return null;

  let tag = newData.tags;
  if (objectType === 'question' && oldData?.tags) {
    const addTags = newData.tags.filter(
      (c) => !oldData?.tags?.find((p) => p.slug_name === c.slug_name),
    );

    let deleteTags = oldData?.tags
      .filter((c) => !newData?.tags.find((p) => p.slug_name === c.slug_name))
      .map((v) => ({ ...v, state: 'delete' }));

    deleteTags = deleteTags?.map((v) => {
      const index = oldData?.tags?.findIndex(
        (c) => c.slug_name === v.slug_name,
      );
      return {
        ...v,
        pre_index: index,
      };
    });

    tag = newData.tags.map((item) => {
      const find = addTags.find((c) => c.slug_name === item.slug_name);
      if (find) {
        return {
          ...find,
          state: 'add',
        };
      }
      return item;
    });

    deleteTags.forEach((v) => {
      tag.splice(v.pre_index, 0, v);
    });
  }

  return (
    <div className={className}>
      {objectType !== 'answer' && opts?.showTitle && (
        <h5
          dangerouslySetInnerHTML={{
            __html: diffText(
              newData.title?.replace(/</gi, '&lt;'),
              oldData?.title?.replace(/</gi, '&lt;'),
            ),
          }}
          className="mb-3"
        />
      )}
      {objectType === 'question' && (
        <div className="mb-4">
          {tag?.map((item) => {
            return (
              <Tag
                key={item.slug_name}
                className="me-1"
                data={item}
                textClassName={`d-inline-block review-text-${item.state}`}
              />
            );
          })}
        </div>
      )}
      {objectType === 'tag' && opts?.showTagUrlSlug && (
        <div
          className={classnames(
            'small font-monospace',
            newData.original_text && 'mb-4',
          )}
          dangerouslySetInnerHTML={{
            __html: `/tags/${
              newData?.main_tag_slug_name
                ? diffText(
                    newData.main_tag_slug_name,
                    oldData?.main_tag_slug_name,
                  )
                : diffText(newData.slug_name, oldData?.slug_name)
            }`,
          }}
        />
      )}
      <div
        dangerouslySetInnerHTML={{
          __html: diffText(newData.original_text, oldData?.original_text),
        }}
        className="pre-line text-break font-monospace small"
      />
    </div>
  );
};

export default memo(Index);

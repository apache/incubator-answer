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

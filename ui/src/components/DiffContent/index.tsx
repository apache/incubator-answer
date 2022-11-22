import { FC, memo } from 'react';

import { Tag } from '@/components';
import { diffText } from '@/utils';

interface Props {
  currentData: Record<string, any>;
  prevData?: Record<string, any>;
  className?: string;
}

const Index: FC<Props> = ({ currentData, prevData, className = '' }) => {
  if (!currentData?.content) return null;

  let tag;
  if (prevData?.tags) {
    const addTags = currentData.tags.filter(
      (c) => !prevData?.tags?.find((p) => p.slug_name === c.slug_name),
    );

    let deleteTags = prevData?.tags
      .filter(
        (c) => !currentData?.tags.find((p) => p.slug_name === c.slug_name),
      )
      .map((v) => ({ ...v, state: 'delete' }));

    deleteTags = deleteTags?.map((v) => {
      const index = prevData?.tags?.findIndex(
        (c) => c.slug_name === v.slug_name,
      );
      return {
        ...v,
        pre_index: index,
      };
    });

    tag = currentData.tags.map((item) => {
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
      <h5
        dangerouslySetInnerHTML={{
          __html: diffText(currentData.title, prevData?.title),
        }}
        className="mb-3"
      />
      <div className="mb-4">
        {tag.map((item) => {
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
      <div
        dangerouslySetInnerHTML={{
          __html: diffText(currentData.content, prevData?.content),
        }}
        className="pre-line"
      />
    </div>
  );
};

export default memo(Index);

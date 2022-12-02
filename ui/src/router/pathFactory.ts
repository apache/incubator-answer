import urlcat from 'urlcat';

import type * as Type from '@/common/interface';
import Pattern from '@/common/pattern';
import { siteInfoStore } from '@/stores';

const tagLanding = (tag: Type.Tag) => {
  let slugName = tag.main_tag_slug_name || tag.slug_name || '';
  slugName = slugName.toLowerCase();
  return urlcat('/tags/:slugName', { slugName });
};
const tagInfo = (slugName: string) => {
  slugName = slugName.toLowerCase();
  return urlcat('/tags/:slugName/info', { slugName });
};
const tagEdit = (tagId: string) => {
  return urlcat('/tags/:tagId/edit', { tagId });
};
const questionLanding = (questionId: string, title: string = '') => {
  const { siteInfo } = siteInfoStore.getState();
  if (siteInfo.permalink) {
    title = title.toLowerCase();
    title = title.trim().replace(/\s+/g, '-');
    title = title.replace(Pattern.emoji, '');
    if (title) {
      return urlcat('/questions/:questionId/:title', { questionId, title });
    }
  }

  return urlcat('/questions/:questionId', { questionId });
};
const answerLanding = (
  questionId: string,
  questionTitle: string = '',
  answerId: string,
) => {
  const questionLandingUrl = questionLanding(questionId, questionTitle);
  return urlcat(`${questionLandingUrl}/:answerId`, {
    answerId,
  });
};

export const pathFactory = {
  tagLanding,
  tagInfo,
  tagEdit,
  questionLanding,
  answerLanding,
};

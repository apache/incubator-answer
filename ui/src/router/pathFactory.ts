import urlcat from 'urlcat';

import Pattern from '@/common/pattern';
import { seoSettingStore } from '@/stores';

const tagLanding = (slugName: string) => {
  if (!slugName) {
    return '/tags';
  }
  slugName = slugName.toLowerCase();
  return urlcat('/tags/:slugName', { slugName });
};
const tagInfo = (slugName: string) => {
  if (!slugName) {
    return '/tags';
  }
  slugName = slugName.toLowerCase();
  return urlcat('/tags/:slugName/info', { slugName });
};
const tagEdit = (tagId: string) => {
  return urlcat('/tags/:tagId/edit', { tagId });
};
const questionLanding = (questionId: string, title: string = '') => {
  const { seo } = seoSettingStore.getState();
  if (seo.permalink === 1) {
    title = title.toLowerCase();
    title = title.trim().replace(/\s+/g, '-');
    title = title.replace(Pattern.emoji, '');
    if (title) {
      return urlcat('/questions/:questionId/:slugPermalink', {
        questionId,
        slugPermalink: title,
      });
    }
  }

  return urlcat('/questions/:questionId', { questionId });
};
const answerLanding = (params: {
  questionId: string;
  questionTitle?: string;
  answerId: string;
}) => {
  const questionLandingUrl = questionLanding(
    params.questionId,
    params.questionTitle,
  );
  return urlcat(`${questionLandingUrl}/:answerId`, {
    answerId: params.answerId,
  });
};

export const pathFactory = {
  tagLanding,
  tagInfo,
  tagEdit,
  questionLanding,
  answerLanding,
};

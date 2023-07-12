import { seoSettingStore } from '@/stores';

const tagLanding = (slugName: string) => {
  const r = slugName ? `/tags/${slugName}` : '/tags';
  return r;
};

const tagInfo = (slugName: string) => {
  const r = slugName ? `/tags/${slugName}/info` : '/tags';
  return r;
};

const tagEdit = (tagId: string) => {
  const r = `/tags/${tagId}/edit`;
  return r;
};

const questionLanding = (questionId: string, slugTitle: string = '') => {
  const { seo } = seoSettingStore.getState();
  if (!questionId) {
    return slugTitle ? `/questions/null/${slugTitle}` : '/questions/null';
  }
  // @ts-ignore
  if (/[13]/.test(seo.permalink) && slugTitle) {
    return `/questions/${questionId}/${slugTitle}`;
  }

  return `/questions/${questionId}`;
};

const answerLanding = (params: {
  questionId: string;
  slugTitle?: string;
  answerId: string;
}) => {
  const questionLandingUrl = questionLanding(
    params.questionId,
    params.slugTitle,
  );
  return `${questionLandingUrl}/${params.answerId}`;
};

export const pathFactory = {
  tagLanding,
  tagInfo,
  tagEdit,
  questionLanding,
  answerLanding,
};

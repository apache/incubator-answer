import type * as Type from '@/common/interface';

const tagLanding = (tag: Type.Tag) => {
  let slugName = tag.slug_name || '';
  slugName = slugName.toLowerCase();
  return `/tags/${encodeURIComponent(slugName)}`;
};
const tagInfo = (slugName: string) => {
  slugName = slugName.toLowerCase();
  return `/tags/${encodeURIComponent(slugName)}/info`;
};
const tagEdit = (tagId: string) => {
  return `/tags/${tagId}/edit`;
};
const questionLanding = (question_id: string) => {
  return `/questions/${question_id}`;
};
const answerLanding = (question_id: string, answer_id: string) => {
  return `/questions/${question_id}/${answer_id}`;
};

export const pathFactory = {
  tagLanding,
  tagInfo,
  tagEdit,
  questionLanding,
  answerLanding,
};

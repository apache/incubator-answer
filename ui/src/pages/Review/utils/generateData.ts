import { pathFactory } from '@/router/pathFactory';

export default (data: any) => {
  if (!data?.object_id) {
    return {
      itemLink: '',
      itemId: '',
      itemTimePrefix: '',
    };
  }

  const {
    object_type = '',
    object_id = '',
    question_id = '',
    answer_id = '',
    comment_id = '',
    title = '',
  } = data;
  let itemLink = '';
  let itemId = '';
  let itemTimePrefix = '';

  if (object_type === 'question') {
    itemLink = pathFactory.questionLanding(String(object_id), title);
    itemId = String(question_id);
    itemTimePrefix = 'asked';
  } else if (object_type === 'answer') {
    itemLink = pathFactory.answerLanding({
      // @ts-ignore
      questionId: question_id,
      slugTitle: title,
      answerId: String(object_id),
    });
    itemId = String(object_id);
    itemTimePrefix = 'answered';
  } else if (object_type === 'comment') {
    if (question_id && answer_id) {
      itemLink = `${pathFactory.answerLanding({
        questionId: question_id,
        answerId: answer_id,
      })}?commentId=${comment_id}`;
    } else {
      itemLink = `${pathFactory.questionLanding(
        String(question_id),
        title,
      )}?commentId=${comment_id}`;
    }
    itemId = String(comment_id);
    itemTimePrefix = 'commented';
  }

  return {
    itemLink,
    itemId,
    itemTimePrefix,
  };
};

import { debounce } from 'lodash';

import {
  DRAFT_QUESTION_STORAGE_KEY,
  DRAFT_ANSWER_STORAGE_KEY,
} from '@/common/constants';
import { storageExpires as storage } from '@/utils';

export type QuestionDraft = {
  params: {
    title: string;
    content: string;
    tags: any[];
    answer_content: string;
  };
  callback?: () => void;
};

export type AnswerDraft = {
  questionId: string;
  content: string;
  callback?: () => void;
};

type DraftType = {
  type: 'question' | 'answer';
};

export type DraftParams = QuestionDraft | AnswerDraft;

class SaveDraft {
  type: DraftType['type'];

  status: 'save' | 'remove';

  constructor({ type = 'question' }: DraftType) {
    this.type = type;
    this.status = 'save';
  }

  save = debounce((data: DraftParams) => {
    // TODO
    if (this.status === 'remove') {
      return;
    }
    if (this.type === 'question') {
      const { params, callback } = data as QuestionDraft;

      this.storeDraft(params, callback);
    }

    if (this.type === 'answer') {
      const { content, questionId, callback } = data as AnswerDraft;
      if (!questionId || !content) {
        return;
      }

      this.storeDraft({ content, questionId }, callback);
    }
  }, 3000);

  remove() {
    this.status = 'remove';
    const that = this;
    if (this.type === 'question') {
      storage.remove(DRAFT_QUESTION_STORAGE_KEY, () => {
        that.status = 'save';
      });
    }
    if (this.type === 'answer') {
      storage.remove(DRAFT_ANSWER_STORAGE_KEY, () => {
        that.status = 'save';
      });
    }
  }

  private storeDraft = (params: any, callback) => {
    const key =
      this.type === 'question'
        ? DRAFT_QUESTION_STORAGE_KEY
        : DRAFT_ANSWER_STORAGE_KEY;
    storage.set(key, params);
    callback?.();
  };
}

export default SaveDraft;

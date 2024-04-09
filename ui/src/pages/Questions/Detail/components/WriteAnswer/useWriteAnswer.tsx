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

import { useState, useEffect } from 'react';
import { useTranslation } from 'react-i18next';

import { marked } from 'marked';

import { usePromptWithUnload, useCaptchaModal } from '@/hooks';
import { Modal } from '@/components';
import {
  AdminSettingsWrite,
  FormDataType,
  PostAnswerReq,
} from '@/common/interface';
import { postAnswer } from '@/services';
import { guard, handleFormError, SaveDraft, storageExpires } from '@/utils';
import { DRAFT_ANSWER_STORAGE_KEY } from '@/common/constants';
import { writeSettingStore } from '@/stores';

interface Props {
  visible?: boolean;
  data: {
    /** question  id */
    qid: string;
    answered?: boolean;
    loggedUserRank: number;
    first_answer_id?: string;
  };
  callback?: (obj) => void;
}

interface WrtieAnswerData {
  showEditor: boolean;
  formData: FormDataType;
  hasDraft: boolean;
  editorFocusState: boolean;
  focusType: string;
  showTips: boolean;
  writeInfo: AdminSettingsWrite;
}

interface WriteAnswerMethods {
  checkValidated: () => boolean;
  resetForm: () => void;
  deleteDraft: () => void;
  handleSubmit: () => void;
  clickBtn: () => void;
  handleFocusForTextArea: (evt) => void;
  translate: (key: string, options?: Record<string, unknown>) => string;
  setFormData: (data: FormDataType) => void;
  removeDraft: () => void;
  setShowEditor: (val: boolean) => void;
  setShowTips: (val: boolean) => void;
  setFocusType: (val: string) => void;
}

interface WriteAnswerReturn {
  data: WrtieAnswerData;
  methods: WriteAnswerMethods;
}

export const useWriteAnswer = ({
  visible = false,
  data,
  callback,
}: Props): WriteAnswerReturn => {
  const saveDraft = new SaveDraft({ type: 'answer' });
  const { t: translate } = useTranslation('translation', {
    keyPrefix: 'question_detail.write_answer',
  });
  const [formData, setFormData] = useState<FormDataType>({
    content: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });
  const [showEditor, setShowEditor] = useState<boolean>(visible);
  const [focusType, setFocusType] = useState('');
  const [editorFocusState, setEditorFocusState] = useState(false);
  const [hasDraft, setHasDraft] = useState(false);
  const [showTips, setShowTips] = useState(data.loggedUserRank < 100);
  const aCaptcha = useCaptchaModal('answer');
  const writeInfo = writeSettingStore((state) => state.write);

  usePromptWithUnload({
    when: Boolean(formData.content.value),
  });

  const removeDraft = () => {
    // immediately remove debounced save
    saveDraft.save.cancel();
    saveDraft.remove();
    setHasDraft(false);
  };

  useEffect(() => {
    const draft = storageExpires.get(DRAFT_ANSWER_STORAGE_KEY);
    if (draft?.questionId === data.qid && draft?.content) {
      setFormData({
        content: {
          value: draft.content,
          isInvalid: false,
          errorMsg: '',
        },
      });
      setShowEditor(true);
      setHasDraft(true);
    }
  }, []);

  useEffect(() => {
    const draft = storageExpires.get(DRAFT_ANSWER_STORAGE_KEY);
    const { content } = formData;

    if (content.value) {
      // save Draft
      saveDraft.save({
        questionId: data?.qid,
        content: content.value,
      });

      setHasDraft(true);
    } else if (draft?.questionId === data.qid && !content.value) {
      removeDraft();
    }
  }, [formData.content.value]);

  const checkValidated = (): boolean => {
    let bol = true;
    const { content } = formData;

    if (!content.value || Array.from(content.value.trim()).length < 6) {
      bol = false;
      formData.content = {
        value: content.value,
        isInvalid: true,
        errorMsg: translate('characters'),
      };
    } else {
      formData.content = {
        value: content.value,
        isInvalid: false,
        errorMsg: '',
      };
    }

    setFormData({
      ...formData,
    });
    return bol;
  };

  const resetForm = () => {
    setFormData({
      content: {
        value: '',
        isInvalid: false,
        errorMsg: '',
      },
    });
  };

  const deleteDraft = () => {
    const res = window.confirm(
      translate('discard_confirm', { keyPrefix: 'draft' }),
    );
    if (res) {
      removeDraft();
      resetForm();
    }
  };

  const handleSubmit = () => {
    if (!guard.tryNormalLogged(true)) {
      return;
    }
    if (!checkValidated()) {
      return;
    }

    aCaptcha.check(() => {
      const params: PostAnswerReq = {
        question_id: data?.qid,
        content: formData.content.value,
        html: marked.parse(formData.content.value),
      };
      const imgCode = aCaptcha.getCaptcha();
      if (imgCode.verify) {
        params.captcha_code = imgCode.captcha_code;
        params.captcha_id = imgCode.captcha_id;
      }
      postAnswer(params)
        .then(async (res) => {
          await aCaptcha.close();
          setShowEditor(false);
          setFormData({
            content: {
              value: '',
              isInvalid: false,
              errorMsg: '',
            },
          });
          removeDraft();
          callback?.(res.info);
        })
        .catch((ex) => {
          if (ex.isError) {
            aCaptcha.handleCaptchaError(ex.list);
            const stateData = handleFormError(ex, formData);
            setFormData({ ...stateData });
          }
        });
    });
  };

  const clickBtn = () => {
    if (!guard.tryNormalLogged(true)) {
      return;
    }

    if (data?.answered && !showEditor) {
      Modal.confirm({
        title: translate('confirm_title'),
        content: translate('confirm_info'),
        confirmText: translate('continue'),
        onConfirm: () => {
          setShowEditor(true);
        },
      });
      return;
    }

    if (!showEditor) {
      setShowEditor(true);
      return;
    }

    handleSubmit();
  };
  const handleFocusForTextArea = (evt) => {
    if (!guard.tryNormalLogged(true)) {
      evt.currentTarget.blur();
      return;
    }
    setFocusType('answer');
    setShowEditor(true);
    setEditorFocusState(true);
  };

  return {
    data: {
      showEditor,
      formData,
      hasDraft,
      editorFocusState,
      focusType,
      showTips,
      writeInfo,
    },
    methods: {
      checkValidated,
      resetForm,
      deleteDraft,
      handleSubmit,
      clickBtn,
      handleFocusForTextArea,
      translate,
      setFormData,
      removeDraft,
      setShowEditor,
      setShowTips,
      setFocusType,
    },
  };
};

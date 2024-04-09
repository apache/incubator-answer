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

import { memo, FC } from 'react';
import { Form, Button, Alert } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { Trans } from 'react-i18next';

import classNames from 'classnames';

import { Editor, TextArea } from '@/components';
import { useWriteAnswer } from '@/pages/Questions/Detail/components/WriteAnswer/useWriteAnswer';

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

const Index: FC<Props> = ({ visible = false, data: inputData, callback }) => {
  const { data: writeAnswerData, methods: writeAnswerMethods } = useWriteAnswer(
    {
      visible,
      data: inputData,
      callback,
    },
  );

  return (
    <Form noValidate className="mt-4">
      {(!inputData.answered || writeAnswerData.showEditor) && (
        <Form.Group className="mb-3">
          <Form.Label>
            <h5>{writeAnswerMethods.translate('title')}</h5>
          </Form.Label>
          <Form.Control
            isInvalid={writeAnswerData.formData.content.isInvalid}
            className="d-none"
          />
          {!writeAnswerData.showEditor && !inputData.answered && (
            <div className="d-flex">
              <TextArea
                className="w-100"
                rows={8}
                autoFocus={false}
                onFocus={writeAnswerMethods.handleFocusForTextArea}
              />
            </div>
          )}
          {writeAnswerData.showEditor && (
            <>
              <Editor
                className={classNames(
                  'form-control p-0',
                  writeAnswerData.focusType === 'answer' && 'focus',
                )}
                value={writeAnswerData.formData.content.value}
                autoFocus={writeAnswerData.editorFocusState}
                onChange={(val) => {
                  writeAnswerMethods.setFormData({
                    content: {
                      value: val,
                      isInvalid: false,
                      errorMsg: '',
                    },
                  });
                }}
                onFocus={() => {
                  writeAnswerMethods.setFocusType('answer');
                }}
                onBlur={() => {
                  writeAnswerMethods.setFocusType('');
                }}
              />

              <Alert
                variant="warning"
                show={
                  inputData.loggedUserRank < 100 && writeAnswerData.showTips
                }
                onClose={() => writeAnswerMethods.setShowTips(false)}
                dismissible
                className="mt-3">
                <p>{writeAnswerMethods.translate('tips.header_1')}</p>
                <ul>
                  <li>
                    <Trans
                      i18nKey="question_detail.write_answer.tips.li1_1"
                      components={{ strong: <strong /> }}
                    />
                  </li>
                  <li>{writeAnswerMethods.translate('tips.li1_2')}</li>
                </ul>
                <p>
                  <Trans
                    i18nKey="question_detail.write_answer.tips.header_2"
                    components={{ strong: <strong /> }}
                  />
                </p>
                <ul className="mb-0">
                  <li>{writeAnswerMethods.translate('tips.li2_1')}</li>
                </ul>
              </Alert>
            </>
          )}

          <Form.Control.Feedback type="invalid">
            {writeAnswerData.formData.content.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
      )}

      {inputData.answered && !writeAnswerData.showEditor ? (
        // the 0th answer is the oldest one
        <Link
          to={`/posts/${inputData.qid}/${inputData.first_answer_id}/edit`}
          className="btn btn-primary">
          {writeAnswerMethods.translate('edit_answer')}
        </Link>
      ) : (
        <Button onClick={writeAnswerMethods.clickBtn}>
          {writeAnswerMethods.translate('btn_name')}
        </Button>
      )}

      {inputData.answered &&
        !writeAnswerData.showEditor &&
        !writeAnswerData.writeInfo.restrict_answer && (
          <Button
            onClick={writeAnswerMethods.clickBtn}
            className="ms-2 "
            variant="outline-primary">
            {writeAnswerMethods.translate('add_another_answer')}
          </Button>
        )}

      {writeAnswerData.hasDraft && (
        <Button
          variant="link"
          className="ms-2"
          onClick={writeAnswerMethods.deleteDraft}>
          {writeAnswerMethods.translate('discard_draft', { keyPrefix: 'btns' })}
        </Button>
      )}
    </Form>
  );
};

export default memo(Index);

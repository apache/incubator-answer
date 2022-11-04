import qs from 'qs';
import useSWR from 'swr';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const uploadImage = (file) => {
  const form = new FormData();

  form.append('file', file);
  return request.post('/answer/api/v1/user/post/file', form);
};
export const useQueryQuestionByTitle = (title) => {
  return useSWR<Record<string, any>>(
    title ? `/answer/api/v1/question/similar?title=${title}` : '',
    request.instance.get,
  );
};

export const useQueryTags = (params) => {
  return useSWR<Type.ListResult>(
    `/answer/api/v1/tags/page?${qs.stringify(params, {
      skipNulls: true,
    })}`,
    request.instance.get,
  );
};

export const useQueryRevisions = (object_id: string | undefined) => {
  return useSWR<Record<string, any>>(
    object_id ? `/answer/api/v1/revisions?object_id=${object_id}` : '',
    request.instance.get,
  );
};

export const useQueryComments = (params) => {
  if (params.page === 0) {
    params.query_cond = 'vote';
    params.page = 1;
  }
  return useSWR<Type.ListResult>(
    `/answer/api/v1/comment/page?${qs.stringify(params, {
      skipNulls: true,
    })}`,
    request.instance.get,
  );
};

export const updateComment = (params) => {
  return request.put('/answer/api/v1/comment', params);
};

export const deleteComment = (id) => {
  return request.delete('/answer/api/v1/comment', {
    comment_id: id,
  });
};

export const addComment = (params) => {
  return request.post('/answer/api/v1/comment', params);
};

export const queryTags = (tag: string) => {
  return request.get(`/answer/api/v1/question/tags?tag=${tag}`);
};

export const useQueryAnswerInfo = (id: string) => {
  return useSWR<{
    info;
    question;
  }>(`/answer/api/v1/answer/info?id=${id}`, request.instance.get);
};

export const modifyQuestion = (
  params: Type.QuestionParams & { id: string },
) => {
  return request.put(`/answer/api/v1/question`, params);
};

export const modifyAnswer = (params: Type.AnswerParams) => {
  return request.put(`/answer/api/v1/answer`, params);
};

export const login = (params: Type.LoginReqParams) => {
  return request.post<Type.UserInfoRes>(
    '/answer/api/v1/user/login/email',
    params,
  );
};

export const register = (params: Type.RegisterReqParams) => {
  return request.post<any>('/answer/api/v1/user/register/email', params);
};

export const logout = () => {
  return request.get('/answer/api/v1/user/logout');
};

export const verifyEmail = (code: string) => {
  return request.get(`/answer/api/v1/email/verify?code=${code}`);
};

export const resendEmail = (params?: Type.ImgCodeReq) => {
  params = qs.parse(
    qs.stringify(params, {
      skipNulls: true,
    }),
  );
  return request.post('/answer/api/v1/user/email/verification/send', {
    ...params,
  });
};

/**
 * @description get login userinfo
 * @returns {UserInfo}
 */
export const getLoggedUserInfo = () => {
  return request.get<Type.UserInfoRes>('/answer/api/v1/user/info');
};

export const modifyPassword = (params: Type.ModifyPasswordReq) => {
  return request.put('/answer/api/v1/user/password', params);
};

export const modifyUserInfo = (params: Type.ModifyUserReq) => {
  return request.put('/answer/api/v1/user/info', params);
};

export const uploadAvatar = (params: Type.AvatarUploadReq) => {
  return request.post('/answer/api/v1/user/avatar/upload', params);
};

export const resetPassword = (params: Type.PasswordResetReq) => {
  return request.post('/answer/api/v1/user/password/reset', params);
};

export const replacementPassword = (params: { code: string; pass: string }) => {
  return request.post('/answer/api/v1/user/password/replacement', params);
};

export const activateAccount = (code: string) => {
  return request.post(`/answer/api/v1/user/email/verification`, { code });
};

export const checkImgCode = (params: Type.CheckImgReq) => {
  return request.get<Type.ImgCodeRes>(
    `/answer/api/v1/user/action/record?${qs.stringify(params)}`,
  );
};

export const setNotice = (params: Type.SetNoticeReq) => {
  return request.post('/answer/api/v1/user/notice/set', params);
};

export const saveQuestion = (params: Type.QuestionParams) => {
  return request.post('/answer/api/v1/question', params);
};

export const questionDetail = (id: string) => {
  return request.get<Type.QuestionDetailRes>(
    `/answer/api/v1/question/info?id=${id}`,
  );
};

export const getAnswers = (params: Type.AnswersReq) => {
  const apiUrl = `/answer/api/v1/answer/page?${qs.stringify(params)}`;
  return request.get<Type.ListResult<Type.AnswerItem>>(apiUrl);
};

export const postAnswer = (params: Type.PostAnswerReq) => {
  return request.post('/answer/api/v1/answer', params);
};

export const bookmark = (params: { group_id: string; object_id: string }) => {
  return request.post('/answer/api/v1/collection/switch', params);
};

export const postVote = (
  params: { object_id: string; is_cancel: boolean },
  type: 'down' | 'up',
) => {
  return request.post(`/answer/api/v1/vote/${type}`, params);
};

export const following = (params: {
  object_id: string;
  is_cancel: boolean;
}) => {
  return request.post<{ follows: number; is_followed: boolean }>(
    '/answer/api/v1/follow',
    params,
  );
};

export const acceptanceAnswer = (params: {
  answer_id?: string;
  question_id: string;
}) => {
  return request.post('/answer/api/v1/answer/acceptance', params);
};

export const reportList = ({
  type,
  action,
  isBackend = false,
}: Type.ReportParams & { isBackend }) => {
  let api = '/answer/api/v1/reasons';
  if (isBackend) {
    api = '/answer/admin/api/reasons';
  }
  return request.get(`${api}?object_type=${type}&action=${action}`);
};

export const postReport = (params: {
  source: Type.ReportType;
  content: string;
  object_id: string;
  report_type: number;
}) => {
  return request.post('/answer/api/v1/report', params);
};

export const deleteQuestion = (params: { id: string }) => {
  return request.delete('/answer/api/v1/question', params);
};

export const deleteAnswer = (params: { id: string }) => {
  return request.delete('/answer/api/v1/answer', params);
};

export const closeQuestion = (params: {
  id: string;
  close_msg?: string;
  close_type: number;
}) => {
  return request.put('/answer/api/v1/question/status', params);
};

export const changeEmail = (params: { e_mail: string }) => {
  return request.post('/answer/api/v1/user/email/change/code', params);
};

export const changeEmailVerify = (params: { code: string }) => {
  return request.put('/answer/api/v1/user/email', params);
};

export const getAppSettings = () => {
  return request.get<Type.SiteSettings>('/answer/api/v1/siteinfo');
};

export const upgradSystem = () => {
  return request.post('/answer/api/v1/upgradation');
};

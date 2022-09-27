import qs from 'qs';
import useSWR from 'swr';

import request from '@answer/utils/request';

import type * as Type from './types';

export const uploadImage = (file) => {
  const form = new FormData();

  form.append('file', file);
  return request.post('/answer/api/v1/user/post/file', form);
};
export const useQueryQuestionByTitle = (title) => {
  return useSWR<Record<string, any>>(
    title ? `/answer/api/v1/question/title/like?title=${title}` : '',
    request.instance.get,
  );
};

export const useQueryTags = (params) => {
  return useSWR<Type.RecordResult>(
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
  return useSWR<Type.RecordResult>(
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
  return request.get(`/answer/api/v1/question/tag/search?tag=${tag}`);
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
  return request.put(`/answer/api/v1/question/modify`, params);
};

export const modifyAnswer = (params: Type.AnswerParams) => {
  return request.post(`/answer/api/v1/answer/modify`, params);
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

export const emailVerify = (code: string) => {
  return request.get(`/answer/api/v1/email/verify?code=${code}`);
};

export const emailReSend = (params?: Type.ImgCodeReq) => {
  return request.get(
    `/answer/api/v1/user/email/verify/send?${qs.stringify(params, {
      skipNulls: true,
    })}`,
  );
};

/**
 * @description get login userinfo
 * @returns {UserInfo}
 */
export const getUserInfo = () => {
  return request.get<Type.UserInfoRes>('/answer/api/v1/user/info');
};

export const modifyPassword = (params: Type.ModifyPassReq) => {
  return request.post('/answer/api/v1/user/password/modify', params);
};

export const modifyUserInfo = (params: Type.ModifyUserReq) => {
  return request.put('/answer/api/v1/user/info', params);
};

export const uploadAvatar = (params: Type.AvatarUploadReq) => {
  return request.post('/answer/api/v1/user/avatar/upload', params);
};

export const passRetrieve = (params: Type.PssRetReq) => {
  return request.post('/answer/api/v1/user/password/retrieve', params);
};

export const passRetrieveSet = (params: { code: string; pass: string }) => {
  return request.post('/answer/api/v1/user/password/retrieve/set', params);
};

export const accountActivate = (code: string) => {
  return request.get(`/answer/api/v1/user/email/verify?code=${code}`);
};

export const checkImgCode = (params: Type.CheckImgReq) => {
  return request.get<Type.ImgCodeRes>(
    `/answer/api/v1/user/action/record?${qs.stringify(params)}`,
  );
};

export const noticeSet = (params: Type.NoticeSetReq) => {
  return request.post('/answer/api/v1/user/notice/set', params);
};

export const saveQuestion = (params: Type.QuestionParams) => {
  return request.post('/answer/api/v1/question/add', params);
};

export const questionDetail = (id: string) => {
  return request.get<Type.QuDetailRes>(`/answer/api/v1/question/info?id=${id}`);
};

export const langConfig = () => {
  return request.get('/answer/api/v1/language/config');
};

export const languages = () => {
  return request.get<Type.LangsType[]>('/answer/api/v1/language/options');
};

export const getAnswers = (params: Type.AnswersReq) => {
  return request.post<Type.AnswerRes>('/answer/api/v1/answer/list', params);
};

export const postAnswer = (params: Type.PostAnswerReq) => {
  return request.post('/answer/api/v1/answer/add', params);
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

export const adoptAnswer = (params: {
  answer_id: string;
  question_id: string;
}) => {
  return request.post('/answer/api/v1/answer/adopted', params);
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

export const questionDelete = (params: { id: string }) => {
  return request.delete('/answer/api/v1/question/remove', params);
};

export const answerDelete = (params: { id: string }) => {
  return request.delete('/answer/api/v1/answer/remove', params);
};

export const closeQuestion = (params: {
  id: string;
  close_msg?: string;
  close_type: number;
}) => {
  return request.post('/answer/api/v1/question/close', params);
};

export const closeReasons = () => {
  return request.get('/answer/api/v1/question/closemsglist');
};

export const changeEmail = (params: { e_mail: string }) => {
  return request.post('/answer/api/v1/user/email/change/code', params);
};

export const changeEmailVerify = (params: { code: string }) => {
  return request.put('/answer/api/v1/user/email/change', params);
};

export const useSiteSettings = () => {
  const apiUrl = `/answer/api/v1/siteinfo`;
  const { data, error } = useSWR<Type.SiteSettings, Error>(
    [apiUrl],
    request.instance.get,
  );

  return {
    data,
    isLoading: !data && !error,
    error,
  };
};

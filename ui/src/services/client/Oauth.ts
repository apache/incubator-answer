import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const oAuthBindEmail = (data: Type.OauthBindEmailReq) => {
  return request.post('/answer/api/v1/connector/binding/email', data);
};

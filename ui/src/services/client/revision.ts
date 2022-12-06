import request from '@/utils/request';
import * as Type from '@/common/interface';

export const editCheck = (id: string) => {
  const apiUrl = `/answer/api/v1/revisions/edit/check?id=${id}`;
  return request.get(apiUrl);
};

export const revisionAudit = (id: string, operation: 'approve' | 'reject') => {
  const apiUrl = `/answer/api/v1/revisions/audit`;
  return request.put(apiUrl, {
    id,
    operation,
  });
};

export const getReviewList = (page: number) => {
  const apiUrl = `/answer/api/v1/revisions/unreviewed?page=${page}`;
  return request.get<Type.ReviewResp>(apiUrl);
};

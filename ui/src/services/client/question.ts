import useSWR from 'swr';
import qs from 'qs';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const useQuestionList = (params: Type.QueryQuestionsReq) => {
  const apiUrl = `/answer/api/v1/question/page?${qs.stringify(params)}`;
  const { data, error } = useSWR<Type.ListResult, Error>(
    [apiUrl],
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
  };
};

export const useHotQuestions = (
  params: Type.QueryQuestionsReq = {
    page: 1,
    page_size: 6,
    order: 'frequent',
  },
) => {
  const apiUrl = `/answer/api/v1/question/page?${qs.stringify(params)}`;
  const { data, error } = useSWR<Type.ListResult, Error>(
    [apiUrl],
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
  };
};

export const useSimilarQuestion = (params: {
  question_id: string;
  page_size: number;
}) => {
  const apiUrl = `/answer/api/v1/question/similar/tag?${qs.stringify(params)}`;

  const { data, error } = useSWR<Type.ListResult, Error>(
    params.question_id ? apiUrl : null,
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
  };
};

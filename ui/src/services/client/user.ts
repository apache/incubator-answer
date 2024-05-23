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

import useSWR from 'swr';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const useQueryContributeUsers = () => {
  const apiUrl = '/answer/api/v1/user/ranking';
  return useSWR<{
    users_with_the_most_reputation: Type.User[];
    users_with_the_most_vote: Type.User[];
    staffs: Type.User[];
  }>(apiUrl, request.instance.get);
};

export const userSearchByName = (name: string) => {
  const apiUrl = '/answer/api/v1/user/info/search';
  return request.get<Type.UserInfoBase[]>(apiUrl, {
    params: {
      username: name,
    },
  });
};

export type UserPermissionKey =
  | 'question.add'
  | 'question.edit'
  | 'question.edit_without_review'
  | 'question.delete'
  | 'question.close'
  | 'question.reopen'
  | 'question.vote_up'
  | 'question.vote_down'
  | 'question.pin'
  | 'question.unpin'
  | 'question.hide'
  | 'question.show'
  | 'answer.add'
  | 'answer.edit'
  | 'answer.edit_without_review'
  | 'answer.delete'
  | 'answer.accept'
  | 'answer.vote_up'
  | 'answer.vote_down'
  | 'answer.invite_someone_to_answer'
  | 'comment.add'
  | 'comment.edit'
  | 'comment.delete'
  | 'comment.vote_up'
  | 'comment.vote_down'
  | 'report.add'
  | 'tag.add'
  | 'tag.edit'
  | 'tag.edit_slug_name'
  | 'tag.edit_without_review'
  | 'tag.delete, tag.synonym'
  | 'link.url_limit'
  | 'vote.detail'
  | 'answer.audit'
  | 'question.audit'
  | 'tag.audit'
  | 'tag.use_reserved_tag';
export const useUserPermission = (
  keys: UserPermissionKey | UserPermissionKey[],
) => {
  const apiUrl = '/answer/api/v1/permission';
  const action = Array.isArray(keys) ? keys.join(',') : keys;

  return useSWR<
    Partial<
      Record<
        UserPermissionKey,
        {
          has_permission: boolean;
          no_permission_tip: string;
        }
      >
    >
  >([apiUrl, { params: { action } }], request.instance.get);
};

export const useSearchUserStaff = (name: string) => {
  const apiUrl = name
    ? `/answer/api/v1/user/staff?username=${name}&page_size=10`
    : null;
  return useSWR<Type.User[]>(apiUrl, request.instance.get);
};

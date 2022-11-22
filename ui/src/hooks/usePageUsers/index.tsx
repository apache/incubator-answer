import { useState } from 'react';

import { uniqBy } from 'lodash';

import * as Types from '@/common/interface';

let globalUsers: Types.PageUser[] = [];
const usePageUsers = () => {
  const [users, setUsers] = useState<Types.PageUser[]>(globalUsers);
  const getUsers = () => {
    return users;
  };
  return {
    getUsers,
    setUsers: (data: Types.PageUser | Types.PageUser[]) => {
      if (data instanceof Array) {
        setUsers(uniqBy([...users, ...data], 'userName'));
        globalUsers = uniqBy([...globalUsers, ...data], 'userName');
      } else {
        setUsers(uniqBy([...users, data], 'userName'));
        globalUsers = uniqBy([...globalUsers, data], 'userName');
      }
    },
  };
};

export default usePageUsers;

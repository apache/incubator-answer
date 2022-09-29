import { useState } from 'react';

import { uniqBy } from 'lodash';

import * as Types from '@answer/common/interface';

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
        setUsers(uniqBy([...users, ...data], 'name'));
        globalUsers = uniqBy([...globalUsers, ...data], 'name');
      } else {
        setUsers(uniqBy([...users, data], 'name'));
        globalUsers = uniqBy([...globalUsers, data], 'name');
      }
    },
  };
};

export default usePageUsers;

import { memo, FC, useState, useEffect } from 'react';
import { Card, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import classNames from 'classnames';

import { Avatar } from '@/components';
import { getInviteUser, putInviteUser } from '@/services';
import type * as Type from '@/common/interface';

import PeopleDropdown from './PeopleDropdown';

interface Props {
  questionId: string;
  readOnly?: boolean;
}
const Index: FC<Props> = ({ questionId, readOnly = false }) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'invite_to_answer',
  });
  const MAX_ASK_NUMBER = 5;
  const [editing, setEditing] = useState(false);
  const [users, setUsers] = useState<Type.UserInfoBase[]>();

  const initInviteUsers = () => {
    if (!questionId) {
      return;
    }
    getInviteUser(questionId)
      .then((resp) => {
        setUsers(resp);
      })
      .catch(() => {
        if (!users) {
          setUsers([]);
        }
      });
  };

  const updateInviteUsers = (user: Type.UserInfoBase) => {
    let userList = [user];
    if (users?.length) {
      userList = [...users, user];
    }
    setUsers(userList);
  };

  const removeInviteUser = (user: Type.UserInfoBase) => {
    const inviteUsers = users!.filter((_) => {
      return _.username !== user.username;
    });
    setUsers(inviteUsers);
  };

  const saveInviteUsers = () => {
    if (!users) {
      return;
    }
    const names = users.map((_) => {
      return _.username;
    });
    putInviteUser(questionId, names)
      .then(() => {
        setEditing(false);
      })
      .catch((ex) => {
        console.log('ex: ', ex);
      });
  };
  useEffect(() => {
    initInviteUsers();
  }, [questionId]);

  const showAddButton = editing && (!users || users.length < MAX_ASK_NUMBER);
  const showInviteDesc = !editing && users?.length === 0;
  const showEditButton = !readOnly && !editing && users?.length;
  const showSaveButton = !readOnly && editing;

  return (
    <Card className="mt-4">
      <Card.Header className="text-nowrap d-flex justify-content-between text-capitalize">
        {t('title')}
        {showSaveButton ? (
          <Button onClick={saveInviteUsers} variant="link" className="p-0">
            {t('save', { keyPrefix: 'btns' })}
          </Button>
        ) : null}
        {showEditButton ? (
          <Button
            onClick={() => setEditing(true)}
            variant="link"
            className="p-0">
            {t('edit', { keyPrefix: 'btns' })}
          </Button>
        ) : null}
      </Card.Header>
      <Card.Body>
        <div
          className={classNames(
            'd-flex align-items-center flex-wrap',
            editing ? 'm-n1' : ' mx-n2 my-n1',
          )}>
          {users?.map((user) => {
            if (editing) {
              return (
                <Button
                  key={user.username}
                  className="m-1 d-inline-flex flex-nowrap"
                  size="sm"
                  variant="outline-secondary">
                  <Avatar avatar={user.avatar} size="20" />
                  <span className="text-nowrap ms-2">{user.display_name}</span>
                  {/* eslint-disable-next-line jsx-a11y/click-events-have-key-events */}
                  <span
                    className="ps-1 pe-1 me-n1"
                    onClick={() => removeInviteUser(user)}>
                    x
                  </span>
                </Button>
              );
            }
            return (
              <Link
                key={user.username}
                to={`/users/${user.username}`}
                className="mx-2 my-1 d-inline-flex flex-nowrap">
                <Avatar avatar={user.avatar} size="24" />
                <span className="text-nowrap ms-2">{user.display_name}</span>
              </Link>
            );
          })}
          {showAddButton ? (
            <PeopleDropdown
              selectedPeople={users}
              onSelect={updateInviteUsers}
            />
          ) : null}
        </div>
        {showInviteDesc ? (
          <>
            <div className="text-muted">{t('desc')}</div>
            {readOnly ? null : (
              <Button
                size="sm"
                variant="outline-primary"
                className="mt-3"
                onClick={() => setEditing(true)}>
                {t('invite')}
              </Button>
            )}
          </>
        ) : null}
      </Card.Body>
    </Card>
  );
};

export default memo(Index);

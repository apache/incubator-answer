import { FC, memo, useEffect, useState } from 'react';
import { Dropdown, Form } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { loggedUserInfoStore } from '@/stores';
import { userSearchByName } from '@/services';
import { Avatar } from '@/components';
import * as Type from '@/common/interface';

interface Props {
  selectedPeople: Type.UserInfoBase[] | undefined;
  onSelect: (people: Type.UserInfoBase) => void;
}

const Index: FC<Props> = ({ selectedPeople = [], onSelect }) => {
  const { user: currentUser } = loggedUserInfoStore();
  const { t } = useTranslation('translation', {
    keyPrefix: 'invite_to_answer',
  });
  const [toggleState, setToggleState] = useState(false);
  const [peopleList, setPeopleList] = useState<Type.UserInfoBase[]>([]);

  const filterAndSetPeople = (source) => {
    const filteredPeople: Type.UserInfoBase[] = [];
    source.forEach((p) => {
      if (currentUser && currentUser.username === p.username) {
        return;
      }
      if (selectedPeople.find((_) => _.username === p.username)) {
        return;
      }
      filteredPeople.push(p);
    });
    setPeopleList(filteredPeople);
  };

  const searchPeople = (evt) => {
    const name = evt.target.value;
    if (!name) {
      return;
    }
    userSearchByName(name).then((resp) => {
      filterAndSetPeople(resp);
    });
  };

  const handleSelect = (idx) => {
    const people = peopleList[idx];
    if (people) {
      onSelect(people);
    }
  };

  useEffect(() => {
    filterAndSetPeople(peopleList);
  }, [selectedPeople]);

  return (
    <Dropdown
      className="d-inline-flex"
      show={toggleState}
      onSelect={handleSelect}
      onToggle={setToggleState}>
      <Dropdown.Toggle
        className="m-1 no-toggle"
        size="sm"
        variant="outline-secondary">
        {t('add')} +
      </Dropdown.Toggle>

      <Dropdown.Menu>
        <Dropdown.Header className="px-2 py-0">
          <Form.Control
            autoFocus
            placeholder={t('search')}
            onChange={searchPeople}
          />
        </Dropdown.Header>
        {peopleList.map((p, idx) => {
          return (
            <Dropdown.Item
              key={p.username}
              eventKey={idx}
              className={idx === 0 ? 'mt-2' : ''}>
              <div className="d-flex align-items-center text-nowrap">
                <Avatar avatar={p.avatar} size="24" />
                <span className="mx-2">{p.display_name}</span>
                <small className="text-secondary">@{p.username}</small>
              </div>
            </Dropdown.Item>
          );
        })}
      </Dropdown.Menu>
    </Dropdown>
  );
};

export default memo(Index);

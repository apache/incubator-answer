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

import { FC, memo } from 'react';
import { ButtonGroup, Button, DropdownButton, Dropdown } from 'react-bootstrap';
import { useSearchParams, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { floppyNavigation } from '@/utils';

interface Props {
  data;
  i18nKeyPrefix: string;
  currentSort: string;
  sortKey?: string;
  className?: string;
  pathname?: string;
  wrapClassName?: string;
}
const MAX_BUTTON_COUNT = 3;
const Index: FC<Props> = ({
  data = [],
  currentSort = '',
  sortKey = 'order',
  i18nKeyPrefix = '',
  className = '',
  pathname = '',
  wrapClassName = '',
}) => {
  const [searchParams, setUrlSearchParams] = useSearchParams();
  const navigate = useNavigate();

  const { t } = useTranslation('translation', {
    keyPrefix: i18nKeyPrefix,
  });

  const handleParams = (order): string => {
    searchParams.delete('page');
    searchParams.set(sortKey, order);
    const searchStr = searchParams.toString();
    return `?${searchStr}`;
  };

  const handleClick = (e, type) => {
    const str = handleParams(type);
    if (floppyNavigation.shouldProcessLinkClick(e)) {
      e.preventDefault();
      if (pathname) {
        navigate(`${pathname}${str}`);
      } else {
        setUrlSearchParams(str);
      }
    }
  };

  const filteredData = data.filter((_, index) => index > MAX_BUTTON_COUNT - 2);
  const currentBtn = filteredData.find((btn) => {
    return (typeof btn === 'string' ? btn : btn.name) === currentSort;
  });
  return (
    <ButtonGroup size="sm" className={wrapClassName}>
      {data.map((btn, index) => {
        const key = typeof btn === 'string' ? btn : btn.sort;
        const name = typeof btn === 'string' ? btn : btn.name;
        return (
          <Button
            as="a"
            key={key}
            variant="outline-secondary"
            active={currentSort === name}
            className={classNames(
              'text-capitalize fit-content',
              data.length > MAX_BUTTON_COUNT &&
                index > MAX_BUTTON_COUNT - 2 &&
                'd-none d-md-block',
              className,
            )}
            style={
              data.length > MAX_BUTTON_COUNT && index === data.length - 1
                ? {
                    borderTopRightRadius: '0.25rem',
                    borderBottomRightRadius: '0.25rem',
                  }
                : {}
            }
            href={
              pathname ? `${pathname}${handleParams(key)}` : handleParams(key)
            }
            onClick={(evt) => handleClick(evt, key)}>
            {t(name)}
          </Button>
        );
      })}
      {data.length > MAX_BUTTON_COUNT && (
        <DropdownButton
          size="sm"
          variant={currentBtn ? 'secondary' : 'outline-secondary'}
          className="d-block d-md-none"
          as={ButtonGroup}
          title={currentBtn ? t(currentSort) : t('more')}>
          {filteredData.map((btn) => {
            const key = typeof btn === 'string' ? btn : btn.sort;
            const name = typeof btn === 'string' ? btn : btn.name;
            return (
              <Dropdown.Item
                as="a"
                key={key}
                active={currentSort === name}
                className={classNames(
                  'text-capitalize',
                  'd-block d-md-none',
                  className,
                )}
                href={
                  pathname
                    ? `${pathname}${handleParams(key)}`
                    : handleParams(key)
                }
                onClick={(evt) => handleClick(evt, key)}>
                {t(name)}
              </Dropdown.Item>
            );
          })}
        </DropdownButton>
      )}
    </ButtonGroup>
  );
};

export default memo(Index);

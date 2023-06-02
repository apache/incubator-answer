/* eslint-disable no-nested-ternary */
import { FC, useState, useEffect } from 'react';
import { Dropdown, FormControl, Button, Form } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { marked } from 'marked';
import classNames from 'classnames';

import { useTagModal, useToast } from '@/hooks';
import type * as Type from '@/common/interface';
import { queryTags, useUserPermission } from '@/services';

import './index.scss';

interface IProps {
  value?: Type.Tag[];
  onChange?: (tags: Type.Tag[]) => void;
  onFocus?: () => void;
  onBlur?: () => void;
  hiddenDescription?: boolean;
  hiddenCreateBtn?: boolean;
  showRequiredTagText?: boolean;
  alwaysShowAddBtn?: boolean;
}

let timer;

const TagSelector: FC<IProps> = ({
  value = [],
  onChange,
  onFocus = () => {},
  onBlur = () => {},
  hiddenDescription = false,
  hiddenCreateBtn = false,
  alwaysShowAddBtn = false,
  showRequiredTagText = false,
}) => {
  const [initialValue, setInitialValue] = useState<Type.Tag[]>([...value]);
  const [currentIndex, setCurrentIndex] = useState<number>(0);
  const [repeatIndex, setRepeatIndex] = useState(-1);
  const [searchValue, setSearchValue] = useState<string>('');
  const [tags, setTags] = useState<Type.Tag[] | null>(null);
  const { t } = useTranslation('translation', { keyPrefix: 'tag_selector' });
  const [visibleMenu, setVisibleMenu] = useState(false);
  const { data: userPermission } = useUserPermission('tag.add');
  const toast = useToast();
  const tagModal = useTagModal({
    onConfirm: (data) => {
      if (!(onChange instanceof Function)) {
        return;
      }
      const findIndex = initialValue.findIndex(
        (item) => item.slug_name.toLowerCase() === data.slug_name.toLowerCase(),
      );
      if (findIndex === -1) {
        onChange([
          ...value,
          {
            ...data,
            parsed_text: marked(data.original_text),
          },
        ]);
      } else {
        setRepeatIndex(findIndex);
        clearTimeout(timer);
        timer = setTimeout(() => {
          setRepeatIndex(-1);
        }, 2000);
      }
    },
  });

  const filterTags = (result) => {
    const tagArray: Type.Tag[] = [];
    result.forEach((item) => {
      const findIndex = value.findIndex((v) => {
        const tagName1 = v.slug_name.toLowerCase();
        const tagName2 =
          typeof item === 'string'
            ? item.toLowerCase()
            : item.slug_name.toLowerCase();

        return tagName1 === tagName2;
      });

      if (findIndex === -1) {
        tagArray.push(typeof item === 'string' ? { slug_name: item } : item);
      }
    });
    return tagArray;
  };

  useEffect(() => {
    setInitialValue(value);
    if (tags) {
      const tagArray: Type.Tag[] = filterTags(tags || []);

      setTags(tagArray);
    }
  }, [value]);

  const fetchTags = (str) => {
    queryTags(str).then((res) => {
      const tagArray: Type.Tag[] = filterTags(res || []);
      setTags(tagArray?.length > 5 ? tagArray.slice(0, 5) : tagArray);
    });
  };

  useEffect(() => {
    fetchTags(searchValue);
  }, [visibleMenu]);

  const resetSearch = () => {
    setCurrentIndex(0);
    setSearchValue('');
    setTags([]);
  };
  const handleClick = (val: Type.Tag) => {
    const findIndex = initialValue.findIndex(
      (item) => item.slug_name.toLowerCase() === val.slug_name.toLowerCase(),
    );
    if (onChange instanceof Function && findIndex === -1) {
      onChange([
        ...value,
        {
          original_text: '',
          parsed_text: '',
          ...val,
        },
      ]);
    } else {
      setRepeatIndex(findIndex);
      clearTimeout(timer);
      timer = setTimeout(() => {
        setRepeatIndex(-1);
      }, 2000);
    }
    resetSearch();
  };

  const handleRemove = (val: Type.Tag) => {
    if (onChange instanceof Function) {
      onChange(
        value.filter((v) => {
          if (v instanceof Object) {
            return v.slug_name.toLowerCase() !== val.slug_name.toLowerCase();
          }
          return v !== val;
        }),
      );
    }
  };

  const handleSearch = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const searchStr = e.currentTarget.value.replace(';', '');
    setSearchValue(searchStr);
    fetchTags(searchStr);
  };

  const handleSelect = (eventKey) => {
    setCurrentIndex(eventKey);
  };
  const handleKeyDown = (e) => {
    e.stopPropagation();

    if (!tags) {
      return;
    }
    const { keyCode } = e;

    if (keyCode === 38 && currentIndex > 0) {
      setCurrentIndex(currentIndex - 1);
    }
    if (keyCode === 40 && currentIndex < tags.length - 1) {
      setCurrentIndex(currentIndex + 1);
    }

    if (keyCode === 13 && currentIndex > -1) {
      e.preventDefault();

      if (tags.length === 0) {
        tagModal.onShow(searchValue);
        return;
      }
      if (currentIndex <= tags.length - 1) {
        handleClick(tags[currentIndex]);
        // if (currentIndex === tags.length - 1 && currentIndex > 0) {
        //   setCurrentIndex(currentIndex - 1);
        // }
      }
    }
  };

  const handleCreate = () => {
    const tagAddPermission = userPermission?.['tag.add'];
    if (!tagAddPermission || tagAddPermission?.has_permission) {
      tagModal.onShow(searchValue);
    } else if (tagAddPermission?.no_permission_tip) {
      toast.onShow({
        msg: tagAddPermission.no_permission_tip,
        variant: 'danger',
      });
    }
  };

  return (
    <div
      className="tag-selector-wrap"
      onFocus={onFocus}
      onBlur={onBlur}
      onKeyDown={handleKeyDown}>
      <div className="d-flex flex-wrap m-n1">
        {initialValue?.map((item, index) => {
          return (
            <Button
              key={item.slug_name}
              className={classNames(
                'm-1 text-nowrap d-flex align-items-center',
                index === repeatIndex && 'bg-fade-out',
              )}
              variant={`outline-${
                item.reserved ? 'danger' : item.recommend ? 'dark' : 'secondary'
              }`}
              size="sm">
              {item.display_name}
              <span className="ms-1" onMouseUp={() => handleRemove(item)}>
                ×
              </span>
            </Button>
          );
        })}
        {initialValue?.length < 5 || alwaysShowAddBtn ? (
          <Dropdown onSelect={handleSelect} onToggle={setVisibleMenu}>
            <Dropdown.Toggle
              className={classNames('m-1')}
              variant="outline-secondary"
              size="sm">
              <span className="me-1">+</span>
              {t('add_btn')}
            </Dropdown.Toggle>
            <Dropdown.Menu>
              {visibleMenu && (
                <Dropdown.Header>
                  <Form
                    onSubmit={(e) => {
                      e.preventDefault();
                    }}>
                    <FormControl
                      placeholder={t('search_tag')}
                      autoFocus
                      value={searchValue}
                      onChange={handleSearch}
                    />
                  </Form>
                </Dropdown.Header>
              )}
              {!searchValue &&
                showRequiredTagText &&
                tags &&
                tags.filter((v) => v.recommend)?.length > 0 && (
                  <h6 className="dropdown-header">{t('tag_required_text')}</h6>
                )}

              {tags?.map((item, index) => {
                return (
                  <Dropdown.Item
                    key={item.slug_name}
                    eventKey={index}
                    active={index === currentIndex}
                    onClick={() => handleClick(item)}>
                    {item.display_name}
                  </Dropdown.Item>
                );
              })}
              {searchValue && tags && tags.length === 0 && (
                <Dropdown.Item disabled className="text-secondary">
                  {t('no_result')}
                </Dropdown.Item>
              )}
              {!hiddenCreateBtn && searchValue && (
                <Button
                  variant="link"
                  className="px-3 btn-no-border w-100 text-start"
                  onClick={handleCreate}>
                  + {t('create_btn')}
                </Button>
              )}
            </Dropdown.Menu>
          </Dropdown>
        ) : null}
      </div>
      {!hiddenDescription && <Form.Text>{t('hint')}</Form.Text>}
    </div>
  );
};

export default TagSelector;

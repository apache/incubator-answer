import { FC, useState, useEffect } from 'react';
import { Dropdown, FormControl, Button, Form } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { marked } from 'marked';
import classNames from 'classnames';

import { useTagModal } from '@answer/hooks';
import type * as Type from '@answer/common/interface';

import { queryTags } from '@/services';

import './index.scss';

interface IProps {
  value?: Type.Tag[];
  onChange?: (tags: Type.Tag[]) => void;
  onFocus?: () => void;
  onBlur?: () => void;
  hiddenDescription?: boolean;
  hiddenCreateBtn?: boolean;
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
}) => {
  const [initialValue, setInitialValue] = useState<Type.Tag[]>([...value]);
  const [currentIndex, setCurrentIndex] = useState<number>(0);
  const [repeatIndex, setRepeatIndex] = useState(-1);
  const [tag, setTag] = useState<string>('');
  const [tags, setTags] = useState<Type.Tag[] | null>(null);
  const { t } = useTranslation('translation', { keyPrefix: 'tag_selector' });
  const [visibleMenu, setVisibleMenu] = useState(false);

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

  useEffect(() => {
    if (!tag) {
      setTags(null);
      return;
    }

    queryTags(tag).then((res) => {
      const tagArray: Type.Tag[] = filterTags(res || []);
      setTags(tagArray);
    });
  }, [tag]);

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
    setTag(e.currentTarget.value.replace(';', ''));
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
    if (
      keyCode === 13 &&
      currentIndex > -1 &&
      currentIndex <= tags.length - 1
    ) {
      e.preventDefault();
      handleClick(tags[currentIndex]);
    }
  };
  return (
    <div
      className="tag-selector-wrap"
      onFocus={onFocus}
      onBlur={onBlur}
      onKeyDown={handleKeyDown}>
      <div className="d-flex flex-wrap mx-n1">
        {initialValue?.map((item, index) => {
          return (
            <Button
              key={item.slug_name}
              className={classNames(
                'm-1 text-nowrap d-flex align-items-center',
                index === repeatIndex && 'warning',
              )}
              variant="outline-secondary"
              size="sm">
              {item.slug_name}
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
                      value={tag}
                      onChange={handleSearch}
                    />
                  </Form>
                </Dropdown.Header>
              )}

              {tags?.map((item, index) => {
                return (
                  <Dropdown.Item
                    key={item.slug_name}
                    eventKey={index}
                    active={index === currentIndex}
                    onClick={() => handleClick(item)}>
                    {item.slug_name}
                  </Dropdown.Item>
                );
              })}
              {tag && tags && tags.length === 0 && (
                <Dropdown.Item disabled className="text-secondary">
                  {t('no_result')}
                </Dropdown.Item>
              )}
              {!hiddenCreateBtn && tag && (
                <Button
                  variant="link"
                  className="px-3 btn-no-border w-100 text-start"
                  onClick={() => {
                    tagModal.onShow();
                  }}>
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

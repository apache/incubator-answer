import { useState } from 'react';
import { Container, Row, Col, Button, Card } from 'react-bootstrap';
import { useParams, useNavigate, Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import {
  Tag,
  TagSelector,
  FormatTime,
  Modal,
  PageTitle,
} from '@answer/components';
import {
  useTagInfo,
  useQuerySynonymsTags,
  saveSynonymsTags,
  deleteTag,
} from '@answer/api';

const TagIntroduction = () => {
  const [isEdit, setEditState] = useState(false);
  const { tagName } = useParams();
  const { data: tagInfo } = useTagInfo({ name: tagName });
  const { t } = useTranslation('translation', { keyPrefix: 'tag_info' });
  const navigate = useNavigate();
  const { data: synonymsTags, mutate } = useQuerySynonymsTags(tagInfo?.tag_id);
  if (!tagInfo) {
    return null;
  }
  if (tagInfo.main_tag_slug_name) {
    navigate(`/tags/${tagInfo.main_tag_slug_name}/info`, { replace: true });
    return null;
  }
  const handleEdit = () => {
    setEditState(true);
  };

  const handleSave = () => {
    saveSynonymsTags({
      tag_id: tagInfo?.tag_id,
      synonym_tag_list: synonymsTags,
    }).then(() => {
      mutate();
      setEditState(false);
    });
  };

  const handleTagsChange = (value) => {
    mutate([...value], {
      revalidate: false,
    });
  };

  const handleEditTag = () => {
    navigate(`/tags/${tagInfo?.tag_id}/edit`);
  };
  const handleDeleteTag = () => {
    if (synonymsTags && synonymsTags.length > 0) {
      Modal.confirm({
        title: t('delete.title'),
        content: t('delete.content2'),
        showConfirm: false,
        cancelText: t('delete.close'),
      });
      return;
    }
    Modal.confirm({
      title: t('delete.title'),
      content: t('delete.content'),
      onConfirm: () => {
        deleteTag(tagInfo.tag_id);
      },
    });
  };
  const onAction = (params) => {
    if (params.action === 'edit') {
      handleEditTag();
    } else if (params.action === 'delete') {
      handleDeleteTag();
    }
  };

  let pageTitle = '';
  if (tagInfo) {
    pageTitle = `'${tagInfo.display_name}' ${t('tag_wiki', {
      keyPrefix: 'page_title',
    })}`;
  }
  return (
    <>
      <PageTitle title={pageTitle} />
      <Container className="pt-4 mt-2 mb-5">
        <Row className="justify-content-center">
          <Col xxl={7} lg={8} sm={12}>
            <h3 className="mb-3">
              <Link
                to={`/tags/${tagInfo?.slug_name}`}
                replace
                className="link-dark">
                {tagInfo.display_name}
              </Link>
            </h3>

            <div className="text-secondary mb-4 fs-14">
              <FormatTime preFix={t('created_at')} time={tagInfo.created_at} />
              <FormatTime
                preFix={t('edited_at')}
                className="ms-3"
                time={tagInfo.updated_at}
              />
            </div>

            <div
              className="content text-break"
              dangerouslySetInnerHTML={{ __html: tagInfo?.parsed_text }}
            />
            <div className="mt-4">
              {tagInfo?.member_actions.map((action, index) => {
                return (
                  <Button
                    key={action.name}
                    variant="link"
                    className={classNames(
                      'link-secondary btn-no-border p-0 fs-14',
                      index > 0 && 'ms-3',
                    )}
                    onClick={() => onAction(action)}>
                    {action.name}
                  </Button>
                );
              })}
            </div>
          </Col>
          <Col xxl={3} lg={4} sm={12} className="mt-5 mt-lg-0">
            <Card>
              <Card.Header className="d-flex justify-content-between">
                <span>{t('synonyms.title')}</span>
                {isEdit ? (
                  <Button
                    variant="link"
                    className="p-0 btn-no-border"
                    onClick={handleSave}>
                    {t('synonyms.btn_save')}
                  </Button>
                ) : (
                  <Button
                    variant="link"
                    className="p-0 btn-no-border"
                    onClick={handleEdit}>
                    {t('synonyms.btn_edit')}
                  </Button>
                )}
              </Card.Header>
              <Card.Body>
                {isEdit && (
                  <>
                    <div className="mb-3">
                      {t('synonyms.text')}{' '}
                      <Tag className="me-2 mb-2" href="#">
                        {tagName}
                      </Tag>
                    </div>
                    <TagSelector
                      value={synonymsTags}
                      onChange={handleTagsChange}
                      hiddenDescription
                    />
                  </>
                )}
                {!isEdit &&
                  (synonymsTags && synonymsTags.length > 0 ? (
                    synonymsTags.map((item) => {
                      return (
                        <Tag
                          key={item.tag_id}
                          className="me-2 mb-2"
                          href={`/tags/${item.slug_name}`}>
                          {item.slug_name}
                        </Tag>
                      );
                    })
                  ) : (
                    <>
                      <div className="text-muted mb-3">
                        {t('synonyms.empty')}
                      </div>
                      <Button
                        variant="outline-primary"
                        size="sm"
                        onClick={handleEdit}>
                        {t('synonyms.btn_add')}
                      </Button>
                    </>
                  ))}
              </Card.Body>
            </Card>
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default TagIntroduction;

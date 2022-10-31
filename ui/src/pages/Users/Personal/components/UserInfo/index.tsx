import { FC, memo } from 'react';
import { Badge, OverlayTrigger, Tooltip } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import { Avatar, Icon } from '@answer/components';
import type { UserInfoRes } from '@answer/common/interface';

interface Props {
  data: UserInfoRes;
}

const Index: FC<Props> = ({ data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });
  if (!data?.username) {
    return null;
  }
  return (
    <div className="d-flex mb-4">
      {data?.status !== 'deleted' ? (
        <Link to={`/users/${data.username}`} reloadDocument>
          <Avatar avatar={data.avatar} size="160px" searchStr="s=128" />
        </Link>
      ) : (
        <Avatar avatar={data.avatar} size="160px" searchStr="s=128" />
      )}

      <div className="ms-4">
        <div className="d-flex align-items-center mb-2">
          {data?.status !== 'deleted' ? (
            <Link
              to={`/users/${data.username}`}
              className="link-dark h3 mb-0"
              reloadDocument>
              {data.display_name}
            </Link>
          ) : (
            <span className="link-dark h3 mb-0">{data.display_name}</span>
          )}
          {data?.is_admin && (
            <div className="ms-2">
              <OverlayTrigger
                placement="top"
                overlay={<Tooltip>{t('mod_long')}</Tooltip>}>
                <Badge bg="light" className="text-body">
                  {t('mod_short')}
                </Badge>
              </OverlayTrigger>
            </div>
          )}
        </div>
        <div className="text-secondary mb-4">@{data.username}</div>

        <div className="d-flex mb-3">
          <div className="me-3">
            <strong className="fs-5">{data.rank || 0}</strong>
            <span className="text-secondary"> {t('x_reputation')}</span>
          </div>
          <div className="me-3">
            <strong className="fs-5">{data.answer_count || 0}</strong>
            <span className="text-secondary"> {t('x_answers')}</span>
          </div>
          <div>
            <strong className="fs-5">{data?.question_count || 0}</strong>
            <span className="text-secondary"> {t('x_questions')}</span>
          </div>
        </div>

        <div className="d-flex text-secondary">
          {data.location && (
            <div className="d-flex align-items-center me-3">
              <Icon name="geo-alt-fill" className="me-2" />
              <span>{data.location}</span>
            </div>
          )}

          {data.website && (
            <div className="d-flex align-items-center">
              <Icon name="house-door-fill" className="me-2" />
              <a
                className="link-secondary"
                href={
                  data.website?.includes('http')
                    ? data.website
                    : `http://${data.website}`
                }>
                {data?.website.replace(/(http|https):\/\//, '').split('/')?.[0]}
              </a>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default memo(Index);

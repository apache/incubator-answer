import { memo, FC, useState } from 'react';
import { useSearchParams, Link } from 'react-router-dom';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { following } from '@/services';
import { tryNormalLogged } from '@/utils/guard';

interface Props {
  data;
}

const reg =
  /(\[.*\])|(is:answer)|(is:question)|(score:\d*)|(user:\S*)|(answers:\d*)/g;
const Index: FC<Props> = ({ data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'search' });
  const [searchParams] = useSearchParams();
  const q = searchParams.get('q');
  const options = q?.match(reg);
  const [followed, setFollowed] = useState(data?.is_follower);

  const follow = () => {
    if (!tryNormalLogged(true)) {
      return;
    }
    following({
      object_id: data?.tag_id,
      is_cancel: followed,
    }).then((res) => {
      setFollowed(res.is_followed);
    });
  };

  return (
    <div className="mb-5">
      <h3 className="mb-3">{t('title')}</h3>
      <p>
        <span className="me-1 text-secondary">{t('keywords')}</span>
        {q?.replace(reg, '')}
        <br />
        {options?.length && (
          <>
            <span className="text-secondary">{t('options')}</span>
            {options?.map((item) => {
              return <code key={item}>{item} </code>;
            })}
          </>
        )}
      </p>
      {data?.slug_name && (
        <>
          {data.excerpt && (
            <p className="text-break">
              {data.excerpt}
              <Link to={`/tags/${data.slug_name}/info`}> [{t('more')}]</Link>
            </p>
          )}

          <Button variant="outline-primary" onClick={follow}>
            {followed ? t('following') : t('follow')}
          </Button>
        </>
      )}
    </div>
  );
};

export default memo(Index);

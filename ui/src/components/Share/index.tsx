import { memo, FC, useState, useEffect } from 'react';
import { Dropdown, OverlayTrigger, Tooltip } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { FacebookShareButton, TwitterShareButton } from 'next-share';
import copy from 'copy-to-clipboard';

import { userInfoStore } from '@answer/stores';

interface IProps {
  type: 'answer' | 'question';
  qid: any;
  aid?: any;
  title: string;
}

const Index: FC<IProps> = ({ type, qid, aid, title }) => {
  const user = userInfoStore((state) => state.user);
  const [show, setShow] = useState(false);
  const [showTip, setShowTip] = useState(false);
  const [canSystemShare, setSystemShareState] = useState(false);
  const { t } = useTranslation();

  let baseUrl =
    type === 'question'
      ? `${window.location.origin}/questions/${qid}`
      : `${window.location.origin}/questions/${qid}/${aid}`;
  if (user.id) {
    baseUrl = `${baseUrl}?shareUserId=${user.username}`;
  }

  const closeShare = () => {
    setShowTip(false);
    setShow(false);
  };

  const handleCopy = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();
    copy(baseUrl);
    setShowTip(true);
    setTimeout(closeShare, 1000);
  };

  const systemShare = () => {
    navigator.share({
      title,
      text: `${title} - Answer：`,
      url: baseUrl,
    });
  };
  useEffect(() => {
    if (window.navigator?.canShare?.({ text: 'can_share' })) {
      setSystemShareState(true);
    }
  }, []);
  return (
    <Dropdown show={show} onToggle={closeShare}>
      <Dropdown.Toggle
        id="dropdown-share"
        as="a"
        className="no-toggle fs-14 link-secondary pointer me-3"
        onClick={() => setShow(true)}
        style={{ lineHeight: '23px' }}>
        {t('share.name')}
      </Dropdown.Toggle>
      <Dropdown.Menu style={{ width: '195px' }}>
        <OverlayTrigger
          trigger="click"
          placement="left"
          show={showTip}
          overlay={<Tooltip>{t('share.copied')}</Tooltip>}>
          <Dropdown.Item onClick={handleCopy} eventKey="copy">
            {t('share.copy')}
          </Dropdown.Item>
        </OverlayTrigger>
        <Dropdown.Item eventKey="facebook">
          <FacebookShareButton
            title={title}
            url={baseUrl}
            className="w-100 py-1 px-3 text-start">
            {t('share.facebook')}
          </FacebookShareButton>
        </Dropdown.Item>
        <Dropdown.Item>
          <TwitterShareButton
            title={title}
            url={baseUrl}
            className="w-100 py-1 px-3 text-start">
            {t('share.twitter')}
          </TwitterShareButton>
        </Dropdown.Item>
        {canSystemShare && (
          <Dropdown.Item onClick={systemShare}>{t('share.via')}</Dropdown.Item>
        )}
      </Dropdown.Menu>
    </Dropdown>
  );
};

export default memo(Index);

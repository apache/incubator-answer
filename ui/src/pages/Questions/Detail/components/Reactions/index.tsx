import { FC, memo } from 'react';
import { Button, OverlayTrigger, Tooltip } from 'react-bootstrap';

import { Icon } from '@/components';

interface Props {}

interface Data {
  title: string;
  icon: string;
  className: string;
}

const data: Data[] = [
  { title: 'hello', icon: 'heart-fill', className: 'text-danger' },
  { title: 'hello', icon: 'emoji-laughing-fill', className: 'text-warning' },
  { title: 'hello', icon: 'emoji-frown-fill', className: 'text-warning' },
];

const Index: FC<Props> = () => {
  const renderTooltip = (props) => (
    <Tooltip id="reaction-button-tooltip" {...props} bsPrefix="tooltip">
      <div className="d-block d-md-flex flex-wrap m-0 p-0">
        {data.map((d) => (
          <Button
            key={d.icon}
            variant="light"
            size="sm"
            onClick={() => alert('hellob')}>
            <Icon name={d.icon} className={d.className} />
          </Button>
        ))}
      </div>
    </Tooltip>
  );

  return (
    <div className="d-block d-md-flex flex-wrap mt-4 mb-3">
      <Button
        variant="outline-secondary"
        className="rounded-pill answer-reaction-btn"
        size="sm">
        <Icon name="chat-text-fill" />
        <span className="ms-1">{6} comments</span>
      </Button>

      <OverlayTrigger trigger="click" placement="top" overlay={renderTooltip}>
        <Button
          variant="outline-secondary"
          size="sm"
          className="rounded-pill ms-2 answer-reaction-btn">
          <Icon name="emoji-smile-fill" />
          <span className="ms-1">+</span>
        </Button>
      </OverlayTrigger>

      {/* <div className="arrow" /> */}

      {data.map((d) => (
        <OverlayTrigger
          placement="top"
          // trigger="hover"
          trigger="click"
          overlay={
            <Tooltip>
              <div className="text-start">
                <b>heart</b> <br /> fenbox, joyqi, robin, andrus, jackathon and
                7 more...
              </div>
            </Tooltip>
          }>
          <Button
            title="hahah"
            variant="outline-secondary"
            className="rounded-pill ms-2 answer-reaction-btn"
            size="sm">
            <Icon name={d.icon} className={d.className} />
            <span className="ms-1">{3}</span>
          </Button>
        </OverlayTrigger>
      ))}
    </div>
  );
};

export default memo(Index);

import { FC } from 'react';

import { base64ToSvg } from '@/utils';

interface IProps {
  base64: string | undefined;
}
const Icon: FC<IProps> = ({ base64 = '' }) => {
  return base64 ? (
    <span
      dangerouslySetInnerHTML={{
        __html: base64ToSvg(base64),
      }}
    />
  ) : null;
};

export default Icon;

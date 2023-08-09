import { FC } from 'react';

import { base64ToSvg } from '@/utils';

interface IProps {
  svgClassName?: string;
  base64: string | undefined;
}
const Icon: FC<IProps> = ({ base64 = '', svgClassName = '' }) => {
  return base64 ? (
    <span
      dangerouslySetInnerHTML={{
        __html: base64ToSvg(base64, svgClassName),
      }}
    />
  ) : null;
};

export default Icon;

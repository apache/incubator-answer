import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import { uploadImage } from '@/services';
import * as Type from '@/common/interface';

interface IProps {
  type: Type.UploadType;
  className?: classnames.Argument;
  children?: React.ReactNode;
  acceptType?: string;
  disabled?: boolean;
  uploadCallback: (img: string) => void;
}

const Index: React.FC<IProps> = ({
  type,
  uploadCallback,
  children,
  acceptType = '',
  className,
  disabled = false,
}) => {
  const { t } = useTranslation();
  const [status, setStatus] = useState(false);

  const onChange = (e: any) => {
    if (status) {
      return;
    }
    if (e.target.files[0]) {
      // const fileSize = e.target.files[0].size || 0;

      // if (maxSize && fileSize / 1024 / 1024 > 2) {
      //   Modal.confirm({
      //     content: '请上传小于 2M 的图片',
      //   });
      //   return;
      // }
      setStatus(true);
      uploadImage({ file: e.target.files[0], type })
        .then((res) => {
          uploadCallback(res);
        })
        .finally(() => {
          setStatus(false);
        });
    }
  };

  return (
    <label
      className={classnames('btn btn-outline-secondary', className, {
        disabled: !!disabled,
      })}>
      {children || (status ? t('upload_img.loading') : t('upload_img.name'))}
      <input
        type="file"
        className="d-none"
        disabled={disabled}
        accept={`image/jpeg,image/jpg,image/png,image/webp${acceptType}`}
        onChange={onChange}
      />
    </label>
  );
};

export default React.memo(Index);

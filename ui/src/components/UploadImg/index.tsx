import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';

interface IProps {
  type: string;
  className?: string;
  children?: React.ReactNode;
  upload: (data: FormData) => Promise<any>;
}

const Index: React.FC<IProps> = ({ type, upload, children, className }) => {
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
      const data = new FormData();

      data.append('file', e.target.files[0]);
      // do
      upload(data).finally(() => {
        setStatus(false);
      });
    }
  };

  return (
    <label className={`btn btn-outline-secondary uploadBtn ${className}`}>
      {children || (status ? t('upload_img.loading') : t('upload_img.name'))}
      <input
        type="file"
        className="d-none"
        accept="image/jpeg,image/jpg,image/png,image/webp"
        onChange={onChange}
        id={type}
      />
    </label>
  );
};

export default React.memo(Index);

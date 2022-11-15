import { FC } from 'react';
import { ButtonGroup, Button } from 'react-bootstrap';

import { Icon, UploadImg } from '@/components';
import { UploadType } from '@/common/interface';

interface Props {
  type: UploadType;
  value: string;
  onChange: (value: string) => void;
}

const Index: FC<Props> = ({ type = 'post', value, onChange }) => {
  const onUpload = (imgPath: string) => {
    onChange(imgPath);
  };

  const onRemove = () => {
    onChange('');
  };
  return (
    <div className="d-flex">
      <div className="bg-gray-300 upload-img-wrap me-2 d-flex align-items-center justify-content-center">
        <img src={value} alt="" height={100} />
      </div>
      <ButtonGroup vertical className="fit-content">
        <UploadImg type={type} uploadCallback={onUpload} className="mb-0">
          <Icon name="cloud-upload" />
        </UploadImg>

        <Button variant="outline-secondary" onClick={onRemove}>
          <Icon name="trash" />
        </Button>
      </ButtonGroup>
    </div>
  );
};

export default Index;

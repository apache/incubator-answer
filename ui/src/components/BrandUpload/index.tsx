import { FC } from 'react';
import { ButtonGroup, Button } from 'react-bootstrap';

import { Icon, UploadImg } from '@/components';

interface Props {
  type: string;
  imgPath: string;
  uploadCallback: (data: FormData) => Promise<any>;
  deleteCallback: (type: string) => void;
}

const Index: FC<Props> = ({
  type,
  imgPath,
  uploadCallback,
  deleteCallback,
}) => {
  return (
    <div className="d-flex">
      <div className="bg-gray-300 upload-img-wrap me-2 d-flex align-items-center justify-content-center">
        <img src={imgPath} alt="" height={100} />
      </div>
      <ButtonGroup vertical className="fit-content">
        <UploadImg type={type} upload={uploadCallback} className="mb-0">
          <Icon name="cloud-upload" />
        </UploadImg>

        <Button
          variant="outline-secondary"
          onClick={() => deleteCallback(type)}>
          <Icon name="trash" />
        </Button>
      </ButtonGroup>
    </div>
  );
};

export default Index;

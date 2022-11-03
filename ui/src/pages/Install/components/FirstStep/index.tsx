import { FC, useEffect, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type { LangsType, FormValue, FormDataType } from '@/common/interface';
import Progress from '../Progress';
import { languages } from '@/services';

interface Props {
  data: FormValue;
  changeCallback: (value: FormDataType) => void;
  nextCallback: () => void;
  visible: boolean;
}
const Index: FC<Props> = ({ visible, data, changeCallback, nextCallback }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'install' });

  const [langs, setLangs] = useState<LangsType[]>();

  const getLangs = async () => {
    const res: LangsType[] = await languages();
    setLangs(res);
  };

  const handleSubmit = () => {
    nextCallback();
  };

  useEffect(() => {
    getLangs();
  }, []);

  if (!visible) return null;
  return (
    <Form noValidate onSubmit={handleSubmit}>
      <Form.Group controlId="lang" className="mb-3">
        <Form.Label>{t('lang.label')}</Form.Label>
        <Form.Select
          value={data.value}
          isInvalid={data.isInvalid}
          onChange={(e) => {
            changeCallback({
              lang: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            });
          }}>
          {langs?.map((item) => {
            return (
              <option value={item.value} key={item.value}>
                {item.label}
              </option>
            );
          })}
        </Form.Select>
      </Form.Group>

      <div className="d-flex align-items-center justify-content-between">
        <Progress step={1} />
        <Button type="submit">{t('next')}</Button>
      </div>
    </Form>
  );
};

export default Index;

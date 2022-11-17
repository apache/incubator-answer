import { Form } from 'react-bootstrap';

import { TIMEZONES } from '@/common/constants';

const TimeZonePicker = (props) => {
  return (
    <Form.Select {...props}>
      {TIMEZONES?.map((item) => {
        return (
          <optgroup label={item.label} key={item.label}>
            {item.options.map((option) => {
              return (
                <option value={option.value} key={option.value}>
                  {option.label}
                </option>
              );
            })}
          </optgroup>
        );
      })}
    </Form.Select>
  );
};

export default TimeZonePicker;

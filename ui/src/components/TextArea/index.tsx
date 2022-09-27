import { FC, useRef, useEffect, memo } from 'react';
import { FormControl, FormControlProps } from 'react-bootstrap';

const TextArea: FC<FormControlProps> = ({ value, onChange, size }) => {
  const ref = useRef<HTMLTextAreaElement>(null);

  const autoGrow = () => {
    if (ref.current) {
      ref.current.style.height = 'auto';
      ref.current.style.height = `${ref.current.scrollHeight}px`;
    }
  };

  useEffect(() => {
    if (ref.current && value) {
      autoGrow();
    }
  }, [ref, value]);

  return (
    <FormControl
      as="textarea"
      className="resize-none font-monospace"
      rows={1}
      size={size}
      value={value}
      onChange={onChange}
      autoFocus
      ref={ref}
      onInput={autoGrow}
    />
  );
};
export default memo(TextArea);

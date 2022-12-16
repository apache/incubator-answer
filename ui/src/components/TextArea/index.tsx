import { FC, useRef, useEffect, memo } from 'react';
import { FormControl, FormControlProps } from 'react-bootstrap';

const TextArea: FC<
  FormControlProps & { rows?: number; autoFocus?: boolean }
> = ({ value, onChange, size, rows = 1, autoFocus = true, ...rest }) => {
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
      rows={rows}
      size={size}
      value={value}
      onChange={onChange}
      autoFocus={autoFocus}
      ref={ref}
      onInput={autoGrow}
      {...rest}
    />
  );
};
export default memo(TextArea);

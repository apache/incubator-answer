import {
  forwardRef,
  useEffect,
  useRef,
  useState,
  memo,
  useImperativeHandle,
} from 'react';

import { marked } from 'marked';

import { htmlRender } from './utils';

let scrollTop = 0;
marked.setOptions({
  breaks: true,
  sanitize: true,
});

const Index = ({ value }, ref) => {
  const [html, setHtml] = useState('');

  const previewRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const previewHtml = marked(value).replace(
      /<img/gi,
      '<img referrerpolicy="no-referrer"',
    );
    scrollTop = previewRef.current?.scrollTop || 0;
    setHtml(previewHtml);
  }, [value]);
  useEffect(() => {
    if (!html) {
      return;
    }

    previewRef.current?.scrollTo(0, scrollTop);
    htmlRender(previewRef.current);
  }, [html]);
  useImperativeHandle(ref, () => {
    return {
      getHtml: () => html,
    };
  });

  return (
    <div
      ref={previewRef}
      className="preview-wrap position-relative p-3 bg-light rounded text-break text-wrap mt-2 fmt"
      dangerouslySetInnerHTML={{ __html: html }}
    />
  );
};

export default memo(forwardRef(Index));

import {
  forwardRef,
  useEffect,
  useRef,
  useState,
  memo,
  useImperativeHandle,
} from 'react';

import { markdownToHtml } from '@/services';
import ImgViewer from '@/components/ImgViewer';

import { htmlRender } from './utils';

let scrollTop = 0;
let renderTimer;

const Index = ({ value }, ref) => {
  const [html, setHtml] = useState('');
  const previewRef = useRef<HTMLDivElement>(null);

  const renderMarkdown = (markdown) => {
    clearTimeout(renderTimer);
    const timeout = renderTimer ? 1000 : 0;
    renderTimer = setTimeout(() => {
      markdownToHtml(markdown).then((resp) => {
        scrollTop = previewRef.current?.scrollTop || 0;
        setHtml(resp);
      });
    }, timeout);
  };
  useEffect(() => {
    renderMarkdown(value);
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
    <ImgViewer>
      <div
        ref={previewRef}
        className="preview-wrap position-relative p-3 bg-light rounded text-break text-wrap mt-2 fmt"
        dangerouslySetInnerHTML={{ __html: html }}
      />
    </ImgViewer>
  );
};

export default memo(forwardRef(Index));

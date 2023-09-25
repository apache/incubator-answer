import {
  useEffect,
  useRef,
  useState,
  ForwardRefRenderFunction,
  forwardRef,
  useImperativeHandle,
} from 'react';

import classNames from 'classnames';

import PluginRender from '../PluginRender';

import {
  BlockQuote,
  Bold,
  Code,
  Heading,
  Help,
  Hr,
  Image,
  Indent,
  Italice,
  Link as LinkItem,
  OL,
  Outdent,
  Table,
  UL,
} from './ToolBars';
import { htmlRender, useEditor } from './utils';
import Viewer from './Viewer';
import { EditorContext } from './EditorContext';

import './index.scss';

export interface EditorRef {
  getHtml: () => string;
}

interface EventRef {
  onChange?(value: string): void;
  onFocus?(): void;
  onBlur?(): void;
}

interface Props extends EventRef {
  editorPlaceholder?;
  className?;
  value;
  autoFocus?: boolean;
}

const MDEditor: ForwardRefRenderFunction<EditorRef, Props> = (
  {
    editorPlaceholder = '',
    className = '',
    value,
    onChange,
    onFocus,
    onBlur,
    autoFocus = false,
  },
  ref,
) => {
  const editorRef = useRef<HTMLDivElement>(null);
  const previewRef = useRef<{ getHtml; element } | null>(null);

  const editor = useEditor({
    editorRef,
    onChange,
    onFocus,
    onBlur,
    placeholder: editorPlaceholder,
    autoFocus,
  });

  const [markdown, setMarkdown] = useState<string>(value || '');

  useEffect(() => {
    if (value !== markdown) {
      setMarkdown(value);
    }
  }, [value]);

  const getHtml = () => {
    return previewRef.current?.getHtml();
  };

  useImperativeHandle(ref, () => ({
    getHtml,
  }));

  useEffect(() => {
    if (!editor) {
      return;
    }
    if (editor.getValue() !== value) {
      editor.setValue(value || '');
    }
  }, [editor, value]);

  return (
    <>
      <div className={classNames('md-editor-wrap rounded', className)}>
        <EditorContext.Provider value={editor}>
          <PluginRender
            type="editor"
            className="toolbar-wrap px-3 d-flex align-items-center flex-wrap"
            editor={editor}
            previewElement={previewRef.current?.element}>
            <Heading />
            <Bold />
            <Italice />
            <div className="toolbar-divider" />
            <Code />
            <LinkItem />
            <BlockQuote />
            <Image />
            <Table />
            <div className="toolbar-divider" />
            <OL />
            <UL />
            <Indent />
            <Outdent />
            <Hr />
            <div className="toolbar-divider" />
            <Help />
          </PluginRender>
        </EditorContext.Provider>

        <div className="content-wrap">
          <div
            className="md-editor position-relative w-100 h-100"
            ref={editorRef}
          />
        </div>
      </div>
      <Viewer ref={previewRef} value={markdown} />
    </>
  );
};
export { htmlRender };
export default forwardRef(MDEditor);

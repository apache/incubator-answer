import {
  useEffect,
  useRef,
  useState,
  ForwardRefRenderFunction,
  forwardRef,
  useImperativeHandle,
} from 'react';

import classNames from 'classnames';

import {
  BlockQuote,
  Bold,
  Chart,
  Code,
  Formula,
  Heading,
  Help,
  Hr,
  Image,
  Indent,
  Italice,
  Link,
  OL,
  Outdent,
  Table,
  UL,
} from './ToolBars';
import { createEditorUtils } from './utils';
import Viewer from './Viewer';
import { CodeMirrorEditor, IEditorContext } from './types';
import { EditorContext } from './EditorContext';
import Editor from './Editor';
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
}

const MDEditor: ForwardRefRenderFunction<EditorRef, Props> = (
  { editorPlaceholder = '', className = '', value, onChange, onFocus, onBlur },
  ref,
) => {
  const [markdown, setMarkdown] = useState<string>(value || '');
  const previewRef = useRef<{ getHtml } | null>(null);
  const [editor, setEditor] = useState<CodeMirrorEditor | null>(null);
  const [context, setContext] = useState<IEditorContext | null>(null);
  const eventRef = useRef<EventRef>();

  useEffect(() => {
    if (!editor) {
      return;
    }

    import('codemirror').then(({ default: codemirror }) => {
      setContext({
        editor,
        ...createEditorUtils(codemirror, editor),
      });
    });
  }, [editor]);

  useEffect(() => {
    if (value !== markdown) {
      setMarkdown(value);
    }
  }, [value]);

  useEffect(() => {
    eventRef.current = {
      onChange,
      onFocus,
      onBlur,
    };
  }, [onChange, onFocus, onBlur]);

  const getEditorInstance = (cm) => {
    setEditor(cm);
  };

  const getHtml = () => {
    return previewRef.current?.getHtml();
  };

  const handleChange = (val) => {
    setMarkdown(val);
    eventRef.current?.onChange?.(val);
  };

  const handleFocus = () => {
    eventRef.current?.onFocus?.();
  };

  const handleBlur = () => {
    eventRef.current?.onBlur?.();
  };

  useImperativeHandle(ref, () => ({
    getHtml,
  }));

  return (
    <>
      <div className={classNames('md-editor-wrap rounded', className)}>
        <EditorContext.Provider value={context}>
          {context && (
            <div className="toolbar-wrap px-3 d-flex align-items-center flex-wrap">
              <Heading {...context} />
              <Bold {...context} />
              <Italice {...context} />
              <div className="toolbar-divider" />
              <Code {...context} />
              <Link {...context} />
              <BlockQuote {...context} />
              <Image {...context} />
              <div className="toolbar-divider" />
              <Table {...context} />
              <Formula {...context} />
              <Chart {...context} />
              <div className="toolbar-divider" />
              <OL {...context} />
              <UL {...context} />
              <Indent {...context} />
              <Outdent {...context} />
              <Hr {...context} />
              <div className="toolbar-divider" />
              <Help />
            </div>
          )}
        </EditorContext.Provider>

        <div className="content-wrap">
          <Editor
            value={markdown}
            onChange={handleChange}
            onFocus={handleFocus}
            onBlur={handleBlur}
            editorPlaceholder={editorPlaceholder}
            getEditorInstance={getEditorInstance}
          />
        </div>
      </div>
      <Viewer ref={previewRef} value={markdown} />
    </>
  );
};
export default forwardRef(MDEditor);

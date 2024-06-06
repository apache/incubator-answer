/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import {
  useEffect,
  useRef,
  ForwardRefRenderFunction,
  forwardRef,
  useImperativeHandle,
} from 'react';

import classNames from 'classnames';

import { PluginType } from '@/utils/pluginKit';
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
          {editor && (
            <PluginRender
              type={PluginType.Editor}
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
              <Image editorInstance={editor} />
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
          )}
        </EditorContext.Provider>

        <div className="content-wrap">
          <div
            className="md-editor position-relative w-100 h-100"
            ref={editorRef}
          />
        </div>
      </div>
      <Viewer ref={previewRef} value={value} />
    </>
  );
};
export { htmlRender };
export default forwardRef(MDEditor);

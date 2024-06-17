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

import { useEffect, useState } from 'react';

import { minimalSetup } from 'codemirror';
import { EditorState, Compartment } from '@codemirror/state';
import { EditorView, placeholder } from '@codemirror/view';
import { markdown, markdownLanguage } from '@codemirror/lang-markdown';
import { languages } from '@codemirror/language-data';

import { Editor } from '../types';

import createEditorUtils from './extension';

const editableCompartment = new Compartment();
export function htmlRender(el: HTMLElement | null) {
  if (!el) return;
  // Replace all br tags with newlines
  // Fixed an issue where the BR tag in the editor block formula HTML caused rendering errors.
  el.querySelectorAll('p').forEach((p) => {
    if (p.innerHTML.startsWith('$$') && p.innerHTML.endsWith('$$')) {
      const str = p.innerHTML.replace(/<br>/g, '\n');
      p.innerHTML = str;
    }
  });

  // change table style

  el.querySelectorAll('table').forEach((table) => {
    if (
      (table.parentNode as HTMLDivElement)?.classList.contains(
        'table-responsive',
      )
    ) {
      return;
    }

    table.classList.add('table', 'table-bordered');
    const div = document.createElement('div');
    div.className = 'table-responsive';
    table.parentNode?.replaceChild(div, table);
    div.appendChild(table);
  });

  // add rel nofollow for link not includes domain
  el.querySelectorAll('a').forEach((a) => {
    const base = window.location.origin;
    const targetUrl = new URL(a.href, base);

    if (targetUrl.origin !== base) {
      a.rel = 'nofollow';
    }
  });
}

export const useEditor = ({
  editorRef,
  placeholder: placeholderText,
  autoFocus,
  onChange,
  onFocus,
  onBlur,
}) => {
  const [editor, setEditor] = useState<Editor | null>(null);
  const [value, setValue] = useState<string>('');
  const init = async () => {
    const theme = EditorView.theme({
      '&': {
        height: '100%',
      },
      '&.cm-focused': {
        outline: 'none',
      },
      '.cm-content': {
        width: '100%',
        padding: '1rem',
      },
      '.cm-line': {
        whiteSpace: 'pre-wrap',
        wordWrap: 'break-word',
        wordBreak: 'break-all',
      },
    });

    const startState = EditorState.create({
      extensions: [
        minimalSetup,
        markdown({
          codeLanguages: languages,
          base: markdownLanguage,
        }),
        theme,
        placeholder(placeholderText),
        EditorView.lineWrapping,
        editableCompartment.of(EditorView.editable.of(true)),
      ],
    });

    const view = new EditorView({
      parent: editorRef.current,
      state: startState,
    });

    const cm = createEditorUtils(view as Editor);

    cm.setReadOnly = (readOnly: boolean) => {
      cm.dispatch({
        effects: editableCompartment.reconfigure(
          EditorView.editable.of(!readOnly),
        ),
      });
    };

    if (autoFocus) {
      setTimeout(() => {
        cm.focus();
      }, 10);
    }

    cm.on('change', () => {
      const newValue = cm.getValue();
      setValue(newValue);
    });

    cm.on('focus', () => {
      onFocus?.();
    });

    cm.on('blur', () => {
      onBlur?.();
    });

    setEditor(cm);

    return cm;
  };

  useEffect(() => {
    onChange?.(value);
  }, [value]);

  useEffect(() => {
    if (!editorRef.current) {
      return;
    }
    if (editorRef.current.children.length > 0 || editor) {
      return;
    }

    init();
  }, [editor]);
  return editor;
};

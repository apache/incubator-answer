/* eslint-disable */
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

import type { Position } from 'codemirror';
import { EditorSelection, EditorState, StateEffect } from '@codemirror/state';
import { EditorView, keymap, KeyBinding, Command } from '@codemirror/view';
import { markdown } from '@codemirror/lang-markdown';
import type CodeMirror from 'codemirror';
import 'codemirror/lib/codemirror.css';

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

  // add rel nofollow for link not inlcludes domain
  el.querySelectorAll('a').forEach((a) => {
    const base = window.location.origin;
    const targetUrl = new URL(a.href, base);

    if (targetUrl.origin !== base) {
      a.rel = 'nofollow';
    }
  });
}
export interface ExtendEditor {
  addKeyMap: (keyMap: Record<string, Command>) => void;
  on: (
    event:
      | 'change'
      | 'focus'
      | 'blur'
      | 'dragenter'
      | 'dragover'
      | 'drop'
      | 'paste',
    callback: (e?) => void,
  ) => void;
  getValue: () => string;
  setValue: (value: string) => void;
  off: (
    event:
      | 'change'
      | 'focus'
      | 'blur'
      | 'dragenter'
      | 'dragover'
      | 'drop'
      | 'paste',
    callback: (e?) => void,
  ) => void;
  getSelection: () => string;
  replaceSelection: (value: string) => void;
  focus: () => void;
  wrapText: (before: string, after?: string, defaultText?: string) => void;
  replaceLines: (
    replace: Parameters<Array<string>['map']>[0],
    symbolLen?: number,
  ) => void;
  appendBlock: (content: string) => Position;
  getCursor: () => Position;
  replaceRange: (value: string, from: Position, to: Position) => void;
}

const createEditorUtils = (editor: EditorView & ExtendEditor) => {
  editor.focus = () => {
    editor.contentDOM.focus();
  };

  editor.getCursor = () => {
    const range = editor.state.selection.ranges[0];
    const line = editor.state.doc.lineAt(range.from).number;
    const from = editor.state.doc.line(line).from;
    const to = editor.state.doc.line(line).to;
    return { from, to, ch: range.from - from };
  };

  editor.addKeyMap = (keyMap) => {
    const array = Object.entries(keyMap).map(([key, value]) => {
      const keyBinding: KeyBinding = {
        key,
        preventDefault: true,
        run: value,
      };
      return keyBinding;
    });

    editor.dispatch({
      effects: StateEffect.appendConfig.of(keymap.of(array)),
    });
  };

  editor.getSelection = () => {
    return editor.state.sliceDoc(
      editor.state.selection.main.from,
      editor.state.selection.main.to,
    );
  };

  editor.replaceSelection = (value: string) => {
    editor.dispatch({
      changes: [
        {
          from: editor.state.selection.main.from,
          to: editor.state.selection.main.to,
          insert: value,
        },
      ],
      selection: EditorSelection.cursor(
        editor.state.selection.main.from + value.length,
      ),
    });
  };

  editor.on = (event, callback) => {
    if (event === 'change') {
      const change = EditorView.updateListener.of((update) => {
        if (update.docChanged) {
          callback();
        }
      });

      editor.dispatch({
        effects: StateEffect.appendConfig.of(change),
      });
    }
    if (event === 'focus') {
      editor.contentDOM.addEventListener('focus', callback);
    }
    if (event === 'blur') {
      editor.contentDOM.addEventListener('blur', callback);
    }

    if (event === 'dragenter') {
      editor.contentDOM.addEventListener('dragenter', callback);
    }

    if (event === 'dragover') {
      editor.contentDOM.addEventListener('dragover', callback);
    }

    if (event === 'drop') {
      editor.contentDOM.addEventListener('drop', callback);
    }

    if (event === 'paste') {
      editor.contentDOM.addEventListener('paste', callback);
    }
  };

  editor.off = (event, callback) => {
    if (event === 'focus') {
      editor.contentDOM.removeEventListener('focus', callback);
    }

    if (event === 'blur') {
      editor.contentDOM.removeEventListener('blur', callback);
    }

    if (event === 'dragenter') {
      editor.contentDOM.removeEventListener('dragenter', callback);
    }

    if (event === 'dragover') {
      editor.contentDOM.removeEventListener('dragover', callback);
    }

    if (event === 'drop') {
      editor.contentDOM.removeEventListener('drop', callback);
    }

    if (event === 'paste') {
      editor.contentDOM.removeEventListener('paste', callback);
    }
  };

  editor.getValue = () => {
    return editor.state.doc.toString();
  };

  editor.setValue = (value: string) => {
    editor.dispatch({
      changes: { from: 0, to: editor.state.doc.length, insert: value },
    });
  };

  editor.wrapText = (before: string, after = before, defaultText) => {
    const range = editor.state.selection.ranges[0];
    const selection = editor.state.sliceDoc(range.from, range.to);
    if (selection) {
      editor.replaceSelection(`${before}${selection}${after}`);
    } else {
      editor.replaceSelection(`${before}${defaultText}${after}`);
    }
  };
  editor.replaceLines = (replace: Parameters<Array<string>['map']>[0]) => {
    const range = editor.state.selection.ranges[0];
    const line = editor.state.doc.lineAt(range.from).number;
    const from = editor.state.doc.line(line).from;
    const to = editor.state.doc.line(line).to;
    const lines = editor.state.sliceDoc(from, to).split('\n');

    const insert = lines.map(replace).join('\n');
    const selectionStart = from;
    const selectionEnd = from + insert.length;

    editor.dispatch({
      changes: [
        {
          from,
          to,
          insert,
        },
      ],
      selection: EditorSelection.single(selectionStart, selectionEnd),
    });
  };

  editor.appendBlock = (content: string): Position => {
    const range = editor.state.selection.ranges[0];
    const line = editor.state.doc.lineAt(range.from).number;
    const from = editor.state.doc.line(line).from;
    const to = editor.state.doc.line(line).to;
    let insert = `\n\n${content}`;
    let selection = EditorSelection.single(to, to + content.length);
    if (from === to) {
      insert = `${content}\n`;
      selection = EditorSelection.single(from, from + content.length);
    }

    editor.dispatch({
      changes: [
        {
          from: to,
          insert,
        },
      ],
      selection,
    });
  };

  editor.replaceRange = (
    value: string,
    selectionStart: Position,
    selectionEnd: Position,
  ) => {
    const from = selectionStart.from;
    const to = selectionEnd.to + selectionEnd.ch;
    editor.dispatch({
      changes: [
        {
          from,
          to,
          insert: value,
        },
      ],
      selection: EditorSelection.cursor(from + value.length),
    });
  };

  return editor;
};
export const useEditor = ({
  editorRef,
  placeholder,
  autoFocus,
  onChange,
  onFocus,
  onBlur,
}) => {
  const [editor, setEditor] = useState<CodeMirror.Editor | null>(null);
  const [value, setValue] = useState<string>('');
  const init = async () => {
    const theme = EditorView.theme({
      '&': {
        width: '100%',
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
    let startState = EditorState.create({
      extensions: [markdown(), theme],
    });

    const view = new EditorView({
      parent: editorRef.current,
      state: startState,
    });

    const editor = createEditorUtils(view as EditorView & ExtendEditor);

    if (autoFocus) {
      setTimeout(() => {
        editor.focus();
      }, 10);
    }

    editor.on('change', () => {
      const newValue = editor.getValue();
      setValue(newValue);
    });

    editor.on('focus', () => {
      onFocus?.();
    });

    editor.on('blur', () => {
      onBlur?.();
    });

    setEditor(editor);

    return editor;
  };

  useEffect(() => {
    onChange?.(value);
  }, [value]);

  useEffect(() => {
    if (!(editorRef.current instanceof HTMLElement) || editor) {
      return;
    }
    init();
  }, [editor]);
  return editor;
};

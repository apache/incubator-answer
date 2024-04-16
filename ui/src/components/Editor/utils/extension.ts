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

import { EditorSelection, StateEffect } from '@codemirror/state';
import { EditorView, keymap, KeyBinding } from '@codemirror/view';

import { Editor, Position } from '../types';

const createEditorUtils = (editor: Editor) => {
  editor.focus = () => {
    editor.contentDOM.focus();
  };

  editor.getCursor = () => {
    const range = editor.state.selection.ranges[0];
    const line = editor.state.doc.lineAt(range.from).number;
    const { from, to } = editor.state.doc.line(line);
    return { from, to, ch: range.from - from, line };
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

  editor.setSelection = (anchor: Position, head?: Position) => {
    editor.dispatch({
      selection: EditorSelection.create([
        EditorSelection.range(
          editor.state.doc.line(anchor.line).from + anchor.ch,
          head
            ? editor.state.doc.line(head.line).from + head.ch
            : editor.state.doc.line(anchor.line).from + anchor.ch,
        ),
      ]),
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
    const text = `${before}${selection || defaultText}${after}`;

    editor.dispatch({
      changes: [
        {
          from: range.from,
          to: range.to,
          insert: text,
        },
      ],
      selection: EditorSelection.range(
        range.from + before.length,
        range.to + before.length,
      ),
    });
  };

  editor.replaceLines = (
    replace: Parameters<Array<string>['map']>[0],
    symbolLen = 0,
  ) => {
    const range = editor.state.selection.ranges[0];
    const line = editor.state.doc.lineAt(range.from).number;
    const { from, to } = editor.state.doc.line(line);
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
      selection: EditorSelection.create([
        EditorSelection.range(selectionStart + symbolLen, selectionEnd),
      ]),
    });
  };

  editor.appendBlock = (content: string) => {
    const range = editor.state.selection.ranges[0];
    const line = editor.state.doc.lineAt(range.from).number;
    const { from, to } = editor.state.doc.line(line);

    let insert = `\n\n${content}`;

    let selection = EditorSelection.single(to, to + content.length);
    if (from === to) {
      insert = `${content}\n`;
      selection = EditorSelection.create([
        EditorSelection.cursor(to + content.length),
      ]);
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
    const from =
      editor.state.doc.line(selectionStart.line).from + selectionStart.ch;
    const to = editor.state.doc.line(selectionEnd.line).from + selectionEnd.ch;
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

export default createEditorUtils;

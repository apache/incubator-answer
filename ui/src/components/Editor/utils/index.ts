import { useEffect, useState } from 'react';

import type { Editor, Position } from 'codemirror';
import type CodeMirror from 'codemirror';
import 'codemirror/lib/codemirror.css';

export function createEditorUtils(
  codemirror: typeof CodeMirror,
  editor: Editor,
) {
  editor.wrapText = (before: string, after = before, defaultText) => {
    const range = editor.somethingSelected()
      ? editor.listSelections()[0]
      : editor.findWordAt(editor.getCursor());

    const from = range.from();
    const to = range.to();
    const text = editor.getRange(from, to) || defaultText;
    const fromBefore = codemirror.Pos(from.line, from.ch - before.length);
    const toAfter = codemirror.Pos(to.line, to.ch + after.length);

    if (
      editor.getRange(fromBefore, from) === before &&
      editor.getRange(to, toAfter) === after
    ) {
      editor.replaceRange(text, fromBefore, toAfter);
      editor.setSelection(
        fromBefore,
        codemirror.Pos(fromBefore.line, fromBefore.ch + text.length),
      );
    } else {
      editor.replaceRange(before + text + after, from, to);
      const cursor = editor.getCursor();

      editor.setSelection(
        codemirror.Pos(cursor.line, cursor.ch - after.length - text.length),
        codemirror.Pos(cursor.line, cursor.ch - after.length),
      );
    }
  };
  editor.replaceLines = (
    replace: Parameters<Array<string>['map']>[0],
    symbolLen = 0,
  ) => {
    const [selection] = editor.listSelections();

    const range = [
      codemirror.Pos(selection.from().line, 0),
      codemirror.Pos(selection.to().line),
    ] as const;
    const lines = editor.getRange(...range).split('\n');

    editor.replaceRange(lines.map(replace).join('\n'), ...range);
    const newRange = range;

    if (symbolLen > 0) {
      newRange[0].ch = symbolLen;
    }
    editor.setSelection(...newRange);
  };
  editor.appendBlock = (content: string): Position => {
    const cursor = editor.getCursor();

    let emptyLine = -1;

    for (let i = cursor.line; i < editor.lineCount(); i += 1) {
      if (!editor.getLine(i).trim()) {
        emptyLine = i;
        break;
      }
    }
    if (emptyLine === -1) {
      editor.replaceRange('\n', codemirror.Pos(editor.lineCount()));
      emptyLine = editor.lineCount();
    }

    editor.replaceRange(`\n${content}`, codemirror.Pos(emptyLine));
    return codemirror.Pos(emptyLine + 1, 0);
  };
}

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

  const onEnter = (cm) => {
    const cursor = cm.getCursor();
    const text = cm.getLine(cursor.line);
    const doc = cm.getDoc();

    const olRegexData = text.match(/^(\s{0,})(\d+)\.\s/);
    const ulRegexData = text.match(/^(\s{0,})(-|\*)\s/);
    const blockquoteData = text.match(/^>\s+?/g);

    if (olRegexData && text !== olRegexData[0]) {
      const num = olRegexData[2];

      doc.replaceSelection(`\n${olRegexData[1]}${Number(num) + 1}. `);
    } else if (ulRegexData && text !== ulRegexData[0]) {
      doc.replaceSelection(`\n${ulRegexData[1]}${ulRegexData[2]} `);
    } else if (blockquoteData && text !== blockquoteData[0]) {
      doc.replaceSelection(`\n> `);
    } else if (
      text.trim() === '>' ||
      text.trim().match(/^\d{1,}\.$/) ||
      text.trim().match(/^(\*|-)$/)
    ) {
      doc.replaceRange(`\n`, { ...cursor, ch: 0 }, cursor);
    } else {
      doc.replaceSelection(`\n`);
    }
  };

  const init = async () => {
    const { default: codeMirror } = await import('codemirror');
    await import('codemirror/mode/markdown/markdown');
    await import('codemirror/addon/display/placeholder');

    const cm = codeMirror(editorRef?.current, {
      mode: 'markdown',
      lineWrapping: true,
      placeholder,
      focus: autoFocus,
    });

    setEditor(cm);
    createEditorUtils(codeMirror, cm);

    cm.on('change', (e) => {
      const newValue = e.getValue();
      setValue(newValue);
    });

    cm.on('focus', () => {
      onFocus?.();
    });
    cm.on('blur', () => {
      onBlur?.();
    });
    cm.setSize('100%', '100%');
    cm.addKeyMap({
      Enter: onEnter,
    });
    return cm;
  };

  useEffect(() => {
    onChange?.(value);
  }, [value]);

  useEffect(() => {
    if (!(editorRef.current instanceof HTMLElement)) {
      return;
    }
    init();
  }, [editorRef]);

  return editor;
};

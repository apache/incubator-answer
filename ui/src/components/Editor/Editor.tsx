import { useEffect, useRef, useState } from 'react';

import 'codemirror/lib/codemirror.css';
import type CodeMirror from 'codemirror';

export interface EditorInstance {
  editor: CodeMirror.Editor | null;
}

export interface EditorProps {
  onChange?(value: string): void;
  onFocus?(): void;
  onBlur?(): void;
}
const Editor = ({
  value,
  onChange,
  onFocus,
  onBlur,
  editorPlaceholder,
  getEditorInstance,
  autoFocus,
}) => {
  const elRef = useRef<HTMLDivElement>(null);
  const [editor, setEditor] = useState<CodeMirror.Editor | null>(null);
  const eventRef = useRef<EditorProps>();
  const isMountedRef = useRef(false);

  useEffect(() => {
    const el = elRef?.current;
    if (!isMountedRef.current && el instanceof HTMLElement) {
      isMountedRef.current = true;
      import('codemirror').then(async ({ default: CodeMirror }) => {
        await import('codemirror/mode/markdown/markdown');
        await import('codemirror/addon/display/placeholder');

        const cm = CodeMirror(el, {
          mode: 'markdown',
          lineWrapping: true,
          placeholder: editorPlaceholder,
        });
        if (autoFocus) {
          cm.focus();
        }
        cm.on('change', (e) => {
          const newValue = e.getValue();
          eventRef.current?.onChange?.(newValue);
        });

        cm.on('focus', () => {
          eventRef.current?.onFocus?.();
        });
        cm.on('blur', () => {
          eventRef.current?.onBlur?.();
        });
        setEditor(cm);
        getEditorInstance(cm);
        cm.setSize('100%', '100%');
        cm.addKeyMap({
          Enter: () => {
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
          },
        });
      });
    }
  }, [elRef]);

  useEffect(() => {
    if (!editor) {
      return;
    }
    if (editor.getValue() !== value) {
      editor.setValue(value || '');
    }
  }, [editor, value]);

  useEffect(() => {
    eventRef.current = {
      onChange,
      onFocus,
      onBlur,
    };
  }, [onChange, onFocus, onBlur]);

  return (
    <div className="md-editor position-relative w-100 h-100" ref={elRef} />
  );
};

export default Editor;

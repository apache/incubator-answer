import type { Editor, Position } from 'codemirror';
import type CodeMirror from 'codemirror';
import 'katex/dist/katex.min.css';

export function createEditorUtils(
  codemirror: typeof CodeMirror,
  editor: Editor,
) {
  return {
    wrapText(before: string, after = before, defaultText) {
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
    },

    replaceLines(replace: Parameters<Array<string>['map']>[0], symbolLen = 0) {
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
    },

    appendBlock(content: string): Position {
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
    },
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

  import('mermaid').then(({ default: mermaid }) => {
    mermaid.initialize({ startOnLoad: false });

    el.querySelectorAll('.language-mermaid').forEach((pre) => {
      const flag = Date.now();
      mermaid.render(
        `theGraph${flag}`,
        pre.textContent || '',
        function (svgCode) {
          const p = document.createElement('p');
          p.className = 'text-center';
          p.innerHTML = svgCode;

          pre.parentNode?.replaceChild(p, pre);
        },
      );
    });
  });
  import('katex/contrib/auto-render/auto-render').then(
    ({ default: render }) => {
      render(el, {
        delimiters: [
          { left: '$$', right: '$$', display: true },
          { left: '$$<br>', right: '<br>$$', display: true },
          {
            left: '\\begin{equation}',
            right: '\\end{equation}',
            display: true,
          },
          { left: '\\begin{align}', right: '\\end{align}', display: true },
          { left: '\\begin{alignat}', right: '\\end{alignat}', display: true },
          { left: '\\begin{gather}', right: '\\end{gather}', display: true },
          { left: '\\(', right: '\\)', display: false },
          { left: '\\[', right: '\\]', display: true },
        ],
      });
    },
  );

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

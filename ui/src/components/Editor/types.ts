import type { Editor } from 'codemirror';

export interface CodeMirrorEditor extends Editor {
  display: any;

  moduleType;
}

export interface IEditorContext {
  editor: Editor;
  wrapText?;
  replaceLines?;
  appendBlock?;
}

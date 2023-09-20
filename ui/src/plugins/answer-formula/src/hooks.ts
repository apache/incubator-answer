import { useEffect } from 'react';

// @ts-ignore
import katexRender from 'katex/contrib/auto-render/auto-render';

const useRenderFormula = (element: HTMLElement) => {
  const render = (element) => {
    katexRender(element, {
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
  };
  useEffect(() => {
    if (!element) {
      return;
    }

    render(element);
    const observer = new MutationObserver(() => {
      render(element);
    });

    observer.observe(element, {
      childList: true,
      attributes: true,
      subtree: true,
    });
  }, [element]);
};

export { useRenderFormula };

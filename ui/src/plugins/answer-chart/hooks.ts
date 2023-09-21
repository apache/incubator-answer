import { useEffect } from 'react';

// @ts-ignore
import mermaid from 'mermaid';

const useRenderChart = (element: HTMLElement) => {
  const render = (element) => {
    mermaid.initialize({ startOnLoad: false });
    element.querySelectorAll('.language-mermaid').forEach((pre) => {
      const flag = Date.now();
      mermaid.render(
        `theGraph${flag}`,
        pre.textContent || '',
        function (svgCode: string) {
          const p = document.createElement('p');
          p.className = 'text-center';
          p.innerHTML = svgCode;
          pre.parentNode?.replaceChild(p, pre);
        },
      );
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

export { useRenderChart };

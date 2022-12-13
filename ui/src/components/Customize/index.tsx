import { FC, memo, useEffect } from 'react';

import { customizeStore } from '@/stores';

const getElementByAttr = (attr: string, elName: string) => {
  let el = document.querySelector(`[${attr}]`);
  if (!el) {
    el = document.createElement(elName);
    el.setAttribute(attr, '');
  }
  return el;
};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const textToDf = (t) => {
  const dummyDoc = document.createElement('div');
  dummyDoc.innerHTML = t;
  const frag = document.createDocumentFragment();
  while (dummyDoc.childNodes.length) {
    frag.appendChild(dummyDoc.children[0]);
  }
  return frag;
};

const injectCustomCSS = (t: string) => {
  if (!t) {
    return;
  }
  const el = getElementByAttr('data-custom-css', 'style');
  el.textContent = t;
  document.head.insertBefore(el, document.head.lastChild);
};

const injectCustomHead = (t: string) => {
  if (!t) {
    return;
  }
  setTimeout(() => {
    const el = getElementByAttr('data-custom-head', 'style');
    el.textContent = t;
    document.head.appendChild(el);
  }, 200);
};

const injectCustomHeader = (t: string) => {
  if (!t) {
    return;
  }
  // const frag = textToDf(t);
  t = ' Customize Header ';
  document.body.insertBefore(
    document.createComment(t),
    document.body.firstChild,
  );
};

const injectCustomFooter = (t: string) => {
  if (!t) {
    return;
  }
  // FIXME
  t = ' Customize Footer ';
  // const frag = textToDf(t);
  document.documentElement.appendChild(document.createComment(t));
};

const Index: FC = () => {
  const { custom_css, custom_head, custom_header, custom_footer } =
    customizeStore((state) => state);
  useEffect(() => {
    injectCustomCSS(custom_css);
    injectCustomHead(custom_head);
    injectCustomHeader(custom_header);
    injectCustomFooter(custom_footer);
  }, []);
  return (
    <>
      {null}
      {/* App customize */}
    </>
  );
};

export default memo(Index);

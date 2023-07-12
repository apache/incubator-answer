import i18next from 'i18next';

import pattern from '@/common/pattern';
import { USER_AGENT_NAMES } from '@/common/constants';

const Diff = require('diff');

function thousandthDivision(num) {
  const reg = /\d{1,3}(?=(\d{3})+$)/g;
  return `${num}`.replace(reg, '$&,');
}

function formatCount($num: number): string {
  let res = String($num);
  if (!Number.isFinite($num)) {
    res = '0';
  } else if ($num < 10000) {
    res = thousandthDivision($num);
  } else if ($num < 1000000) {
    res = `${Math.round($num / 100) / 10}k`;
  } else if ($num >= 1000000) {
    res = `${Math.round($num / 100000) / 10}m`;
  }
  return res;
}

function scrollToElementTop(element) {
  if (!element) {
    return;
  }
  const offset = 120;
  const bodyRect = document.body.getBoundingClientRect().top;
  const elementRect = element.getBoundingClientRect().top;
  const elementPosition = elementRect - bodyRect;
  const offsetPosition = elementPosition - offset;

  window.scrollTo({
    top: offsetPosition,
    behavior: 'instant' as ScrollBehavior,
  });
}

const scrollToDocTop = () => {
  setTimeout(() => {
    window.scrollTo({
      top: 0,
      left: 0,
      behavior: 'instant' as ScrollBehavior,
    });
  });
};

const bgFadeOut = (el) => {
  if (el && !el.classList.contains('bg-fade-out')) {
    el.classList.add('bg-fade-out');
    setTimeout(() => {
      el.classList.remove('bg-fade-out');
    }, 3200);
  }
};

/**
 * Extract user info from markdown
 * @param markdown string
 * @returns Array<{displayName: string, userName: string}>
 */
function matchedUsers(markdown) {
  const globalReg = /\B@([\w|]+)/g;
  const reg = /\B@([\w\\_\\.]+)/;

  const users = markdown.match(globalReg);
  if (!users) {
    return [];
  }
  return users.map((user) => {
    const matched = user.match(reg);
    return {
      userName: matched[1],
    };
  });
}

/**
 * Identify user information from markdown
 * @param markdown string
 * @returns string
 */
function parseUserInfo(markdown) {
  const globalReg = /\B@([\w\\_\\.\\-]+)/g;
  return markdown.replace(globalReg, '[@$1](/u/$1)');
}

function formatUptime(value) {
  const t = i18next.t.bind(i18next);
  const second = parseInt(value, 10);

  if (second > 60 * 60 && second < 60 * 60 * 24) {
    return `${Math.floor(second / 3600)} ${t('dates.hour')}`;
  }
  if (second > 60 * 60 * 24) {
    return `${Math.floor(second / 3600 / 24)} ${t('dates.day')}`;
  }

  return `< 1 ${t('dates.hour')}`;
}

function escapeRemove(str: string) {
  if (!str || typeof str !== 'string') return str;
  let temp: HTMLDivElement | null = document.createElement('div');
  temp.innerHTML = str;
  const output = temp?.innerText || temp.textContent;
  temp = null;
  return output;
}
function mixColor(color_1, color_2, weight) {
  function d2h(d) {
    return d.toString(16);
  }
  function h2d(h) {
    return parseInt(h, 16);
  }

  weight = typeof weight !== 'undefined' ? weight : 50;
  let color = '#';

  for (let i = 0; i <= 5; i += 2) {
    const v1 = h2d(color_1.substr(i, 2));
    const v2 = h2d(color_2.substr(i, 2));
    let val = d2h(Math.floor(v2 + (v1 - v2) * (weight / 100.0)));

    while (val.length < 2) {
      val = `0${val}`;
    }

    color += val;
  }

  return color;
}

function colorRgb(sColor) {
  sColor = sColor.toLowerCase();
  const reg = /^#([0-9a-fA-f]{3}|[0-9a-fA-f]{6})$/;
  if (sColor && reg.test(sColor)) {
    if (sColor.length === 4) {
      let sColorNew = '#';
      for (let i = 1; i < 4; i += 1) {
        sColorNew += sColor.slice(i, i + 1).concat(sColor.slice(i, i + 1));
      }
      sColor = sColorNew;
    }
    const sColorChange: number[] = [];
    for (let i = 1; i < 7; i += 2) {
      sColorChange.push(parseInt(`0x${sColor.slice(i, i + 2)}`, 16));
    }
    return sColorChange.join(',');
  }
  return sColor;
}

function labelStyle(color, hover) {
  const textColor = mixColor('000000', color.replace('#', ''), 40);
  const backgroundColor = mixColor('ffffff', color.replace('#', ''), 80);
  const rgbBackgroundColor = colorRgb(backgroundColor);
  return {
    color: textColor,
    backgroundColor: `rgba(${colorRgb(rgbBackgroundColor)},${hover ? 1 : 0.5})`,
  };
}

function handleFormError(
  error: { list: Array<{ error_field: string; error_msg: string }> },
  data: any,
  keymap?: Array<{ from: string; to: string }>,
) {
  if (error.list?.length > 0) {
    error.list.forEach((item) => {
      if (keymap?.length) {
        const key = keymap.find((k) => k.from === item.error_field);
        if (key) {
          item.error_field = key.to;
        }
      }
      const errorFieldObject = data[item.error_field];
      if (errorFieldObject) {
        errorFieldObject.isInvalid = true;
        errorFieldObject.errorMsg = item.error_msg;
      }
    });
  }
  return data;
}

function escapeHtml(str: string) {
  const tagsToReplace = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#39;',
    '`': '&#96;',
  };
  return str.replace(/[&<>"'`]/g, (tag) => tagsToReplace[tag] || tag);
}

function diffText(newText: string, oldText?: string): string {
  if (!newText) {
    return '';
  }

  if (typeof oldText !== 'string') {
    return escapeHtml(newText);
  }
  const diff = Diff.diffChars(escapeHtml(oldText), escapeHtml(newText));
  const result = diff.map((part) => {
    if (part.added) {
      if (part.value.replace(/\n/g, '').length <= 0) {
        return `<span class="review-text-add d-block">${part.value.replace(
          /\n/g,
          '↵\n',
        )}</span>`;
      }
      return `<span class="review-text-add">${part.value}</span>`;
    }
    if (part.removed) {
      if (part.value.replace(/\n/g, '').length <= 0) {
        return `<span class="review-text-delete text-decoration-none d-block">${part.value.replace(
          /\n/g,
          '↵\n',
        )}</span>`;
      }
      return `<span class="review-text-delete">${part.value}</span>`;
    }

    return part.value;
  });

  return result.join('');
}

function base64ToSvg(base64: string) {
  // base64 to svg xml
  const svgxml = atob(base64);

  // svg add class btnSvg
  const parser = new DOMParser();
  const doc = parser.parseFromString(svgxml, 'image/svg+xml');
  const parseError = doc.querySelector('parsererror');
  const svg = doc.querySelector('svg');
  let str = '';
  if (svg && !parseError) {
    svg.classList.add('btnSvg');
    svg.classList.add('me-2');

    // transform svg to string
    const serializer = new XMLSerializer();
    str = serializer.serializeToString(doc);
  }
  return str;
}

// Determine whether the user is in WeChat or Enterprise WeChat or DingTalk, and return the corresponding type

function getUaType() {
  const ua = navigator.userAgent.toLowerCase();
  if (pattern.uaWeCom.test(ua)) {
    return USER_AGENT_NAMES.WeCom;
  }
  if (pattern.uaWeChat.test(ua)) {
    return USER_AGENT_NAMES.WeChat;
  }
  if (pattern.uaDingTalk.test(ua)) {
    return USER_AGENT_NAMES.DingTalk;
  }
  return null;
}

export {
  thousandthDivision,
  formatCount,
  scrollToElementTop,
  scrollToDocTop,
  bgFadeOut,
  matchedUsers,
  parseUserInfo,
  formatUptime,
  escapeRemove,
  mixColor,
  colorRgb,
  labelStyle,
  handleFormError,
  diffText,
  base64ToSvg,
  getUaType,
};

import i18next from 'i18next';

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

function scrollTop(element) {
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
  });
}

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

function escapeRemove(str) {
  if (!str || typeof str !== 'string') return str;
  const arrEntities = {
    lt: '<',
    gt: '>',
    nbsp: ' ',
    amp: '&',
    quot: '"',
    '#39': "'",
  };

  return str.replace(/&(lt|gt|nbsp|amp|quot|#39);/gi, function (all, t) {
    return arrEntities[t];
  });
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
) {
  if (error.list?.length > 0) {
    error.list.forEach((item) => {
      data[item.error_field].isInvalid = true;
      data[item.error_field].errorMsg = item.error_msg;
    });
  }
  return data;
}

function diffText(newText: string, oldText: string): string {
  if (!newText) {
    return '';
  }

  if (typeof oldText !== 'string') {
    return newText
      ?.replace(/\n/gi, '<br>')
      ?.replace(/<kbd/gi, '&lt;kbd')
      ?.replace(/<\/kbd>/gi, '&lt;/kbd&gt;')
      ?.replace(/<iframe/gi, '&lt;iframe')
      ?.replace(/<input/gi, '&lt;input');
  }
  const diff = Diff.diffChars(oldText, newText);
  const result = diff.map((part) => {
    if (part.added) {
      if (part.value.replace(/\n/g, '').length <= 0) {
        return `<span class="review-text-add d-block">${part.value.replace(
          /\n/g,
          '↵\n',
        )}</span>`;
      }
      return `<span class="review-text-add d-block">${part.value}</span>`;
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

  return result
    .join('')
    ?.replace(/<iframe/gi, '&lt;iframe')
    ?.replace(/<kbd/gi, '&lt;kbd')
    ?.replace(/<\/kbd>/gi, '&lt;/kbd&gt;')
    ?.replace(/<input/gi, '&lt;input');
}

export {
  thousandthDivision,
  formatCount,
  scrollTop,
  matchedUsers,
  parseUserInfo,
  formatUptime,
  escapeRemove,
  mixColor,
  colorRgb,
  labelStyle,
  handleFormError,
  diffText,
};

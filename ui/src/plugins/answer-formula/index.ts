import Formula from './Formula';
import i18nConfig from './i18n';
import { useRenderFormula } from './hooks';

export default {
  info: {
    type: 'editor',
    slug_name: 'formula',
  },
  component: Formula,
  i18nConfig,
  hooks: {
    useRender: [useRenderFormula],
  },
};

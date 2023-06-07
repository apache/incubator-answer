import { FC, ReactNode } from 'react';

import builtin from '@/plugins/builtin';
import * as plugins from '@/plugins';
import { Plugin } from '@/utils/pluginKit';

interface Props {
  slug_name: string;
  children?: ReactNode;
  [prop: string]: any;
}

const findPluginBySlugName: (l, n) => Plugin | null = (source, slug_name) => {
  let ret: Plugin | null = null;
  if (source) {
    Object.keys(source).forEach((k) => {
      const p = source[k];
      if (p && p.info && p.info.slug_name === slug_name && p.component) {
        ret = p;
      }
    });
  }

  return ret;
};

const Index: FC<Props> = ({ slug_name, children, ...props }) => {
  const bp = findPluginBySlugName(builtin, slug_name);
  const vp = findPluginBySlugName(plugins, slug_name);
  const plugin = bp || vp;

  if (!plugin) {
    return null;
  }
  /**
   * TODO: Rendering control for non-builtin plug-ins
   * ps: Logic such as version compatibility determination can be placed here
   */

  const PluginComponent = plugin.component;
  // @ts-ignore
  return <PluginComponent {...props}>{children}</PluginComponent>;
};

export default Index;

import { FC, ReactNode, memo } from 'react';

import builtin from '@/plugins/builtin';
import * as plugins from '@/plugins';
import { Plugin, PluginType } from '@/utils/pluginKit';

/**
 * Note：Please set at least either of the `slug_name` and `type` attributes, otherwise no plugins will be rendered.
 *
 * @field slug_name: The `slug_name` of the plugin needs to be rendered.
 *                   If this property is set, `PluginRender` will use it first (regardless of whether `type` is set)
 *                   to find the corresponding plugin and render it.
 * @field type: Used to formulate the rendering of all plugins of this type.
 *              (if the `slug_name` attribute is set, it will be ignored)
 * @field prop: Any attribute you want to configure, e.g. `className`
 */

interface Props {
  slug_name?: string;
  type?: PluginType;
  children?: ReactNode;
  [prop: string]: any;
}

const findPlugin: (s, k: 'slug_name' | 'type', v) => Plugin[] = (
  source,
  k,
  v,
) => {
  const ret: Plugin[] = [];
  if (source) {
    Object.keys(source).forEach((i) => {
      const p = source[i];
      if (p && p.component && p.info && p.info[k] === v) {
        ret.push(p);
      }
    });
  }
  return ret;
};

const Index: FC<Props> = ({ slug_name, type, children, ...props }) => {
  const fk = slug_name ? 'slug_name' : 'type';
  const fv = fk === 'slug_name' ? slug_name : type;
  const bp = findPlugin(builtin, fk, fv);
  const vp = findPlugin(plugins, fk, fv);
  const pluginSlice = [...bp, ...vp];

  if (!pluginSlice.length) {
    return null;
  }
  /**
   * TODO: Rendering control for non-builtin plug-ins
   * ps: Logic such as version compatibility determination can be placed here
   */

  return (
    <>
      {pluginSlice.map((ps) => {
        const PluginFC = ps.component;
        return (
          // @ts-ignore
          <PluginFC key={ps.info.slug_name} {...props}>
            {children}
          </PluginFC>
        );
      })}
    </>
  );
};

export default memo(Index);

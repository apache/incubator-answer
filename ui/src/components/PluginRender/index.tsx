import React, { FC, ReactNode, memo } from 'react';

import PluginKite, { Plugin, PluginType } from '@/utils/pluginKit';
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

const Index: FC<Props> = ({
  slug_name,
  type,
  children,
  className,
  ...props
}) => {
  const pluginSlice: Plugin[] = [];
  const plugins = PluginKite.getPlugins();

  plugins.forEach((plugin) => {
    if (type && slug_name) {
      if (plugin.info.slug_name === slug_name && plugin.info.type === type) {
        pluginSlice.push(plugin);
      }
    } else if (type) {
      if (plugin.info.type === type) {
        pluginSlice.push(plugin);
      }
    } else if (slug_name) {
      if (plugin.info.slug_name === slug_name) {
        pluginSlice.push(plugin);
      }
    }
  });

  /**
   * TODO: Rendering control for non-builtin plug-ins
   * ps: Logic such as version compatibility determination can be placed here
   */

  if (type === 'editor') {
    const nodes = React.Children.map(children, (child, index) => {
      if (index === 15) {
        return (
          <>
            {child}
            {pluginSlice.map((ps) => {
              const PluginFC = ps.component;
              return (
                // @ts-ignore
                <PluginFC key={ps.info.slug_name} {...props} />
              );
            })}
            <div className="toolbar-divider" />
          </>
        );
      }
      return child;
    });

    return <div className={className}>{nodes}</div>;
  }

  return (
    <>
      {pluginSlice.map((ps) => {
        const PluginFC = ps.component;
        return (
          // @ts-ignore
          <PluginFC key={ps.info.slug_name} className={className} {...props} />
        );
      })}
    </>
  );
};

export default memo(Index);

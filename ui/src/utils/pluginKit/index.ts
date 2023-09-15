import { NamedExoticComponent, FC } from 'react';

import builtin from '@/plugins/builtin';
import * as allPlugins from '@/plugins';

/**
 * This information is to be defined for all components.
 * It may be used for feature upgrades or version compatibility processing.
 *
 * @field slug_name: Unique identity string for the plugin, usually configured in `info.yaml`
 * @field type: The type of plugin is defined and a single type of plugin can have multiple implementations.
 *              For example, a plugin of type `connector` can have a `google` implementation and a `github` implementation.
 *              `PluginRender` automatically renders the plug-in types already included in `PluginType`.
 * @field name: Plugin name, optionally configurable. Usually read from the `i18n` file
 * @field description: Plugin description, optionally configurable. Usually read from the `i18n` file
 */

export type PluginType = 'connector' | 'search' | 'editor';
export interface PluginInfo {
  slug_name: string;
  type: PluginType;
  name?: string;
  description?: string;
}

export interface Plugin {
  info: PluginInfo;
  component: NamedExoticComponent | FC;
}

class Plugins {
  plugins: Plugin[] = [];

  constructor() {
    this.registerBuiltin();
    this.registerPlugins();
  }

  validate(plugin: Plugin) {
    if (!plugin) {
      return false;
    }
    const { info } = plugin;
    const { slug_name, type } = info;

    if (!slug_name) {
      return false;
    }

    if (!type) {
      return false;
    }

    return true;
  }

  registerBuiltin() {
    Object.keys(builtin).forEach((key) => {
      const plugin = builtin[key];
      this.register(plugin);
    });
  }

  registerPlugins() {
    Object.keys(allPlugins).forEach((key) => {
      const plugin = allPlugins[key];
      this.register(plugin);
    });
  }

  register(plugin: Plugin) {
    const bool = this.validate(plugin);
    if (!bool) {
      return;
    }
    this.plugins.push(plugin);
  }

  getPlugin(slug_name: string) {
    return this.plugins.find((p) => p.info.slug_name === slug_name);
  }

  getPluginsByType(type: PluginType) {
    return this.plugins.filter((p) => p.info.type === type);
  }

  getPlugins() {
    return this.plugins;
  }
}

const plugins = new Plugins();

export default plugins;

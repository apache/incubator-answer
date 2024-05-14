/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import React from 'react';

import builtin from '@/plugins/builtin';
import * as allPlugins from '@/plugins';
import type * as Type from '@/common/interface';
import { LOGGED_TOKEN_STORAGE_KEY } from '@/common/constants';
import { getPluginsStatus } from '@/services';
import Storage from '@/utils/storage';

import { initI18nResource, Plugin, PluginInfo, PluginType } from './utils';

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

class Plugins {
  plugins: Plugin[] = [];

  constructor() {
    this.registerBuiltin();
    this.registerPlugins();

    getPluginsStatus().then((plugins) => {
      this.activatePlugins(plugins);
    });
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
      // builttin plugins are always activated
      // Use own internal rendering logic'
      plugin.activated = true;
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
    if (plugin.i18nConfig) {
      initI18nResource(plugin.i18nConfig);
    }
    this.plugins.push(plugin);
  }

  activatePlugins(activatedPlugins: Type.ActivatedPlugin[]) {
    this.plugins.forEach((plugin: any) => {
      const { slug_name } = plugin.info;
      const activatedPlugin: any = activatedPlugins?.find(
        (p) => p.slug_name === slug_name,
      );
      if (activatedPlugin) {
        plugin.activated = activatedPlugin?.enabled;
      }
    });
  }

  changePluginActiveStatus(slug_name: string, active: boolean) {
    const plugin = this.getPlugin(slug_name);
    if (plugin) {
      plugin.activated = active;
    }
  }

  getPlugin(slug_name: string) {
    return this.plugins.find((p) => p.info.slug_name === slug_name);
  }

  getOnePluginHooks(slug_name: string) {
    const plugin = this.getPlugin(slug_name);
    return plugin?.hooks;
  }

  getPlugins() {
    return this.plugins;
  }
}

const plugins = new Plugins();

const useRenderHtmlPlugin = (element: HTMLElement | null) => {
  plugins
    .getPlugins()
    .filter((plugin) => plugin.activated && plugin.hooks?.useRender)
    .forEach((plugin) => {
      plugin.hooks?.useRender?.forEach((hook) => {
        hook(element);
      });
    });
};

const getRoutePlugins = () => {
  return plugins.getPlugins().filter((plugin) => plugin.info.type === 'route');
};

const defaultProps = () => {
  const token = Storage.get(LOGGED_TOKEN_STORAGE_KEY) || '';
  return {
    key: token,
    headers: {
      Authorization: token,
    },
  };
};

const mergeRoutePlugins = (routes) => {
  const routePlugins = getRoutePlugins();
  if (routePlugins.length === 0) {
    return routes;
  }
  routes.forEach((route) => {
    if (route.page === 'pages/Layout') {
      route.children?.forEach((child) => {
        if (child.page === 'pages/SideNavLayout') {
          routePlugins.forEach((plugin) => {
            const { slug_name, route: path } = plugin.info;
            const Component = plugin.component;

            child.children.push({
              page: `plugin/${slug_name}`,
              path,
              element: React.createElement(Component, defaultProps(), null),
            });
          });
        }
      });
    }
  });
  return routes;
};

// Only one captcha type plug-in can be enabled at the same time
const useCaptchaPlugin = (key: Type.CaptchaKey) => {
  const captcha = plugins
    .getPlugins()
    .filter((plugin) => plugin.info.type === 'captcha' && plugin.activated);
  const pluginHooks = plugins.getOnePluginHooks(captcha[0]?.info.slug_name);
  return pluginHooks?.useCaptcha?.({
    captchaKey: key,
    commonProps: defaultProps(),
  });
};

export type { Plugin, PluginInfo, PluginType };

export { useRenderHtmlPlugin, mergeRoutePlugins, useCaptchaPlugin };
export default plugins;

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

import { NamedExoticComponent, FC } from 'react';

import builtin from '@/plugins/builtin';
import * as allPlugins from '@/plugins';
import type * as Type from '@/common/interface';
import { getPluginsStatus } from '@/services';

import { initI18nResource } from './utils';
import type { PluginInfo } from './interface';

export interface Plugin {
  info: PluginInfo;
  component: NamedExoticComponent | FC;
  i18nConfig?;
  hooks?: {
    useRender?: Array<(element: HTMLElement | null) => void>;
    useCaptcha?: (key: Type.CaptchaKey) => {
      getCaptcha: () => Record<string, any>;
      check: (t: () => void) => void;
      handleCaptchaError: (error) => any;
      close: () => Promise<void>;
      resolveCaptchaReq: (data) => void;
    };
  };
  activated?: boolean;
}

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

// Only one captcha type plug-in can be enabled at the same time
const useCaptchaPlugin = (key: Type.CaptchaKey) => {
  const captcha = plugins
    .getPlugins()
    .filter((plugin) => plugin.info.type === 'captcha');
  const pluginHooks = plugins.getOnePluginHooks(captcha[0]?.info.slug_name);
  return pluginHooks?.useCaptcha?.(key);
};

export { useRenderHtmlPlugin, useCaptchaPlugin };
export default plugins;

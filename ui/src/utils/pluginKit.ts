import { NamedExoticComponent, FC } from 'react';

import i18next from 'i18next';

/**
 * This information is to be defined for all components.
 * It may be used for feature upgrades or version compatibility processing.
 *
 * @field slug_name: Unique identity string for the plugin, usually configured in `info.yaml`
 * @field type: The type of plugin is defined and a single type of plugin can have multiple implementations.
 *              For example, a plugin of type `Connector` can have a `google` implementation and a `github` implementation.
 *              `PluginRender` automatically renders the plug-in types already included in `PluginType`.
 * @field name: Plugin name, optionally configurable. Usually read from the `i18n` file
 * @field description: Plugin description, optionally configurable. Usually read from the `i18n` file
 */

const I18N_NS = 'plugin';

export type PluginType = 'Connector';
export interface PluginInfo {
  slug_name: string;
  type?: PluginType;
  name?: string;
  description?: string;
}

export interface Plugin {
  info: PluginInfo;
  component: NamedExoticComponent | FC;
}

interface I18nResource {
  [lng: string]: {
    plugin: {
      [slug_name: string]: {
        ui: any;
      };
    };
  };
}

const addResourceBundle = (resource: I18nResource) => {
  if (resource) {
    Object.keys(resource).forEach((lng) => {
      const r = resource[lng];

      i18next.addResourceBundle(lng, I18N_NS, r.plugin, true, true);
    });
  }
};

const initI18nResource = (resource: I18nResource) => {
  addResourceBundle(resource);
  /**
   * Note: In development mode,
   * when the base i18n file is changed, `i18next` will reinitialise the updated resource file,
   * which will cause the resource package added in the plugin to be lost
   * and will need to be automatically re-added by listening for events
   */
  i18next.on('initialized', () => {
    addResourceBundle(resource);
  });
};

const getTransNs = () => {
  return I18N_NS;
};

const getTransKeyPrefix = (info: PluginInfo) => {
  const kp = `${info.slug_name}.ui`;
  return kp;
};

export default {
  initI18nResource,
  getTransNs,
  getTransKeyPrefix,
};

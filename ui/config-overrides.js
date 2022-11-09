const path = require('path');
const i18nLocaleTool = require('./scripts/i18n-locale-tool');

module.exports = {
  webpack: function (config, env) {
    if (env === 'production') {
      config.output.publicPath = process.env.REACT_APP_PUBLIC_PATH;
      i18nLocaleTool.resolvePresetLocales();
    }

    for (let _rule of config.module.rules) {
      if (_rule.oneOf) {
        _rule.oneOf.unshift({
          test: /\.ya?ml$/,
          use: 'yaml-loader'
        });
        break;
      }
    }

    config.resolve.alias = {
      ...config.resolve.alias,
      '@': path.resolve(__dirname, 'src'),
    };

    return config;
  },

  devServer: function (configFunction) {
    i18nLocaleTool.autoSync();

    return function (proxy, allowedHost) {
      const config = configFunction(proxy, allowedHost);
      config.proxy = {
        '/answer': {
          target: 'http://10.0.10.98:2060',
          changeOrigin: true,
          secure: false,
        },
        '/installation': {
          target: 'http://10.0.10.98:2060',
          changeOrigin: true,
          secure: false,
        },
      };
      return config;
    };
  },
};

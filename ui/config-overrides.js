const path = require('path');

module.exports = {
  webpack: function (config, env) {
    if (env === 'production') {
      config.output.publicPath = process.env.REACT_APP_PUBLIC_PATH;
    }
    config.resolve.alias = {
      ...config.resolve.alias,
      '@': path.resolve(__dirname, 'src'),
    };

    return config;
  },

  devServer: function (configFunction) {
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

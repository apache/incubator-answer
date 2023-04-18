const {
  addWebpackModuleRule,
  addWebpackAlias,
} = require("customize-cra");

const path = require("path");
const i18nPath = path.resolve(__dirname, "../i18n");

module.exports = {
  webpack: function(config, env) {
    addWebpackAlias({
      "@": path.resolve(__dirname, "src"),
      "@i18n": i18nPath
    })(config);

    addWebpackModuleRule({
      test: /\.ya?ml$/,
      use: "yaml-loader"
    })(config);

    // add i18n dir to ModuleScopePlugin allowedPaths
    const moduleScopePlugin = config.resolve.plugins.find(_ => _.constructor.name === "ModuleScopePlugin");
    if (moduleScopePlugin) {
      moduleScopePlugin.allowedPaths.push(i18nPath);
    }

    return config;
  },
  devServer: function(configFunction) {
    return function(proxy, allowedHost) {
      const config = configFunction(proxy, allowedHost);
      config.proxy = [
        {
          context: ['/answer', '/installation'],
          target: process.env.REACT_APP_API_URL,
          changeOrigin: true,
          secure: false,
        },
        {
          context: ['/custom.css'],
          target: process.env.REACT_APP_API_URL,
        }
      ];
      return config;
    };
  }
};

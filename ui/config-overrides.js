const {
  addWebpackModuleRule,
  addWebpackAlias
} = require("customize-cra");

const path = require("path");
const i18nPath = path.resolve(__dirname, "../i18n");

module.exports = {
  webpack: function(config, env) {
    if (env === "production") {
      config.output.publicPath = process.env.REACT_APP_PUBLIC_PATH;
    }

    addWebpackAlias({
      ["@"]: path.resolve(__dirname, "src"),
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
      config.proxy = {
        "/answer": {
          target: "http://10.0.20.88:8080/",
          changeOrigin: true,
          secure: false
        },
        "/installation": {
          target: "http://10.0.20.88:8080/",
          changeOrigin: true,
          secure: false
        }
      };
      return config;
    };
  }
};

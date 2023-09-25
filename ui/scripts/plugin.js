const path = require('path');
const fs = require('fs');

const pluginPath = path.join(__dirname, '../src/plugins');
const pluginFolders = fs.readdirSync(pluginPath);

pluginFolders.forEach((folder) => {
  const pluginFolder = path.join(pluginPath, folder);
  const stat = fs.statSync(pluginFolder);
  if (stat.isDirectory() && folder !== 'builtin') {
    if (!fs.existsSync(path.join(pluginFolder, 'index.ts'))) {
      return;
    }

    // add plugin to package.json
    const packageJson = require(path.join(pluginFolder, 'package.json'));
    const packageName = packageJson.name;
    const packageJsonPath = path.join(__dirname, 'package.json');
    const packageJsonContent = require(packageJsonPath);
    packageJsonContent.dependencies[packageName] = 'workspace:*';

    fs.writeFileSync(
      packageJsonPath,
      JSON.stringify(packageJsonContent, null, 2),
    );
  }
});

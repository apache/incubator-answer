const path = require('path');
const fs = require('fs');
const humps = require('humps');

const template = `
package {{slugName}}

import "github.com/answerdev/answer/plugin"

type {{pluginName}} struct {
}

func init() {
	plugin.Register(&{{pluginName}}{})
}

func (d {{pluginName}}) Info() plugin.Info {
	return plugin.Info{
		Name:        plugin.MakeTranslator("i18n.{{slugName}}.name"),
		SlugName:    "{{slugName}}",
		Description: plugin.MakeTranslator("i18n.{{slugName}}.description"),
		Author:      "{{author}}",
		Version:     "{{version}}",
	}
}
`;

const pluginPath = path.join(__dirname, '../src/plugins');
const pluginFolders = fs.readdirSync(pluginPath);

pluginFolders.forEach((folder) => {
  const pluginFolder = path.join(pluginPath, folder);
  const stat = fs.statSync(pluginFolder);
  if (stat.isDirectory() && folder !== 'builtin') {
    if (!fs.existsSync(path.join(pluginFolder, 'index.ts'))) {
      return;
    }

    const tsFile = fs.readFileSync(
      path.join(pluginFolder, 'index.ts'),
      'utf-8',
    );
    const slugName = tsFile.match(/slug_name: '(.*)'/)[1];
    const pluginName = humps.pascalize(slugName) + 'Plugin';
    const packageJson = require(path.join(pluginFolder, 'package.json'));
    const author = packageJson.author;
    const version = packageJson.version;
    const content = template
      .replace(/{{slugName}}/g, slugName)
      .replace(/{{pluginName}}/g, pluginName)
      .replace(/{{author}}/g, author)
      .replace(/{{version}}/g, version);
    fs.writeFileSync(path.join(pluginFolder, `${slugName}.go`), content);
  }
});

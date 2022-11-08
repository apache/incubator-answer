/* eslint-disable import/no-extraneous-dependencies */
const path = require('node:path');
const fs = require('node:fs');

const chokidar = require('chokidar');

const SRC_PATH = path.resolve(__dirname, '../../i18n');
const DEST_PATH = path.resolve(__dirname, '../src/i18n/locales');
const PRESET_LANG = ['en_US', 'zh_CN'];

const cleanLocales = () => {
  fs.readdirSync(DEST_PATH).forEach((fp) => {
    fs.rmSync(path.resolve(DEST_PATH, fp), { force: true, recursive: true });
  });
};

const copyLocaleFile = (filePath) => {
  const targetFilePath = path.resolve(DEST_PATH, path.basename(filePath));
  fs.copyFile(
    filePath,
    targetFilePath,
    fs.constants.COPYFILE_FICLONE,
    (err) => {
      if (err) {
        throw err;
      }
    },
  );
};

const watchAndSync = () => {
  chokidar
    .watch(path.resolve(SRC_PATH, '*.yaml'), {
      awaitWriteFinish: true,
    })
    .on('all', (evt, filePath) => {
      copyLocaleFile(filePath);
    });
};

const autoSync = () => {
  cleanLocales();
  watchAndSync();
};

const resolvePresetLocales = () => {
  PRESET_LANG.forEach((lng) => {
    const sp = path.resolve(SRC_PATH, `${lng}.yaml`);
    copyLocaleFile(sp);
  });
};

module.exports = {
  autoSync,
  resolvePresetLocales,
};

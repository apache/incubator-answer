const fs = require('fs');
const path = require('path');
const yaml = require('js-yaml');

const configFilePath = path.resolve(__dirname, '../../configs/config.yaml');
const envFilePath = path.resolve(__dirname, '../.env.production');

// Read config.yaml file
const config = yaml.load(fs.readFileSync(configFilePath, 'utf8'));

// Generate .env file content
let envContent = 'TSC_COMPILE_ON_ERROR=true\nESLINT_NO_DEV_ERRORS=true\n';
for (const key in config.ui) {
  const value = config.ui[key];
  envContent += `${key.toUpperCase()}=${value}\n`;
}

// Write .env file
fs.writeFileSync(envFilePath, envContent);

console.log(`Generated ${envFilePath}`);

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

const path = require('path');
const fs = require('fs');

const pluginPath = path.join(__dirname, '../src/plugins');
const pluginFolders = fs.readdirSync(pluginPath);

function pascalize(str) {
  return str.replace(/\b\w/g, (match) => match.toUpperCase())
    .replace(/[-_\s]+/g, '');
}

function resetPackageJson() {
  const packageJsonPath = path.join(__dirname, '..', 'package.json');
  const packageJsonContent = require(packageJsonPath);
  const dependencies = packageJsonContent.dependencies;
  for (const key in dependencies) {
    if (dependencies[key].startsWith('workspace')) {
      delete dependencies[key];
    }
  }
  fs.writeFileSync(
    packageJsonPath,
    JSON.stringify(packageJsonContent, null, 2),
  );
}

function resetIndexTs() {
  const indexTsPath = path.join(pluginPath, 'index.ts');
  fs.writeFileSync(indexTsPath, '');
}

function addPluginToPackageJson(packageName) {
  const packageJsonPath = path.join(__dirname, '..', 'package.json');
  const packageJsonContent = require(packageJsonPath);
  packageJsonContent.dependencies[packageName] = 'workspace:*';

  fs.writeFileSync(
    packageJsonPath,
    JSON.stringify(packageJsonContent, null, 2),
  );
}

function addPluginToIndexTs(packageName) {
  const indexTsPath = path.join(pluginPath, 'index.ts');
  const indexTsContent = fs.readFileSync(indexTsPath, 'utf-8');
  const lines = indexTsContent.split('\n');
  const ComponentName = pascalize(packageName);
  const importLine = `export { default as ${ComponentName} } from '${packageName}';`;
  if (!lines.includes(importLine)) {
    lines.push(importLine);
  }
  fs.writeFileSync(indexTsPath, lines.join('\n'));
}

const pluginLength = pluginFolders.filter((folder) => {
  const pluginFolder = path.join(pluginPath, folder);
  const stat = fs.statSync(pluginFolder);
  return stat.isDirectory() && folder !== 'builtin';
}).length

if (pluginLength > 0) {
  resetIndexTs();
}

resetPackageJson();

pluginFolders.forEach((folder) => {
  const pluginFolder = path.join(pluginPath, folder);
  const stat = fs.statSync(pluginFolder);

  if (stat.isDirectory() && folder !== 'builtin') {
    if (!fs.existsSync(path.join(pluginFolder, 'index.ts'))) {
      return;
    }
    const packageJson = require(path.join(pluginFolder, 'package.json'));
    const packageName = packageJson.name;


    addPluginToPackageJson(packageName);
    addPluginToIndexTs(packageName);
  }
});

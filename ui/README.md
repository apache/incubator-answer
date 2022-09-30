# Answer

`Answer` is a modern Q&A community application ✨

To learn more about the philosophy and goals of the project, visit [Answer](https://answer.dev.segmentfault.com).

### 📦 Prerequisites

- [Node.js](https://nodejs.org/) `>=16.17`
- [pnpm](https://pnpm.io/) `>=7`

pnpm is required by building the Answer project. To installing the pnpm tools with below commands:

```bash
corepack enable
corepack prepare pnpm@v7.12.2 --activate
```

With Node.js v16.17 or newer, you may install the latest version of pnpm by just specifying the tag:

```bash
corepack prepare pnpm@latest --activate
```

## 🔨 Development

clone the repo locally and run following command in your terminal:

```shell
$ git clone `answer repo` answer
$ cd answer
$ pnpm install
$ pnpm run start
```

now, your browser should already open automatically, and autoload `http://localhost:3000`.
you can also manually visit it.

## 👷 Workflow

when cloning repo, and run `pnpm install` to init dependencies. you can use project commands below:

- `pnpm run start` run Answer web locally.
- `pnpm run build:dev` build code for environment `dev`
- `pnpm run build:test` build code for environment `test`
- `pnpm run build:prod` build code for environment `prod`
- `pnpm run lint` lint and fix the code style
- `pnpm run cz` run `git commit` by `commitizen`

## 🖥 Environment Support

| [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/edge/edge_48x48.png" alt="Edge" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)<br>Edge | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/firefox/firefox_48x48.png" alt="Firefox" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)<br>Firefox | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/chrome/chrome_48x48.png" alt="Chrome" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)<br>Chrome | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/safari/safari_48x48.png" alt="Safari" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)<br>Safari |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| last 2 versions                                                                                                                                                                                          | last 2 versions                                                                                                                                                                                                  | last 2 versions                                                                                                                                                                                              | last 2 versions                                                                                                                                                                                              |

## Build with

- [React.js](https://reactjs.org/) - Our front end is a React.js app.
- [Bootstrap](https://getbootstrap.com/) - UI library.
- [React Bootstrap](https://react-bootstrap.github.io/) - UI Library(rebuilt for React.)
- [axios](https://github.com/axios/axios) - Request library
- [SWR](https://swr.bootcss.com/) - Request library
- [react-i18next](https://react.i18next.com/) - International library
- [zustand](https://github.com/pmndrs/zustand) - State-management library

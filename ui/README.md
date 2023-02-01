# Answer

`Answer` is a modern Q&A community application ✨

To learn more about the philosophy and goals of the project, visit [Answer](https://answer.dev).

## ⚙️ Prerequisites

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
$ git clone git@github.com:answerdev/answer.git answer
$ cd answer/ui
$ pnpm install
$ pnpm start
```

now, your browser should already open automatically, and autoload `http://localhost:3000`.
you can also manually visit it.

## 👷 Workflow

when cloning repo, and run `pnpm install` to init dependencies. you can use project commands below:

- `pnpm start` run Answer web locally.
- `pnpm build` build Answer for production
- `pnpm lint` lint and fix the code style

## 🌍 I18n(Multi-language)
If you need to add or edit a language entry, just go to the `/i18n/en_US.yaml` file,
all front-end language entries are placed under the `ui` field.

If you would like to help us with the i18n translation, please visit [Answer@Crowdin](https://crowdin.com/translate/answer)

## 💡 Project instructions

```
.
├── cmd
├── configs
├── docs
├── i18n
      ├── en_US.yaml (basic language file)
      ├── i18n.yaml (language list)
├── internal
├── ...
└── ui (front-end project starts here)
      ├── build (built results directory, usually without concern)
      ├── public (html template for public)
      ├── scripts (some scripting tools on front-end project)
      ├── src (almost all front-end resources are here)
            ├── assets (static resources)
            ├── common (project information/data defined here)
            ├── components (all components of the project)
            ├── hooks (all hooks of the project)
            ├── i18n (Used only to initialize the front-end i18n tool)
            ├── pages (all pages of the project)
            ├── router (Project routing definition)
            ├── services (all data api of the project)
            ├── stores (all data stores of the project)
            ├── utils (all utils of the project)
```

## 🤝 Contributing

#### Fix Bug
If you find a bug, please don't hesitate to [submit an issue](https://github.com/answerdev/answer/issues) to us.
If you can fix it, please include a note with your issue submission.
If it is a bug definitely, you can submit your PR after we confirm it, which will ensure you don't do anything useless.

#### Code Review & Comment
In our development, some codes are not logical we know. If you find it, please don't hesitate to submit PR to us.
In the same way, some function has no comment. We would appreciate it if you could help us supplement it.

#### Translation
All our translations are placed in the i18n directory.

1. If you find that the corresponding key in the language you are using does not have a translation, you can submit your translation.
2. If you want to submit a new language translation, please add your language to the `i18n.yaml` file.

#### Features or Plugin
1. We developed the features for the plan based on the [roadmap](https://github.com/orgs/answerdev/projects/1). If you are suggestions for new functions, please confirm whether they have been planned.
2. Plugins will be available in the future, so stay tuned.

## 📱Environment Support

| [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/edge/edge_48x48.png" alt="Edge" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)<br />Edge | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/firefox/firefox_48x48.png" alt="Firefox" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)<br />Firefox | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/chrome/chrome_48x48.png" alt="Chrome" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)<br />Chrome | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/safari/safari_48x48.png" alt="Safari" width="24px" height="24px" />](http://godban.github.io/browsers-support-badges/)<br />Safari |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| last 2 versions                                                                                                                                                                                        | last 2 versions                                                                                                                                                                                                    | last 2 versions                                                                                                                                                                                                | last 2 versions                                                                                                                                                                                                |

## ⭐ Built with
- [TypeScript](https://www.typescriptlang.org/) - strongly typed JavaScript
- [React.js](https://reactjs.org/) - Our front end is a React.js app
- [React Router](https://reactrouter.com/en/main) - Router library
- [Bootstrap](https://getbootstrap.com/) - UI library.
- [React Bootstrap](https://react-bootstrap.github.io/) - UI Library(rebuilt for React)
- [axios](https://github.com/axios/axios) - Request library
- [SWR](https://swr.bootcss.com/) - Request library
- [react-i18next](https://react.i18next.com/) - International library
- [zustand](https://github.com/pmndrs/zustand) - State-management library

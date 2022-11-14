// There is a bug when using npm to install: the execution of preinstall is after install, so when this prompt is displayed, the dependent packages have already been installed.
if (!/pnpm/.test(process.env.npm_execpath)) {
    console.warn(`\u001b[33mThis repository requires using pnpm as the package manager for scripts to work properly.\u001b[39m\n`)
    process.exit(1)
}
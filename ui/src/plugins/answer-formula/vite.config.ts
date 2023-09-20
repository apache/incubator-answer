import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';
import cssInjectedByJsPlugin from 'vite-plugin-css-injected-by-js';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), cssInjectedByJsPlugin()],
  build: {
    lib: {
      entry: 'src/index.ts',
      name: 'answer-formula',
      fileName: (format) => `answer-formula.${format}.js`,
    },
    rollupOptions: {
      external: [
        'react',
        'react-dom',
        'react-i18next',
        'react-bootstrap',
        'katex',
      ],
      output: {
        globals: {
          react: 'React',
          'react-dom': 'ReactDOM',
          'react-i18next': 'reactI18next',
          'react-bootstrap': 'reactBootstrap',
          katex: 'katex',
        },
      },
    },
  },
});

import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    lib: {
      entry: 'src/index.ts',
      name: 'answer-chart',
      fileName: (format) => `answer-chart.${format}.js`,
    },
    rollupOptions: {
      external: [
        'react',
        'react-dom',
        'react-i18next',
        'react-bootstrap',
        'mermaid',
      ],
      output: {
        globals: {
          react: 'React',
          'react-dom': 'ReactDOM',
          'react-i18next': 'reactI18next',
          'react-bootstrap': 'reactBootstrap',
          mermaid: 'mermaid',
        },
      },
    },
  },
});

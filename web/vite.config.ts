import { defineConfig, esmExternalRequirePlugin } from 'vite';
import path from 'path';
import vue from '@vitejs/plugin-vue';
import Components from 'unplugin-vue-components/vite';

import { NaiveUiResolver } from 'unplugin-vue-components/resolvers';
// https://vitejs.dev/config/
export default defineConfig({
  server: {
    host: '0.0.0.0',
    // 优化开发服务器性能
    fs: {
      strict: false,
    },
    cors: true,
  },
  // 优化依赖预构建
  optimizeDeps: {
    include: [
      'vue',
      'pinia',
      'vue-router',
      'axios',
      'naive-ui',
      'lodash',
      'moment',
      'ethers',
    ],
    exclude: [],
  },
  plugins: [
    vue(),
    Components({
      resolvers: [NaiveUiResolver()],
    }),
    // esmExternalRequirePlugin({
    //   external: [/^node:/]
    // }),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
    },
  },
  build: {
    chunkSizeWarningLimit: 1000,
    rollupOptions: {
      output: {
        manualChunks: {
          'vue-vendor': ['vue', 'vue-router', 'pinia'],
          'naive-ui': ['naive-ui'],
          'utils': ['axios', 'lodash', 'moment'],
          'ethers': ['ethers'],
        },
      },
    },
  },
});

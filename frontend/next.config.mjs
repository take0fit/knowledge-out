// next.config.mjs
import path from 'path';
import { fileURLToPath } from 'url';

// ES Modulesで __dirname 相当の値を取得
const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default {
  // Base directory where the `pages` folder is located
  experimental: {
    appDir: path.join(__dirname, 'src'),
    sassOptions: {
      includePaths: [path.join(__dirname, 'src/styles')],
    },
  }
};

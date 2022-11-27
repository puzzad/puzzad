import {sveltekit} from '@sveltejs/kit/vite'
import type {UserConfig} from 'vite'

const config: UserConfig = {
  plugins: [
    sveltekit(),
  ],
  optimizeDeps: {
    force: true,
  },
  server: {
    host: "0.0.0.0"
  }
}

export default config

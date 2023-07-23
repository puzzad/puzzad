import {sveltekit} from '@sveltejs/kit/vite'
import type {UserConfig} from 'vite'
import { ValidateEnv } from '@julr/vite-plugin-validate-env';
import { z } from 'zod'

const config: UserConfig = {
  plugins: [
    sveltekit(),
    ValidateEnv({
      validator: 'zod',
      schema: {
        VITE_API_URL: z.string().transform((value) => value.endsWith('/') ? value : `${value}/`),
        VITE_HCAPTCHA_SITE_KEY: z.string().min(1)
      }
    })
  ],
  optimizeDeps: {
    force: true,
  },
  server: {
    host: "0.0.0.0"
  }
}

export default config
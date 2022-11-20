import sveltePreprocess from 'svelte-preprocess'
import adapter from '@sveltejs/adapter-node';

/** @type {import('@sveltejs/kit').Config} */
export default {
  preprocess: sveltePreprocess(),
  vitePlugin: {
    experimental: {
      useVitePreprocess: true,
    },
  },
  kit: {
    prerender: {
      enabled: false,
      entries: [],
    },
    alias: {
      $components: 'src/components',
      $assets: 'src/assets',
    },
    adapter: adapter({
      out: 'dist'
    }),
  },
}

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
    alias: {
      $components: 'src/components',
      $assets: 'src/assets',
    },
    adapter: adapter({
      out: 'dist'
    }),
    prerender: {
      entries: [],
    },
  },
}

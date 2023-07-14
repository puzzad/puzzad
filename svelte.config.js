import sveltePreprocess from 'svelte-preprocess'
import adapter from '@sveltejs/adapter-static';

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
      fallback: '200.html',
    }),
  },
}
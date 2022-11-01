import sveltePreprocess from "svelte-preprocess";

export default {
  preprocess: sveltePreprocess(),
  vitePlugin: {
    experimental: {
      useVitePreprocess: true
    }
  }
};

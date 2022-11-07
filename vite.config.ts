import { svelte } from '@sveltejs/vite-plugin-svelte'
import routify from '@roxi/routify/vite-plugin'
import { defineConfig } from 'vite'
import { resolve } from "path";
import preprocess from "svelte-preprocess";

const production = process.env.NODE_ENV === 'production'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        routify(),
        svelte({
            preprocess: [ preprocess()],
        }),
    ],
    resolve: {
        alias: {
            $lib: resolve(__dirname, "./src/lib"),
            $comps: resolve(__dirname, "./src/components"),
        },
    },
})

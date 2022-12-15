<script lang="ts">
  import NavBar from '$components/NavBar.svelte'
  import logo from '$assets/logo.png'
  import {FlatToast, ToastContainer} from 'svelte-toasts'
  import {title} from '$lib/title'
  import * as Sentry from '@sentry/svelte'
  import {BrowserTracing} from '@sentry/tracing'
  export const ssr = false;
  Sentry.init({
    dsn: import.meta.env.PROD ?
        'https://8d95b82da5804fdcaf8cbbeb7132edeb@glitch.puzzad.com/2' :
        'https://6c58e17321094c88bd40d834fd1f687e@glitch.puzzad.com/1',
    integrations: [new BrowserTracing()],
    tracesSampleRate: import.meta.env.PROD ? 0.01 : 1.0,
    environment: import.meta.env.PROD ? 'production' : 'development',
  })
</script>
<svelte:head>
  <title>{$title}</title>
</svelte:head>
<NavBar logo="{logo}"/>
<main>
<slot/>
</main>
  <ToastContainer placement="bottom-right" theme="dark" let:data={data}>
  <FlatToast {data}/>
</ToastContainer>
<footer>
  <p>&copy; 2022 Puzzad.com. All rights reserved. <a href="/about">About &amp; Privacy</a></p>
</footer>
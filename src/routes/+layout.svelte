<script lang="ts">
  import '../../src/style/global.scss'
  import {ToastContainer} from 'svelte-toasts'
  import {title} from '$lib/title'
  import * as Sentry from '@sentry/svelte'
  import {BrowserTracing} from '@sentry/tracing'
  import {logout, isLoggedIn, isAdmin} from '$lib/auth'
  import Nav from '$components/Nav.svelte'
  import Footer from '$components/Footer.svelte'
  import Toast from '$components/Toast.svelte'

  export const ssr = false

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
<Nav loggedIn={$isLoggedIn} on:logout={logout} />
<main>
  <slot/>
</main>
<ToastContainer placement="bottom-right" theme="dark" let:data={data}>
  <Toast type={data.type} title={data.title} message={data.description}/>
</ToastContainer>

<Footer isAdmin={$isAdmin}/>
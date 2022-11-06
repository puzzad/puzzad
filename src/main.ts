import App from "./App.svelte";
import * as Sentry from "@sentry/svelte";
import {BrowserTracing} from "@sentry/tracing";

Sentry.init({
    dsn: import.meta.env.PROD ?
        "https://8d95b82da5804fdcaf8cbbeb7132edeb@glitch.puzzad.com/2" :
        "https://6c58e17321094c88bd40d834fd1f687e@glitch.puzzad.com/1",
    integrations: [new BrowserTracing()],
    tracesSampleRate: import.meta.env.PROD ? 0.01 : 1.0,
    environment: import.meta.env.PROD ? 'production' : 'development',
});

const app = new App({
    target: document.getElementById("app"),
});

export default app;

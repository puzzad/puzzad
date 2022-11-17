<script lang="ts">
  import type { ActionData } from './$types';
  import { enhance } from '$app/forms';
  export let form: ActionData;
</script>
<style>
  form {
    display: grid;
  }
  input {
    margin-bottom: 1em;
  }
  button {
    margin-top: 1em;
    margin-right: 0;
  }
</style>
<h2>Contact us</h2>
<form method="POST" use:enhance>
  <label for="name">Name</label>
  {#if form?.nameMissing}<p class="error">The name field is required</p>{/if}
  <input disabled id="name" name="name" value="{form?.name ?? ''}" />
  <label for="email">Email</label>
  {#if form?.emailMissing}<p class="error">The email field is required</p>{/if}
  <input disabled type="email" id="email" name="email" value="{form?.email ?? ''}"/>
  <label for="message">Query</label>
  {#if form?.messageMissing}<p class="error">The query field is required</p>{/if}
  <textarea disabled id="message" name="message">{form?.message ?? ''}</textarea>
  <button type="submit" disabled>Submit</button>
</form>
{#if form?.success}
 <p>Your message has been sent</p>
{/if}
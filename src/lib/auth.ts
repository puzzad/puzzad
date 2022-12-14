import {readable} from 'svelte/store'
import type {Subscriber} from 'svelte/store'
import {supabase} from '$lib/db'
import {goto} from '$app/navigation'

export const isLoggedIn = readable<bool | null>(null, (set: Subscriber<bool | null>) => {
  supabase.auth.getSession().then(response => {
    set(response.data.session !== null)
  })
  const auth = supabase.auth.onAuthStateChange(async ({}, session) => {
    set(session !== null)
  })
  return auth.data.subscription.unsubscribe
})

export const logout = async () => {
  let {error} = await supabase.auth.signOut()
  if (!error) {
    await goto('/')
  }
}

export const isAdmin = readable<bool | null>(null, (set: Subscriber<bool | null>) => {
  supabase.auth.getSession().then(response => {
    set(response.data.session !== null && response.data.session.user.email?.endsWith(import.meta.env.VITE_ADMIN_DOMAIN))
  })
  const auth = supabase.auth.onAuthStateChange(async ({}, session) => {
    set(session !== null && session.user.email?.endsWith(import.meta.env.VITE_ADMIN_DOMAIN))
  })
  return auth.data.subscription.unsubscribe
})
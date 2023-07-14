import {readable, writable} from 'svelte/store'
import type {Subscriber} from 'svelte/store'
import PocketBase, {Admin, LocalAuthStore, Record} from 'pocketbase'

export const client = new PocketBase(import.meta.env.VITE_SUPABASE_URL, new LocalAuthStore('puzzad-user-key'))
client.autoCancellation(false)

const gameClient = new PocketBase(import.meta.env.VITE_SUPABASE_URL, new LocalAuthStore('puzzad-game-key'))
gameClient.autoCancellation(false)

export const currentUser = writable<Record | Admin | null>(client.authStore.model)

client.authStore.onChange((token, model) => {
  currentUser.set(model)
}, true)

export const getGameClient = async (code: string) => {
  if (gameClient.authStore.model && gameClient.authStore.model.username === code) {
    return Promise.resolve(gameClient)
  } else {
    return gameClient.collection('games').
        authWithPassword(code, 'puzzad').
        then(_ => gameClient)
  }
}

export const loginNative = async (email: string, password: string) => {
  return client.collection('users').authWithPassword(email, password)
}

export const loginOauth = async (provider: string) => {
  return client.collection('users').authWithOAuth2({provider: provider})
}

export const requestEmailVerification = async (email: string) => {
  return client.collection('users').requestVerification(email)
}

export const logout = async () => {
  client.authStore.clear()
}

export const isAdmin = readable<boolean | null>(false, (set: Subscriber<boolean | null>) => {
  return set(client.authStore.model instanceof Admin)
})
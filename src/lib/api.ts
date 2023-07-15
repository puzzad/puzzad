import {writable} from 'svelte/store'
import PocketBase, {Admin, LocalAuthStore, Record} from 'pocketbase'
import {goto} from '$app/navigation'

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
        then(_ => gameClient).
        then(c => { registerGameToAccount(c.authStore.model?.id); return c })
  }
}

const registerGameToAccount = (id: string | undefined) => {
  if (client.authStore.model !== null && id && !client.authStore.model.games.includes(id)) {
    client.collection('users').
        update(client.authStore.model.id, {games: client.authStore.model.games.concat([id])}).
        then(r => {})
  }
}

export const logout = async () => {
  client.authStore.clear()
  await goto("/")
}
import {get, readable, writable} from 'svelte/store'
import type {Subscriber} from 'svelte/store'
import PocketBase, {Admin, Record} from 'pocketbase'

export const pb = new PocketBase(import.meta.env.VITE_SUPABASE_URL);
pb.autoCancellation(false);

export const currentUser = writable<Record | Admin | null>(pb.authStore.model);

pb.authStore.onChange((token, model) => {
  currentUser.set(model);
}, true);

export const loginNative = async (email, password) => {
  return pb.collection("users").authWithPassword(email, password)
}

export const loginOauth = async (provider) => {
  return pb.collection("users").authWithOAuth2({provider: provider})
}

export const requestEmailVerification = async (email) => {
  return pb.collection("users").requestVerification(email)
}

export const logout = async () => {
  pb.authStore.clear()
}

export const signup = async(email, password) => {
  pb.collection("users").create({
    email: email,
    password: password,
    passwordConfirm: password,
  })
      .then(_ => login(email,password))
      .catch(error => console.log(error))
}

export const isAdmin = readable<bool | null>(false, (set: Subscriber<bool | null>) => {
  return set(pb.authStore.model instanceof Admin)
})
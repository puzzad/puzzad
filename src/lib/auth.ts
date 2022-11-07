import {readable} from 'svelte/store';
import type {Session, User} from "@supabase/supabase-js";
import {supabase} from "$lib/db.ts";
import {replace} from "svelte-spa-router";

export const session = readable(null, (set: Subscriber<Session | null>) => {
    supabase.auth.getSession().then(response => {
        set(response.data.session)
    })
    const auth = supabase.auth.onAuthStateChange(async ({}, session) => {
        set(session)
    });
    return auth.data.subscription.unsubscribe;
})
export const logout = async () => {
    let {error} = await supabase.auth.signOut()
    if (!error) {
        await replace("/");
    }
}
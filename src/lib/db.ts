import {createClient, RealtimeClient, SupabaseClient} from "@supabase/supabase-js";
import type {Database} from "./schema";
import {decodeJWTPayload} from "@supabase/gotrue-js/src/lib/helpers";

export const supabase = createClient<Database>(
    import.meta.env.VITE_SUPABASE_URL,
    import.meta.env.VITE_SUPABASE_ANON_KEY
);

export async function getGameClient(gameCode: string): Promise<SupabaseClient<Database>> {
    console.log(gameCode)
    return createClient<Database>(
        import.meta.env.VITE_SUPABASE_URL,
        import.meta.env.VITE_SUPABASE_ANON_KEY,
        {
            global: {
                headers: {
                    Authorization: `Bearer ${await getGameToken(gameCode)}`
                }
            }
        }
    )
}

export async function getRealTimeClient(gameCode: string) {
    const client = new RealtimeClient(
        (import.meta.env.VITE_SUPABASE_URL + `/realtime/v1`).replace(/^http/i, 'ws'),
        {
            params: {
                // Kong requires us to pass a known key to authenticate
                apikey: import.meta.env.VITE_SUPABASE_ANON_KEY
            },
        }
    )
    // This is the actual key we'll be using for dealing with RLS etc
    client.setAuth(await getGameToken(gameCode))
    return client
}

async function getGameToken(gameCode: string): Promise<string> {
    const existingKey = localStorage.getItem('puzzad-game-token')
    if (isValidToken(existingKey, gameCode)) {
        return existingKey
    }

    return await supabase.rpc('getteamcodejwt', {teamcode: gameCode})
        .then(({data: code, error}) => {
            if (error) {
                throw error
            } else {
                return code.toString()
            }
        })
        .then((code) => {
            localStorage.setItem('puzzad-game-token', code)
            return code
        })
}

function isValidToken(token: string, code: string): boolean {
    if (!token) {
        return false
    }

    const payload = decodeJWTPayload(token)
    return payload.exp > (Date.now() / 1000 + 60 * 60) && payload.code == code
}
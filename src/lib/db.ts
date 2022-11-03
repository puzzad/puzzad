import {createClient, SupabaseClient} from "@supabase/supabase-js";
import type {Database} from "./schema";
import {decodeJWTPayload} from "@supabase/gotrue-js/src/lib/helpers";

export const supabase = createClient<Database>(
    import.meta.env.VITE_SUPABASE_URL,
    import.meta.env.VITE_SUPABASE_ANON_KEY
);

export function getGameClient(gameCode: string): PromiseLike<SupabaseClient> {
    const existingKey = localStorage.getItem('puzzad-game-token')
    const tokenPromise = (existingKey && isValidToken(existingKey)) ?
        Promise.resolve(existingKey) :
        supabase.rpc('getteamcodejwt', {teamcode: gameCode})
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

    return tokenPromise.then((token) => {
            const client = createClient<Database>(
                import.meta.env.VITE_SUPABASE_URL,
                import.meta.env.VITE_SUPABASE_ANON_KEY,
                {
                    global: {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    }
                }
            )
            // @ts-ignore - realtime isn't exposed, and doesn't use headers for auth :|
            client.realtime.setAuth(token)
            return client
        }
    )
}

function isValidToken(token: string): boolean {
    if (!token) {
        return false
    }

    const payload = decodeJWTPayload(token)
    return payload.exp > (Date.now() / 1000 + 60 * 60)
}
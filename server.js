import express from 'express'
import bodyParser from 'body-parser'
import jwt from 'jsonwebtoken'
import {config} from 'dotenv'
import {createClient} from '@supabase/supabase-js'

config()
// Should have JWT_SECRET and SERVICE_ROLE from supabase studio
config({path: '.env.local'})
const port = process.env.PORT || 3000
const app = express()
const supabase = createClient(process.env.VITE_SUPABASE_URL, process.env.SERVICE_ROLE)

app.disable('x-powered-by')

app.use(bodyParser.urlencoded({extended: false}))
app.use(express.static('dist'))

app.post('/teamcodelogin', (req, res) => {
    checkCode(req.body.code)
        .then(result => {
            createJWT(result)
                .then(resp => res.send(`${resp}`))
                .catch(error => res.send(`${error}`))
        })
        .catch(error => res.send(`${error}\n`))
})
app.listen(port, "0.0.0.0", () => {
    console.log(`Webserver running: http://0.0.0.0:${port}`)
})

async function checkCode(code) {
    let {data: resp, error} = await supabase
        .from('games')
        .select('id, code')
        .eq('code', code)
    if (error) {
        return Promise.reject(new Error('Error checking code'))
    }
    if (resp.length === 1) {
        return code
    }
    return Promise.reject(new Error('Invalid login code'))
}

async function createJWT(code) {
    let date = Date.now()
    let expires = new Date(date)
    expires.setDate(expires.getDate() + 1)
    return jwt.sign(
        {
            'iss'    : 'supabase',
            'ref'    : process.env.SUPBASE_ID,
            'role'   : 'anon',
            'iat'    : date,
            'expires': expires.getTime(),
            'code'   : code,
        },
        process.env.JWT_SECRET,
        {
            noTimestamp: true,
        },
        null,
    )
}
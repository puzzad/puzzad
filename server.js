import express from 'express'
import bodyParser from 'body-parser'

const app = express()
const port = process.env.PORT || 3000;

app.use(bodyParser.urlencoded({ extended: false }))
app.use(express.static('dist'))
app.disable('x-powered-by')
app.post('/teamcodelogin', (req, res) => {
    let code = req.body.code
    //Do something with the code
})
app.listen(port, () => {
    console.log(`Webserver running: http://0.0.0.0:${port}`)
})
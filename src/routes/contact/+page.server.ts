import Mailgun from 'mailgun.js'
import {MAILGUN_API_KEY} from '$env/static/private'
import formData from 'form-data'
import type { Actions } from './$types';
import {invalid} from '@sveltejs/kit'

export const actions: Actions = {
  default: async (event) => {
    const requestData = await event.request.formData();
    const name = requestData.get('name')
    const email = requestData.get('email')
    const message = requestData.get('message')
    let error = false
    let data = {}
    if (!name) {
      data.nameMissing = true
      error = true
    }
    if (!email || !/[a-z0-9\._%+!$&*=^|~#%'`?{}/\-]+@([a-z0-9\-]+\.){1,}([a-z]{2,16})/.test(email)) {
      data.emailMissing = true
      error = true
    }
    if (!message) {
      error = true
      data.messageMissing = true
    }
    data.name = name
    data.email = email
    data.message = message
    if (error) {
      return invalid(400, data)
    }
    // const mailgun = new Mailgun(formData)
    // const mg = mailgun.client({username: 'api', key: MAILGUN_API_KEY, url: 'https://api.eu.mailgun.net'})
    // mg.messages.create('notify.puzzad.com', {
    //   from: 'Puzzad Contact Form <postmaster@notify.puzzad.com>',
    //   to: ['greg@puzzad.com'],
    //   subject: 'Contact Form',
    //   text: "Name: "+name+"\nEmail: "+email+"\nMessage: \n"+message,
    //   html: "Name: "+name+"<br>Email: "+email+"<br>Message: <br>"+message,
    // }).catch(err => {
    //   return { success: false }
    // })
    return { success: true }
  }
};
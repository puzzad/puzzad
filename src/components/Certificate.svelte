<script lang="ts">
  import certificateFrame from '$assets/certificate-frame.webp'
  import Spinner from './Spinner.svelte'
  import Error from './Error.svelte'

  export let teamName = ''
  export let adventureName = ''
  export let completionDate = ''

  let data = new Promise((resolve, reject) => {
    const image = new Image()
    image.addEventListener('load', () => resolve(image))
    image.addEventListener('error', (e) => reject(e))
    image.src = certificateFrame
  }).then((image: HTMLImageElement) => {
    const canvas = document.createElement('canvas')
    canvas.width = image.width
    canvas.height = image.height
    const ctx = canvas.getContext('2d')
    ctx.drawImage(image, 0, 0)
    ctx.font = '28px serif'
    ctx.textAlign = 'center'
    ctx.fillText(teamName, 944, 430, 420)
    ctx.fillText(adventureName, 944, 595, 420)
    ctx.fillText((new Date(completionDate)).toLocaleDateString('en-gb',
        {weekday: 'long', year: 'numeric', month: 'long', day: 'numeric'}), 415, 805, 420)
    return canvas.toDataURL('image/webp')
  })
</script>

{#await data}
  <Spinner></Spinner>
{:then data}
  <img src={data}>
{:catch error}
  <Error error={error}></Error>
{/await}
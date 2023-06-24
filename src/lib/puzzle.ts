import {getGameClient} from '$lib/db'

const embedRegex = /\$[^$]+?\$/g
const sectionRegex = /<section class="(.*?)">.*?<\/section>/gs

export async function parsePuzzleContent(gameCode: string, storageSlug: string, content: string) {
  return Promise.all([
    getGameClient(gameCode),
    storageSlug,
    content,
  ]).then(([gameClient, storageSlug, content]) => {
    const fileNames = content.match(embedRegex) || []
    return Promise.all([
      content,
      ...fileNames.map((f) => f.substring(1, f.length - 1)).
          map((f) => gameClient.storage.from('puzzles').
              createSignedUrl(`${storageSlug}/${f}`, 60 * 60).
              then(({data, error}) => {
                if (error) {
                  throw error
                }
                return data.signedUrl
              }),
          ),
    ])
  }).then(([content, ...urls]) =>
      content.replace(embedRegex, () => urls.shift() || ''),
  ).then(content => {
    let sections: { [key: string]: string; } = {}
    for (const [section, type] of content.matchAll(sectionRegex)) {
      sections[type] = (sections[type] || '') + section
    }
    return sections
  })
}
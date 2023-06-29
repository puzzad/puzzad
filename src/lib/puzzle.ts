import {getGameClient} from '$lib/db'

const embedRegex = /\$[^$]+?\$/g
const sectionRegex = /<section class="(.*?)">.*?<\/section>/gs

export async function parsePuzzleContent(gameCode: string, storageSlug: string, content: string) {
  return Promise.all([
    content,
  ]).then(content => {
    let sections: { [key: string]: string; } = {}
    for (const [section, type] of content.matchAll(sectionRegex)) {
      sections[type] = (sections[type] || '') + section
    }
    return sections
  })
}
export const formatDuration = function(start: string, end: string) {
  const diff = ((new Date(end)).getTime() - (new Date(start)).getTime()) / 1000
  const text = (value: number, label: string) => {
    const rounded = Math.floor(value)
    if (rounded === 1) {
      return `${rounded} ${label}`
    } else if (rounded > 1) {
      return `${rounded} ${label}s`
    } else {
      return ''
    }
  }

  return [
    text(diff % 60, 'second'),
    text((diff / 60) % 60, 'minute'),
    text((diff / 3600) % 24, 'hour'),
    text(diff / 86400, 'day'),
  ].filter((t) => t !== '').reverse().join(', ')
}
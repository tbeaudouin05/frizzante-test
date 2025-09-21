import type { Config } from 'tailwindcss'

export default {
  content: [
    './lib/**/*.{html,js,svelte,ts}',
    './app.client.svelte',
    './index.html'
  ]
} satisfies Config

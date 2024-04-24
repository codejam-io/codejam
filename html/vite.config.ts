import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/

// We're running HMR on a separate port so that when we're in dev mode, and proxying through the Go backend, the UI
// can still connect to HMR on its own port instead of us needing to proxy or redirect (which doesn't seem to work)
export default defineConfig({
	plugins: [svelte()],
	server: {
		hmr: {
			host: 'localhost',
			port: 5174,
		},
	},
})

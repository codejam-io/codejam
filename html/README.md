# CodeJam Website

Framework: [SvelteKit](https://kit.svelte.dev/docs/introduction)

Component Library: [flowbite-svelte](https://flowbite-svelte.com/)

Icon Library: [FontAwesome](https://fontawesome.com/icons)

## Initial Setup

The first time using this repo, run 
```bash
npm install
```

To run the site locally:
```bash
npm run dev
```
or
```
# to start the server and open the app in a new browser tab
npm run dev -- --open
```


## Dev Info

### Site Rendering
The site will be built as an SPA, by way of the `prerender` flag being set to true in `routes/+layout.js`, and
using [`adapter-static`](https://kit.svelte.dev/docs/adapter-static)


### Routing & Page Content
The routing mostly follows the standard [SvelteKit concept](https://kit.svelte.dev/docs/routing) with the exception 
that the actual page content is stored in a descriptive Svelte component next to the +page.svelte file.

For example, the home page for the `/` route is named HomePage.svelte, and the +page.svelte will simply import that 
component and use it.

The reason for this is personal preference, to be able to quickly open files by name.  I don't like having a bunch
of frequently edited +page.svelte files where the only differentiator is the directory they reside in.


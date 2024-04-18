import { writable } from 'svelte/store';

export const activeContent = writable('');
export const loggedIn = writable(false);
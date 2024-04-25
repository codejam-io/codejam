
import {writable, derived} from 'svelte/store';

export const activeContent = writable('');
export const userStore = writable(null);
export const loggedInStore = derived(userStore, (userData) => userData != null);
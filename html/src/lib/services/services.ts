import {readable} from "svelte/store";
import {eventStore, userStore} from "../stores/stores";

// This shouldn't ever need to be set since dev and prod environments will just use relative endpoints
export let baseApiUrl : string = "";

const originalFetch = window.fetch;

// Override standard "fetch" so we have a place to generically handle errors
window.fetch = (...args) => {
    return new Promise((resolve) => {
        originalFetch(...args)
            .then((response) => {
                resolve(response);
            })
            .catch((err) => {
                console.error("fetch error:", err, ...args);
                throw err;
            })
    });
}

export async function getUser() {
    return fetch(baseApiUrl + "/user/", {method: 'GET', credentials: 'include'})
        .then((response) => {
            if (response.status === 401) {
                userStore.set(null);
            } else {
                response.json()
                    .then((data) => {
                        userStore.set(data);
                    })
                    .catch((err) => {
                        console.error("error deserializing user", err);
                    });
            }
        });
}

export async function logout() {
    return fetch(baseApiUrl + "/user/logout")
        .then((response) => {
            userStore.set(null);
        })
        .catch((err) => {
            console.error("Logout error", err);
        });
}

export async function getEvent(id: string) {
    return fetch(baseApiUrl + "/event/" + id)
        .then((response) => {
            if (response.status === 401) {
                userStore.set(null);
            } else {
                response.json()
                    .then((data) => {
                        if (Array.isArray(data)) {
                            eventStore.set(data[0]);
                        } else {
                            eventStore.set(data);
                        }
                    })
                    .catch((err) => {
                        console.error("error deserializing event", response, err);
                    });
            }
        });
}

// Always call at startup to get the initial states
async function initialLoad() {
    getUser();
    getEvent(""); // loads all events
}

initialLoad();
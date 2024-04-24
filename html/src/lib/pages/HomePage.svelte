<script lang="ts">
    import {Button} from "flowbite-svelte";
    import DiscordIcon from "../components/DiscordIcon.svelte";
    import {baseApiUrl, logout} from "../services/services";
    import {userStore} from "../stores/stores";

    let loggedIn : boolean = false;

    userStore.subscribe((value) => {
        loggedIn = value != null;
    });

    function login() {
        window.location = baseApiUrl +  '/oauth/redirect';
    }

</script>


<h1>Home Page</h1>

{#if !loggedIn}
    <Button on:click={login}>
        <span class="mr-3">Login with Discord</span> <DiscordIcon size="2x"/>
    </Button>
{:else}
    <Button on:click={logout}>
        Logout
    </Button>
{/if}

<p>Logged In: {loggedIn}</p>


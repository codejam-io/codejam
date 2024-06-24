<script lang="ts">

    import {Avatar} from "flowbite-svelte";
    import type {User} from "../models/user";

    export let user : User | null;

</script>


{#if user && user.AvatarUrl}
    <Avatar border src="https://cdn.discordapp.com/avatars/{user.ServiceUserId}/{user.AvatarUrl}.png?size=48" {...$$restProps}></Avatar>
{:else}
    <Avatar border {...$$restProps}></Avatar>
{/if}
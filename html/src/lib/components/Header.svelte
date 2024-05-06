<script lang="ts">

import Banner from "./Banner.svelte";
import {Navbar, NavHamburger, NavLi, NavUl} from "flowbite-svelte";
import {loggedInStore} from "../stores/stores";
import DiscordIcon from "./DiscordIcon.svelte";


</script>


<div>
    <Banner size="normal" />
    <div class="absolute right-0 top-0">
        <Navbar class="rounded-bl-3xl">
            <NavHamburger/>
            <NavUl>
                <NavLi href="/">Home</NavLi>
                {#if $loggedInStore}
                    <NavLi href="/">My Team</NavLi>
                    <NavLi href="/">Profile</NavLi>
                    <NavLi href="/user/logout">Logout</NavLi>
                {:else}
                    <NavLi href="/oauth/redirect">Login with Discord <DiscordIcon/></NavLi>
                {/if}
            </NavUl>
        </Navbar>

    </div>
</div>
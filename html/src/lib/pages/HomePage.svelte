<script lang="ts">
    import '../../styles/styles.css';
	import Nav from '../components/Nav.svelte';
	import Banner from '../components/Banner.svelte';
	//import GradientChange from '../components/GradientChange.svelte';
	import Card from '../components/Card.svelte';
	import { activeContent } from '../stores/stores';
	import Panel from '../components/Panel.svelte';
	import Content from '../components/Content.svelte';
    import {Button} from "flowbite-svelte";
    import DiscordIcon from "../components/DiscordIcon.svelte";
    import {baseApiUrl, logout} from "../services/services";
    import {userStore} from "../stores/stores";
	import { faArrowUpRightFromSquare } from '@fortawesome/free-solid-svg-icons';
    
    let loggedIn : boolean = false;
    
    userStore.subscribe((value) => {
        // if value is not null, set loggedIn to True
        // if value is null (default), set loggedIn to False
        loggedIn = value != null
        }
    )

    function login() {
        if (typeof window !== 'undefined') {
            window.location.assign(baseApiUrl +  '/oauth/redirect');
        }
    }
</script>


<div class="flex col">
	<!-- <GradientChange /> -->

	<!-- <Card /> -->

	<!-- <div class="content">
        {#if $activeContent === 'goals'}
        <div class="card">{@html goals}</div>
        {:else if $activeContent === 'timeline'}
        <div class="card">{@html timeline}</div>
        {/if}
    </div> -->

	<div class="flex row">
		<div id="left-side" class="flex col">
			<Banner size="normal" />

            {#if !loggedIn}
            <Button on:click={login}>
                <span class="mr-3">LOGIN WITH DISCORD!!!!!!!!!!!!!!!!!!!!!!!!!</span> <DiscordIcon size="2x"/>
            </Button>
            {:else}
                {$userStore?.DisplayName}
                
                <Button on:click={logout}>
                    Logout
                </Button>
            {/if}

			
			<!-- {#if !$loggedIn}
            <!-- dont need big aesthetic when already logged in -->
				<Panel />
			<!-- {/if} --> 
			<Nav />
			<Content />
		</div>
		<div id="right-side" class="flex col">
			<Card />
			<Card />
		</div>
	</div>
</div>

<style>
	#right-side {
		max-width: 50%;
	}
</style>

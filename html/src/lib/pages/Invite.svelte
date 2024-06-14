<script lang="ts">
	import Page from '../components/Page.svelte';
	import { Button, Card } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import CodeJamTeam from '../models/team';
    import { type User } from '../models/user'
	import { getTeamByInvite, joinTeam } from '../services/services';
	import { Label, Input } from 'flowbite-svelte';
	import type TeamMember from '../models/TeamMember';
	import CodeJamEvent from '../models/event';
	import { loggedInStore, userStore } from '../stores/stores';
	import DiscordIcon from '../components/DiscordIcon.svelte';

	export let params: any; // set by svelte-spa-router
    console.log(params) // returns Object { invitecode: "d1869a59b4fdf3" }
    console.log(params.invitecode)

	let teamData: CodeJamTeam | null = null;
	let teamMembers: TeamMember[] = [];
	let teamEvent: CodeJamEvent | null = null;
	let loading = true;
	let error: any = null;

	async function loadData(invitecode: string) {
		try {
			const response = await getTeamByInvite(invitecode);
			const data = await response.json();
			teamData = data.Team;
			teamMembers = data.Members;
			teamEvent = data.Event;
		} catch (err) {
			error = 'Failed to load team data.';
			console.error(err);
		} finally {
			loading = false;
		}
	}

	$: if (params) {
		loadData(params.invitecode);
	}

    // query teams table for matching invite_code. 
    // getTeamByInvite()
    // return table id
    
    // joinTeam()

    // checks if user is logged in. 
    // if not logged in, display: login to join team
    // if logged in: button to fetch


</script>

<Page>
	<Card>
		<h3>Join</h3>
		{#if $loggedInStore}
        <div class="py-4">
                <div>Hi {$userStore?.DisplayName},</div> 
                Click below to join {teamMembers[0]?.DisplayName}'s team: 
            </div>

            <Button>Join {teamData?.Name}</Button>
		{:else}
        <div class="py-4">
			Must be logged in to join a team.
        </div>
            <Button>
                <a href="/oauth/redirect">Login with Discord <DiscordIcon /></a> 
            </Button>
        
            
		{/if}


	</Card>
</Page>

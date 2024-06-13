<script lang="ts">
	import Page from '../components/Page.svelte';
	import { Button, Card } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import CodeJamTeam from '../models/team';
	import { getTeam, joinTeam } from '../services/services';
	import { Label, Input } from 'flowbite-svelte';
	import type TeamMember from '../models/TeamMember';
	import CodeJamEvent from '../models/event';
	import { loggedInStore, userStore } from '../stores/stores';

	export let params: any; // set by svelte-spa-router

	let teamData: CodeJamTeam | null = null;
	let teamMembers: TeamMember[] = [];
	let teamEvent: CodeJamEvent | null = null;
	let loading = true;
	let error: any = null;

	async function loadData(id: string) {
		try {
			const response = await getTeam(id);
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
		loadData(params);
	}

    


</script>

<Page>WAHT
    <Card>
        Click below to join {teamMembers[0]?.DisplayName}'s team: {teamData?.Name}

        <Button>Join Team</Button>
    </Card>
</Page>
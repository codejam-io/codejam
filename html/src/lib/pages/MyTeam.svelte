<script lang="ts">
	import Page from '../components/Page.svelte';
	import { Breadcrumb, BreadcrumbItem, Button, Card } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import CodeJamTeam from '../models/team';
	import { getTeamById } from '../services/services';
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
			const response = await getTeamById(id);
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
		loadData(params.id);
	}
</script>

<Page>
    <Breadcrumb solid class="mb-4 w-full max-w-screen-xl">
		<BreadcrumbItem href="/#/" home>Home</BreadcrumbItem>
		<BreadcrumbItem href="/#/team">Team Options</BreadcrumbItem>
		<BreadcrumbItem href="/#/team/create">Create Team</BreadcrumbItem>
        <BreadcrumbItem>View Teams</BreadcrumbItem>
	</Breadcrumb>

	<Card size="md" class="w-full flex">
		{#if loading}
			<div class="p-4">Loading...</div>
		{:else if error}
			<div class="p-4 text-red-500">{error}</div>
		{:else if teamData !== null}
			<div class="p-4">{teamMembers[0]?.DisplayName}, your team has been successfully created!</div>
			<h3 class="p-4">{teamEvent?.Title}</h3>
			<Card size="xl" class="flex w-full p-8 px-4 py-6 space-y-3">
				<center class="p-2">
					<h4>Team {teamData.Name}</h4>
					<span><small>(edit)</small></span>
				</center>
				<span>
					<b>Team Members: </b>
					<ul>
						{#each teamMembers as member}
							<li>{member.DisplayName}</li>
						{/each}
					</ul>
				</span>
				<span>
					<b>Visibility: </b>{teamData.Visibility}
				</span>
				<span>
					<b>Technologies: </b>{teamData.Technologies}
				</span>
				<span>
					<b>Availability: </b>{teamData.Availability}
				</span>
				<span>
					<b>Description: </b>{teamData.Description}
				</span>
				<p>Invite Link: <a href="/#/team/invite/{teamData.InviteCode}">here </a></p>
			</Card>
		{/if}
	</Card>
</Page>

<script lang="ts">

import Page from "../components/Page.svelte";
import {Button, Card} from "flowbite-svelte";
import {getTeams} from "../services/services";
import CodeJamTeam from '../models/team';
import type TeamMember from '../models/TeamMember';
export const params: Record<string, never> = {};

let teamData: CodeJamTeam | null = null;
let teamMembers: TeamMember[] = [];
let loading: boolean = true;
let error: string | null = null;
let allTeams: CodeJamTeam[]

async function loadData() {
    try {
        const response = await getTeams();
        allTeams = await response.json();  // Array of teams...
        console.log(allTeams)

    } catch (err) {
        error = `Failed to load team data: ${err}`;
        console.error(error)
    } finally {
        loading = false;
    }
}

loadData();

</script>
    
<Page>
    <Card size="md" class="w-full flex">
        <h3>Browse All Teams</h3>
        
		{#if loading}
			<div class="p-4">Loading...</div>
		{:else if error}
			<div class="p-4 text-red-500">{error}</div>
        
		{:else if allTeams === null}
            <div>
                Looks like you don't have any teams! Go to <a href="/#/">browse teams</a> to join one!
            </div>
        {:else if allTeams.length === 0}
            <div>
                Error, please contact admin.
            </div>
        {:else}

            {#each allTeams as Team}
            <Card size="xl" class="flex w-full p-8 px-4 py-6 space-y-3">
				<center class="p-2">
					<h4>Team {Team.Name}</h4>
				</center>
                <span>
					<b>Visibility: </b>{Team.Visibility}
				</span>
				<span>
					<b>Technologies: </b>{Team.Technologies}
				</span>
				<span>
					<b>Availability: </b>{Team.Availability}
				</span>
				<span>
					<b>Description: </b>{Team.Description}
				</span>


			</Card>


            {/each}

			
		{/if}
	</Card>

</Page>
    
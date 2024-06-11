<script lang="ts">
    import Page from "../components/Page.svelte";
    import {Button, Card} from "flowbite-svelte";
    import { onMount } from 'svelte';
    import CodeJamTeam from '../models/team';
    import {getTeam} from "../services/services"
    import { Label, Input } from "flowbite-svelte";
	import type TeamMember from "../models/TeamMember";
    import CodeJamEvent from "../models/event";

    export let params: any; // set by svelte-spa-router

    let teamData: CodeJamTeam | null = null;
    let teamMembers: TeamMember | Array<any>; 
    let teamEvent: CodeJamEvent | null = null;
    console.log("params: ", params)


    // getTeam here returns an Object with {Team (CodeJamTeam), Event, Members(TeamMember)}
    // Am I getting Members and Events info from GEtTeamInfo or from... 

    async function loadData(id: string) {
        getTeam(params.id).then((response: any) => {
                console.log("line 17print: ", params.id, "== response: ", response)
                response.json().then((data: any) => {
                    console.log(data)
                    teamData = data.Team as CodeJamTeam;
                    teamMembers = data.Members
                    console.log("team members:", teamMembers)
                    teamEvent = data.Event as CodeJamEvent
                });
            }); 
    }
    $: {
        console.log(params)
        if (params) {
            loadData(params.id)
        }
    }

</script>
    <Page>
        hi
        <Card>
            {#if teamData !== null}
            {teamEvent?.Title}
            <center>
                <b>Team {teamData.Name}</b>
            </center>
            <span><b>Team Members: </b>
                <ul>
                    {#each teamMembers as member}
                        <li>{member.DisplayName} </li>
                    {/each}
                </ul>
                
            </span>

            <!--TODO: the the Event name and the Owner's name-->
            <span>
                <b>Visibility: </b>{teamData?.Visibility}
            </span>
            <span>
                <b>Technologies: </b>{teamData?.Technologies}
            </span>
            <span>
                <b>Availability: </b>{teamData?.Availability}
            </span>
            <span>
                <b>Description: </b>{teamData?.Description}
            </span>
            {/if}
            
            
        </Card>
    </Page>
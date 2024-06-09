<script lang="ts">
    import Page from "../components/Page.svelte";
    import {Button, Card} from "flowbite-svelte";
    import { onMount } from 'svelte';
    import CodeJamTeam from '../models/team';
    import {getTeam} from "../services/services"
    import { Label, Input } from "flowbite-svelte";
	import type TeamMember from "../models/TeamMember";

    export let params: any; // set by svelte-spa-router

    let teamData: CodeJamTeam | null = null;
    let teamMembers: TeamMember | null = null;
    console.log("params: ", params)

    async function loadData(id: string) {
        getTeam(params.id).then((response: any) => {
                console.log("line 17print: ", params.id, "== response: ", response)
                response.json().then((data: any) => {
                    console.log("line 19 data: ", data)
                    teamData = data.Team as CodeJamTeam;
                    console.log(teamData)
                    teamMembers = data[2] as TeamMember
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
            
            <center>
                <b>Team {teamData.Name}</b>
            </center>
            <span><b>Team Members: </b>
                <ul>
                    <!-- {#each teamData.Members as {DisplayName}}
                        <li>{DisplayName} </li>
                    {/each} -->
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
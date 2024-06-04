<script lang="ts">
    import Page from "../components/Page.svelte";
    import {Button, Card} from "flowbite-svelte";
    import { onMount } from 'svelte';
    import CodeJamTeam from '../models/team';
    import {getTeam} from "../services/services"
    import { Label, Input } from "flowbite-svelte";

    export let params: any; // set by svelte-spa-router

    let teamData: CodeJamTeam | null = null;
    console.log("params: ", params)

    async function loadData(id: string) {
        getTeam(params.id).then((response: any) => {
                console.log("line 17print: ", params.id, "== response: ", response)
                response.json().then((data: any) => {
                    console.log("line 19 data: ", data)
                    teamData = data as CodeJamTeam;
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
            
            <span>
                <b>Team Name: </b>{teamData?.Name}
            </span>
            <span>
                <b> Team Visibility: </b>{teamData?.Visibility}
            </span>
            
                
            {/if}
            
            
        </Card>
    </Page>
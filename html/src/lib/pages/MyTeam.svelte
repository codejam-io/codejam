<script lang="ts">

import Page from "../components/Page.svelte";
import {Button, Card} from "flowbite-svelte";
import {FontAwesomeIcon} from "@fortawesome/svelte-fontawesome";
import {faPenToSquare} from "@fortawesome/free-solid-svg-icons";
import CodeJamEvent from "../models/event";
import {onMount} from "svelte";
import {eventStatusStore} from "../stores/stores";



let statuses = null




</script>

    <Page>
    
        <Card size="xl" class="w-full">
            <h2>Team Options</h2>
            <div class="flex flex-row gap-8 my-8">
            </div>

            <Card class="my-2" size="xl">
                <h3>Find a Team to Join</h3>
                <p class="py-4">Browse other teams looking for members like you! </p>
                <p class="py-4"><Button href="/#/teams/browse">Browse</Button></p>
                
                
            </Card>
            <Card class="my-2" size="xl">
                <h3>Start a Team</h3>
                <p class="py-4">By creating a team, you're in control of the team page. Decide whether anyone can join or just people with the invite code. </p>
                <p class="py-4"><Button href="/#/teams/create">Create team</Button></p>
            </Card>
            
            <Card class="my-2" size="xl">
                <h3>Be a Solo Participant</h3>
                <p class="py-4">For those who want to work alone. You can change your mind later.</p>
                <p class="py-4"><Button href="/#/teams/create">Join solo</Button></p>
            </Card>
        </Card>
    
    </Page>

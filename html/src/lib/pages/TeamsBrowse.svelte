<script lang="ts">

    import Page from "../components/Page.svelte";
    import {Button, Card} from "flowbite-svelte";
    import {FontAwesomeIcon} from "@fortawesome/svelte-fontawesome";
    import {faPenToSquare} from "@fortawesome/free-solid-svg-icons";
    import CodeJamEvent from "../models/event";
    import {onMount} from "svelte";
    import {eventStatusStore} from "../stores/stores";
    
    

    </script>
    
        <Page>
        
            <Card size="xl" class="w-full">
                <h2>Current Teams</h2>
                <div class="flex flex-row gap-8 my-8">
                </div>
    
                <Card class="my-2" size="xl">
                    <h3>Team 1</h3>
                    <p class="py-4">Join us, we're the best! </p>
                    <p class="py-4"><Button href="/#/teams/browse">Join</Button></p>
                    
                    
                </Card>
            </Card>
        
        </Page>
    
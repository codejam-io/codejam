<script>
import TextGoals from "./Info/TextGoals.svelte";
import TextTimeline from "./Info/TextTimeline.svelte";
import TextRules from "./Info/TextRules.svelte";
import { activeContent, loggedIn } from "../stores";
	import SubProfile from "./SubProfile.svelte";


</script>
<style>
    .card-hidden {
        display: none;
    }
    .card-module {
        display: flex;
        position: fixed;
        border-radius: 15px;
        top: 50;
        left: left; 
        transform: translate(-50%, -50%);
    }
</style>


<div class="content"> 


    <div id="profile"><SubProfile /></div>

    <!-- I want these divs to appear as a module: -->
    <div id="goals" class:card-hidden={!$loggedIn} class:card-module={$loggedIn} class="card"><TextGoals /></div>
    <div id="timeline" class="card"><TextTimeline /></div>
    <div id="rules" class="card"><TextRules /></div>

<!-- 
    {#if !$loggedIn} -->
    <!-- Not logged in Content Display -->
    <!-- <div id="goals" class="card"><TextGoals /></div>
    <div id="timeline" class="card"><TextTimeline /></div>
    <div id="rules" class="card"><TextRules /></div>
    {/if} -->

</div>

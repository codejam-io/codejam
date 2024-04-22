<script>
import TextGoals from "./Info/TextGoals.svelte";
import TextTimeline from "./Info/TextTimeline.svelte";
import TextRules from "./Info/TextRules.svelte";
import { activeContent, loggedIn } from "../stores";
import SubProfile from "./SubProfile.svelte";


</script>
<style>
    .tg {
        background-color: rgb(255, 234, 241);
    }
    .tt {
        background-color: rgb(237, 224, 250);
    }
    .tr {
        background-color: rgb(206, 233, 243);
    }
</style>


<div class="content"> 


    {#if $loggedIn}
    <!-- profile not a module: -->
    <div id="profile" class="card"><SubProfile /></div>
    <!--I want these divs to appear as a module: -->
    {:else}
    <!-- Not logged in Content Display -->

    <div id="goals" class="card tg"><TextGoals /></div>
    <div id="timeline" class="card tt"><TextTimeline /></div>
    <div id="rules" class="card tr"><TextRules /></div>
    {/if}

    
</div>

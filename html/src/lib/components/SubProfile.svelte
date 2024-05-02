<script>
    import Card from "./Card.svelte";
    import BtnCreateTeam from "./BtnCreateTeam.svelte";
    //import CreateTeamForm from "./CreateTeamForm.svelte";
    import {userStore, loggedInStore} from "./../stores/stores";

</script>
<div class="flex center col">
    <h1>Hi {$userStore?.DisplayName}!</h1>
    <h1>Profilad participate in the codej asdae/</h1>
    <p>In order to participate in the codejam, first step is to join a team. You can scroll through teams on the right, or you can <span><a href="#">create a new team</a></span>.</p>
    <h1>Filler Text Header/</h1>
    <p>stuff stuff order to participate in the codejam, first order to participate in the codejam, first order to participate in the codejam, first order to participate in the codejam, first  </p>
    <!-- <BtnCreateTeam />  -->
</div>
<!-- <CreateTeamForm /> -->
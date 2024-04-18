<script lang="ts">
    import '../../styles/styles.css';
    import Nav from "../components/Nav.svelte";
    import Banner from "../components/Banner.svelte";
    import GradientChange from "../components/GradientChange.svelte";
    import Card from "../components/Card.svelte";
    import { activeContent, loggedIn } from '../stores';
    import Panel from '../components/Panel.svelte';
	import Content from '../components/Content.svelte';

</script>
<style>

</style>

<div class="flex col">
    
    <!-- <GradientChange /> -->
    
    <!-- <Card /> -->

    <!-- <div class="content">
        {#if $activeContent === 'goals'}
        <div class="card">{@html goals}</div>
        {:else if $activeContent === 'timeline'}
        <div class="card">{@html timeline}</div>
        {/if}
    </div> -->

    <div class="flex row">
        <div id="left-side" class="flex col">
            <Banner size="normal" />
            <Panel />
            <Nav />
            <Content />
    
        </div>
        <div id="right-side" class="flex col">
            <Card />
            <Card />
            
        </div>
    
    </div>


</div>
<script lang="ts">
    import '../styles/styles.css';
    import Nav from "../lib/components/Nav.svelte";
    import LoginBtn from "../lib/components/LoginBtn.svelte";
    import Banner from "../lib/components/Banner.svelte";
    import { activeContent } from '../lib/stores';

    let src = '/toast1.png'
    let goals = "the goals of the <b>codejam</b> include..."
    let timeline = "codejam will begin in either July or August"

</script>
<style>
    .content {
        padding: 3rem 5rem;
    }
    .panel {
        height: 50vh;
        border: 5px solid blue;
        padding: 3rem 5rem;
    }
</style>
<div class="flex col">
    
    <Nav />
    <Banner />

    <div class="flex row sp-around">
        <LoginBtn />
    </div>

    <div class="flex panel" >
        <div class="panel-text">
            Ingredients for Codejam:
            <ul>
                <li>
                    sugar
                </li>
                <li>
                    code
                </li>
            </ul>
        </div>
        <img {src} alt="toast with jam and code on it">
    </div>

    <div class="content">
        {#if $activeContent === 'goals'}
        <div>{@html goals}</div>
        {:else if $activeContent === 'timeline'}
        <div>{@html timeline}</div>
        {/if}
    </div>

</div>
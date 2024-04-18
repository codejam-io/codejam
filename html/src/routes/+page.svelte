<script lang="ts">
    import '../styles/styles.css';
    import Nav from "../lib/components/Nav.svelte";
    import LoginBtn from "../lib/components/LoginBtn.svelte";
    import Banner from "../lib/components/Banner.svelte";
    import Card from "../lib/components/Card.svelte";
    import { activeContent } from '../lib/stores';

    let src = '/pinktoasttransparent.png'
    let goals = "the goals of the <b>codejam</b> include..."
    let timeline = "codejam will begin in either July or August"

</script>
<style>
    .content {
        padding: 3rem 5rem;
    }
    .panel {
        height: 30vh;
        border: 0px solid blue;
        margin: 3rem 5rem;
    }
    .card {
        padding: 3rem 5rem;
        color: black;
        border-radius: 50px;
        margin-left: 3rem;
        margin-right: 3rem;

        border-radius: 50px;
        background: #e0e0e0;
        /* box-shadow:  -20px 20px 60px #bebebe,
                    20px -20px 60px #ffffff; */
    }
</style>

<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Fira+Code:wght@300..700&family=Liu+Jian+Mao+Cao&family=Space+Grotesk:wght@300..700&family=Ubuntu+Mono:wght@400;700&display=swap" rel="stylesheet">

<div class="flex col">
    
    <Nav />
    <Banner />

    <div class="flex panel center">
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
            <div class="flex row sp-around">
                <LoginBtn />
            </div>
        </div>
        <img {src} alt="toast with jam and code on it">
    </div>
    
    <!-- <Card /> -->

    <!-- <div class="content">
        {#if $activeContent === 'goals'}
        <div class="card">{@html goals}</div>
        {:else if $activeContent === 'timeline'}
        <div class="card">{@html timeline}</div>
        {/if}
    </div> -->


        <div class="card">{@html goals}</div>

        <div class="card">{@html timeline}</div>



</div>
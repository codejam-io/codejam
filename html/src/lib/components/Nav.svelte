<script lang="ts">
	import LoginBtn from './LoginBtn.svelte';
    import { activeContent } from '../stores';

    function setActiveContent(content: string) {
        activeContent.set(content)
    }
</script>

<nav>
	<ul class="flex row sp-between nav-ul">
		<li><a href="#" on:click|preventDefault={()=> setActiveContent('goals')}>Goal</a></li>  
		<li><a href="#" on:click|preventDefault={()=> setActiveContent('timeline')}>Timeline</a></li>
		<li><a href="#" on:click|preventDefault={()=> setActiveContent('rules')}>Rules</a></li>
	</ul>
</nav>

<style>
	.nav-ul {
		padding: 1rem;
		justify-content: space-around;
	}
	li {
		list-style-type: None;
		gap: 20px;
	}
</style>

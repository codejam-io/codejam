<script lang="ts">
    import { activeContent } from '../stores';

    function setActiveContent(content: string) {
        activeContent.set(content)
        if ($loggedIn) {
            document.getElementById(content)?.classList.add('card-module')
        }
    }
    import { loggedIn } from '../stores'; // Import the loggedIn store

</script>

<nav>
	<ul class="flex row sp-between nav-ul">

		<li><a href="#" on:click|preventDefault={()=> setActiveContent('goals')}>Goal</a></li>  
		<li><a href="#" on:click|preventDefault={()=> setActiveContent('timeline')}>Timeline</a></li>
		<li><a href="#" on:click|preventDefault={()=> setActiveContent('rules')}>Rules</a></li>
        <li><a href="#" on:click|preventDefault={()=> setActiveContent('profile')}>Profile</a></li>

        <!-- I want to use this only if accessed from #/ in svelte
        <li><a href="#goals">Goals</li>
        <li><a href="#timeline">Timeline</li>
        <li><a href="#rules">Rules</li> -->
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
        font-size: 1rem;
	}
</style>

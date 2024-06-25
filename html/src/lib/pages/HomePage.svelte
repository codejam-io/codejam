<script lang="ts">
	import Page from "../components/Page.svelte";
	import {activeEventStore} from "../stores/stores";

</script>

<Page>
	<div class="flex flex-col">
		<div class="flex panel center row my-8">
			<div class="panel-text flex col">
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
			<div class="flex col">
				<img src="/static/pinktoasttransparent.webp" class="object-contain max-h-72" alt="toast with jam and code on it">
			</div>
		</div>

		{#if $activeEventStore !== null}
			<div id="timeline" class="card !bg-[#ede0fa]">
				<h3>Timeline</h3>
				<pre class="whitespace-pre-wrap">{@html $activeEventStore?.Timeline}</pre>
			</div>
			<div id="goals" class="card !bg-pink-100">
				<h3>Goals</h3>
				<pre class="whitespace-pre-wrap">{@html $activeEventStore?.Description}</pre>
			</div>
			<div id="rules" class="card !bg-[#cee9f3]">
				<h3>Rules</h3>
				<pre class="whitespace-pre-wrap">{@html $activeEventStore?.Rules}</pre>
			</div>
		{:else}
			<h1>No Active Events, Stay Tuned!</h1>
		{/if}
	</div>
</Page>


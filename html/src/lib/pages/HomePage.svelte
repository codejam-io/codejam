<script lang="ts">
	import Page from "../components/Page.svelte";
	import {eventStore} from "../stores/stores";

</script>

<Page>
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
			<img src="/static/pinktoasttransparent.png" class="object-contain max-h-72" alt="toast with jam and code on it">
		</div>
	</div>

	<div id="timeline" class="card !bg-[#ede0fa]">
		<h3>Timeline</h3>
		{$eventStore?.Timeline}
	</div>
	<div id="goals" class="card !bg-pink-100">
		<h3>Goals</h3>
		<pre class="whitespace-pre-wrap">{$eventStore?.Description}</pre>
	</div>
	<div id="rules" class="card !bg-[#cee9f3]">
		<h3>Rules</h3>
		{$eventStore?.Rules}
	</div>
</Page>


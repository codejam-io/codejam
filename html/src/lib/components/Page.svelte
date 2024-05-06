<script lang="ts">
    import '../../styles/styles.css';
    import Header from "./Header.svelte";

</script>

<div class="flex col px-8">
    <Header/>
    <slot/>
</div>
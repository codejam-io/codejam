<script lang="ts">
    import '../../styles/styles.css';
    import Header from "./Header.svelte";

</script>

<div class="w-full">
    <Header/>
    <div class="flex flex-row justify-center px-8 w-full">
        <slot/>
    </div>
</div>


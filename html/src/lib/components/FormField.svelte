<script lang="ts">

    // TODO - using svelte/internal may not be ideal, look for alternative way
    import { get_current_component } from "svelte/internal";
    import {getContext, onMount} from "svelte";

    export let name : string = '';

    export let error : string = '';

    let register = getContext("register");
    let this_component = get_current_component();

    export function setError(err: string) {
        error = err;
    }

    onMount(() => {
        register(this_component, name, setError);
    });

</script>


<div class="flex flex-col gap-2">
    <span class="fieldLabel">{name}</span>
    <slot/>
    <span class="fieldError">{error}</span>
</div>



<style>
    .fieldLabel {
        color: #101010;
        font-weight: bold;
        margin-left: 10px;
    }

    .fieldError {
        color: #E02424;
        font-weight: bold;
        margin-left: 10px;
    }
</style>




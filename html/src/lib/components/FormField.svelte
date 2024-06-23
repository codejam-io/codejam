<script lang="ts">

    import {getContext, onMount} from "svelte";

    // form field name
    export let name : string = '';

    export let error : string = '';

    // form field label
    export let label: string = '';

    export function setError(err: string) {
        error = err;
    }

    let formFunctions: any = getContext("formFunctions");

    onMount(() => {
        formFunctions.register(name, setError);
    });

</script>


<div class="flex flex-col gap-2">
    <span class="fieldLabel">{label}</span>
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




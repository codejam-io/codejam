<script lang="ts">
import routes from "./routes";
import Router from "svelte-spa-router";
import {Toaster} from "svelte-sonner";
</script>


<Toaster richColors/>

<main>
    <Router {routes}/>
</main>
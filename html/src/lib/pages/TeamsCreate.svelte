<script lang="ts">

    import Page from "../components/Page.svelte";
    import {Breadcrumb, BreadcrumbItem, Button, Card} from "flowbite-svelte";
    import {FontAwesomeIcon} from "@fortawesome/svelte-fontawesome";
    import {faPenToSquare} from "@fortawesome/free-solid-svg-icons";
    import CodeJamTeam from "../models/team";
    import {onMount} from "svelte";
    import {eventStatusStore} from "../stores/stores";
    import CreateTeamForm from "../components/CreateTeamForm.svelte";
    
</script>
    
    <Page>
        <Breadcrumb solid class="mb-4 w-full max-w-screen-xl">
            <BreadcrumbItem href="/#/" home>Home</BreadcrumbItem>
            <BreadcrumbItem href="/#/team">Team Options</BreadcrumbItem>
            <BreadcrumbItem>Create Team</BreadcrumbItem>
        </Breadcrumb>
        <Card size="xl" class="w-full">
            <h2>Create your team!</h2>
            <div class="flex flex-row gap-8 my-8">
            </div>
            <CreateTeamForm />
        </Card>

        
    
    </Page>
    
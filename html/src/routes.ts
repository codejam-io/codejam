import HomePage from "./lib/pages/HomePage.svelte";
import EventEdit from "./lib/pages/EventEdit.svelte";
import EventList from "./lib/pages/EventList.svelte";
import MyTeam from "./lib/pages/MyTeam.svelte";
import TeamsBrowse from "./lib/pages/TeamsBrowse.svelte";
import TeamsCreate from "./lib/pages/TeamsCreate.svelte";

export default {
    '/': HomePage,
    '/home': HomePage,
    '/admin/events': EventList,
    '/admin/event/:id': EventEdit,
    '/my-team': MyTeam,
    '/teams/browse': TeamsBrowse,
    '/teams/create': TeamsCreate
}
import HomePage from "./lib/pages/HomePage.svelte";
import EventEdit from "./lib/pages/EventEdit.svelte";
import EventList from "./lib/pages/EventList.svelte";
import TeamOptions from "./lib/pages/TeamOptions.svelte";
import TeamsBrowse from "./lib/pages/TeamsBrowse.svelte";
import TeamsCreate from "./lib/pages/TeamsCreate.svelte";
import MyTeam from "./lib/pages/MyTeam.svelte";
import Invite from "./lib/pages/Invite.svelte";


export default {
    '/': HomePage,
    '/home': HomePage,
    '/admin/events': EventList,
    '/admin/event/:id': EventEdit,
    '/team': TeamOptions,
    //'/team/my-teams/': MyTeams // link to all your teams
    '/team/:id': MyTeam, // link to one of your teams (sharable)
    '/team/invite/:id': Invite,
    '/teams/browse': TeamsBrowse,
    '/teams/create': TeamsCreate,
}
<script lang="ts">

    import {FontAwesomeIcon} from "@fortawesome/svelte-fontawesome";

    import {
        type IconLookup,
        type IconDefinition,
        findIconDefinition,
        library
    } from '@fortawesome/fontawesome-svg-core';
    import {
        faDiscord
    } from '@fortawesome/free-brands-svg-icons';
    library.add(faDiscord)
    const discordLookup: IconLookup = { prefix: 'fab', iconName: 'discord' }
    const discordIconDefinition: IconDefinition = findIconDefinition(discordLookup)

    export let size : string = '1x';
</script>

<FontAwesomeIcon icon={discordIconDefinition} size={size}/>
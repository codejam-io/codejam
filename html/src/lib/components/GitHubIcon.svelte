<script lang="ts">

    import {FontAwesomeIcon} from "@fortawesome/svelte-fontawesome";

    import {
        type IconLookup,
        type IconDefinition,
        findIconDefinition,
        library
    } from '@fortawesome/fontawesome-svg-core';
    import {
        faSquareGithub
    } from '@fortawesome/free-brands-svg-icons';
    library.add(faSquareGithub)
    const githubLookup: IconLookup = { prefix: 'fab', iconName: 'square-github' }
    const githubIconDefinition: IconDefinition = findIconDefinition(githubLookup)

    export let size : string = '1x';
</script>

<FontAwesomeIcon icon={githubIconDefinition} size={size}/>
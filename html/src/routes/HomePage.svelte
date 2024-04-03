<script lang="ts">
  import {faGithub} from "@fortawesome/free-brands-svg-icons";
  import {Button} from "flowbite-svelte";
  import {Fa} from "svelte-fa";

  function githubLogin() {
    window.location = "https://github.com/login/oauth/authorize?client_id=a6e4222158ba2bb454e3&redirect_url=https://codejam.io/oauth/callback&allow_signup=false";
  }

</script>

<Button on:click={githubLogin}><Fa icon={faGithub}/>Sign in with GitHub</Button>


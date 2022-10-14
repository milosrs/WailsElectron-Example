<script lang="ts">
  import Login from "./components/Login/Login.svelte";
  import Profiles from './components/Profiles/Profiles.svelte';

  var loggedIn: boolean = false;

  function onLoginSuccess() {
    loggedIn = true
  }
</script>

<main id="main">
  {#if !loggedIn}
    <Login onLogin={onLoginSuccess}/>    
  {:else}
    <Profiles/>
  {/if}
</main>
<style>
  #main {
    height: 100%;
    width: 100%;
  }
</style>


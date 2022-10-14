<script lang="ts">
    import {Login} from '../../../wailsjs/go/main/App.js'
  
    let error = ""
    let username:  string = "";
    let password: string = "";
    let loginDisabled = true;

    export let onLogin: () => void;
  
    function login(): void {
      loginDisabled = true;

      Login(username, password)
      .then(resp => {
        console.log('Got token: ', resp);
        window.localStorage.setItem('token', resp)
        onLogin()
      })
      .catch(err => {
        error = err;
        loginDisabled = false;
      })
    }
  
    function onInputChange(): void {
      error = ""
      loginDisabled = password.length === 0 || username.length === 0;
    }
</script>
<div id="login-page">
    <div class="login-box">
        <h1 class="h3 mb-3 fw-normal">Please sign in</h1>

        <div class="form-floating">
            <input type="email" bind:value={username} class="form-control" id="floatingInput" placeholder="name@example.com" on:keyup={onInputChange} on:change={onInputChange}>
            <label for="floatingInput">Email address</label>
        </div>
        <div class="form-floating">
            <input type="password" bind:value={password} class="form-control" id="floatingPassword" placeholder="Password" on:keyup={onInputChange} on:change={onInputChange}>
            <label for="floatingPassword">Password</label>
        </div>
        <div class="form-floating">
            {#if error != ""}
            <div class="error-label">{error}</div>
            {/if}
        </div>    
        <button id="loginBtn" class="w-100 btn btn-lg btn-primary" type="submit" on:click={login} disabled={loginDisabled}>Sign in</button>
    </div>
</div>
<style>

    #loginBtn[disabled=true] {
      cursor: not-allowed;
    }

    #loginBtn[disabled=false] {
      cursor: pointer;
    }

    #login-page {
      color:black;
      display: flex;
      justify-content: center;
      align-items: center;
      background-color: black;
      height: 100%;
    }
  
    .login-box {
      width: 50%;
    }
  
    .login-box > * {
      margin-bottom: 6%;   
    }
  
    .error-label {
      color: red;
      text-align: left;
    }
  
    h1 {
      color: white;
      text-align: center;
    }
  
  </style>
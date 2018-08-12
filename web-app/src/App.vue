<template>
  <div id="app">
    <h1>Human Resources App</h1>
    <div class="app__cool-side-bar"></div>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
* {
    font-family: "Trebuchet MS", Helvetica, sans-serif;
}
body {
    background: #f54a61;
}
input {
    border: 0;
    border-radius: 9px;
    font-size: 1em;
    margin: .4em;
    padding: 1em;
    background: '##f3cdd2';
}
button {
    margin: 1em .5em;
    padding: .5em 1em;
    font-size: 1em;
    border-radius: 2em;
    background: white;
    color: #f54a61;
    font-weight: bold;
    border-width: 0;
}
#app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
    padding: 2em;
}
.app__cool-side-bar {
    background: #d4384c;
    height: 100%;
    left: 0;
    position: fixed;
    top: 0;
    width: 0.5em;
}
</style>

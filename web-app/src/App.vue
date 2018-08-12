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
    position: absolute;
    top: 0;
    width: 0.5em;
}
</style>

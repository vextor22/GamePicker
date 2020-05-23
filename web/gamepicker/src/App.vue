<template>
  <div id="app">
    <h1>{{ msg }}</h1>
    <app-test></app-test>
    <p
      v-for="(game, index) in info.games"
      :key="index"
    >{{ game.name}}: {{ (game.playtime_forever / 60).toFixed(0) }} Hours</p>
    <p>poop</p>
  </div>
</template>

<script>
const axios = require("axios").default;

import Test from "./components/Test.vue";
export default {
  name: "app",
  data() {
    return {
      msg: "Welcome to Your Vue.js App again",
      info: null
    };
  },
  components: {
    appTest: Test
  },
  mounted() {
    axios
      .get("./app/user/76561197995406081")
      .then(response => (this.info = response.data));
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>

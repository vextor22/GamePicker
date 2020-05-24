<template>
  <div id="app" class="container">
    <h1>{{ msg }}</h1>
    <user-submit @submitSteamID="requestUserData"></user-submit>
    <div v-if="info">
      <p
        v-for="(game, index) in info.games"
        :key="index"
      >{{ game.name}}: {{ (game.playtime_forever / 60).toFixed(2) }} Hours</p>
    </div>
  </div>
</template>

<script>
import UserSubmit from "./components/UserSubmit.vue";
const axios = require("axios").default;

export default {
  name: "app",
  data() {
    return {
      msg: "Steam Game Picker",
      info: null
    };
  },
  methods: {
    requestUserData(event) {
      axios
        .get(`./app/user/${event}`)
        .then(response => (this.info = response.data));
    }
  },
  components: {
    userSubmit: UserSubmit
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

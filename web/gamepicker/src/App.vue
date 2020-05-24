<template>
  <div id="app" class="container-fluid">
    <h1>{{ msg }}</h1>
    <user-submit @submitSteamID="requestUserData"></user-submit>
    <div class="row justify-content-sm-center" style="margin: 20px;">
        <div class="col-xs-6 col-sm-6 col-md-6 col-lg-6">
          <div class="card">

          <game-list :_games="info.games"></game-list>
          </div>
        </div>
        <div class="col-xs-6 col-sm-6 col-md-6 col-lg-6">
          <div class="card">

          <game-list :_games="info.games"></game-list>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import UserSubmit from "./components/UserSubmit.vue";
import GameList from "./components/GameList.vue";
const axios = require("axios").default;

export default {
  name: "app",
  data() {
    return {
      msg: "Steam Game Picker",
      info: {}
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
    userSubmit: UserSubmit,
    gameList: GameList
  }
};
</script>

<style scoped>
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

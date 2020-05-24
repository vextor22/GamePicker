<template>
  <div class="card row justify-content-sm-center">
    <ul v-if="games" class="list-group game-list" style>
      <li
        class="list-group-item text-left"
        v-for="(game, index) in games"
        :key="index"
      >{{ game.name}}: {{ (game.playtime_forever / 60).toFixed(2) }} Hours</li>
    </ul>
    <div class="card align-items-center" v-else>
      <div class="card-text">
        <h2>Instructions</h2>
        <p>
          By default, Steam hides user's game info. To make this information public,
          please update your
          <a
            href="https://steamcommunity.com/my/edit/settings"
          >privacy information</a>.
        </p>
        <div class="row">
          <ol class="list-group">
            <li class="list-group-item text-left">
              1. Follow the link above, or click
              <b>Edit Profile</b>
            </li>
            <li class="list-group-item text-left">
              2. Navigate to
              <b>My Privacy Settings</b>
            </li>
            <li class="list-group-item text-left">
              3. Change
              <b>Game Details</b> to
              <b>Public</b> and uncheck box for
              <b>private playtime</b>
            </li>
          </ol>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    _games: {
      type: Array
    }
  },
  computed: {
    games() {
      if (this._games) {
        console.log("Thing here");
        // ES6 shallow copy of this._games before sorting to avoid side effect
        return [...this._games].sort((a, b) =>
          a.playtime_forever < b.playtime_forever ? 1 : -1
        );
      }

      return null;
    }
  }
};
</script>

<style scoped>
.row {
  margin: 20px;
}
.game-list {
  margin: 20px;
  padding: 5px;
  height: 60vh;
  overflow-y: auto;
  -webkit-overflow-scrolling: touch;
}
</style>

<template>
  <div class="card-body game-card">
    <h2 class="card-title">Random Game</h2>
    <div class="row">
      <div class="row" style="width: 100%;" v-if="gameIndex">
        <div class="game-display">
          <div>
            <img alt="Game Logo" :src="gameLogo" class="image" />
          </div>
          <div class="description">
            <p>{{ _games[gameIndex].name }}</p>
          </div>
        </div>
        <div class="align-left">
          <div class="btn btn-primary" @click="newGame">Pick Another Game</div>
        </div>
      </div>
      <div class="text-left" v-else>
        <div class="card-text">
          <p>Your randomly selected game will be displayed here.</p>
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
  data() {
    return {
      data: 0
    };
  },
  methods: {
    randomInt(max) {
      return Math.floor(Math.random() * max);
    },
    newGame() {
      this.$emit("rerender", {});
    }
  },
  computed: {
    gameIndex() {
      if (this._games) {
        return this.randomInt(this._games.length);
      }
      return null;
    },
    gameLogo() {
      return `http://media.steampowered.com/steamcommunity/public/images/apps/${
        this._games[this.gameIndex].appid
      }/${this._games[this.gameIndex].img_logo_url}.jpg`;
    }
  }
};
</script>

<style scoped>
.row {
  margin: 20px;
}
.game-card {
  margin: 20px;
  padding: 5px;
  height: 60vh;
}
.game-display {
  min-height: 100px;
  width: 100%;
  border: 2px solid darkgray;
  padding: 5px;
  box-shadow: 4px 2px 2px gray;
  overflow: hidden;
}

.description {
  border-left: 2px solid darkgray;
  margin-left: 170px;
  height: 100px;
}
.image {
  margin: 5px 5px 5px 5px;
  height: 80px;
  width: 160px;
  float: left;
  background-color: black;
}
</style>

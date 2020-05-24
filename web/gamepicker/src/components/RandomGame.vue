<template>
  <div class="card-body game-card">
    <h2 class="card-title">Random Game</h2>
    <div class="row">
      <div class="row" style="width: 100%;" v-if="game">
        <div class="game-display">
          <div>
            <img alt="Game Logo" class="image" />
          </div>
          <div class="description">
            <p :key="rerender">{{ game.name }}</p>
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
    game() {
      if (this._games) {
        return this._games[this.randomInt(this._games.length)];
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
.game-card {
  margin: 20px;
  padding: 5px;
  height: 60vh;
}
h2 {
  background-color: darkorange;
}
.game-display {
  min-height: 100px;
  width: 100%;
  border: 2px solid darkgray;
  background-color: #cccccc;
  padding: 5px;
  box-shadow: 4px 2px 2px gray;
  overflow: hidden;
}

.description {
  border-left: 2px solid black;
  margin-left: 95px;
  height: 100px;
}
.image {
  margin: 5px 5px 5px 5px;
  height: 80px;
  width: 80px;
  float: left;
  background-color: black;
}
</style>

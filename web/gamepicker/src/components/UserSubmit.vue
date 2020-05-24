<template>
  <div class="row justify-content-sm-center">
    <div class="col-sm-6">
      <div class="steamid-input card">
        <form>
          <div class="form-group" style="width: 75%; margin: auto;">
            <label for="SteamID">SteamID</label>
            <input
              type="text"
              class="form-control"
              id="SteamID"
              aria-describedby="steamID Input"
              placeholder="Enter SteamID"
              v-model="steamID"
            />
            <button type="submit" @click.prevent="submitted" class="btn btn-primary">Submit SteamID</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script scoped>
export default {
  data() {
    return {
      steamID: ""
    };
  },
  methods: {
    submitted(e) {
      console.log("Clicked me", this.steamID);
      this.$emit("submitSteamID", this.steamID);
    }
  }
};
</script>

<style>
.steamid-input {
  padding: 10px;
}
.btn {
  margin: 5px;
}
</style>

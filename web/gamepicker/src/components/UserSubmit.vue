<template>
  <div class="row justify-content-center">
    <div class="col-sm-6">
      <div class="steamid-input card">
        <form>
          <div class="form-group" style="width: 75%; margin: auto;">
            <label for="SteamID">SteamID</label>
            <input
              type="text"
              :class="{'form-control': true, 'is-invalid': warningText}"
              id="SteamID"
              aria-describedby="steamID Input"
              placeholder="Enter SteamID"
              @input="userChanged"
            />
            <small class="text-danger" v-if="warningText">{{ warningText }}</small>
            <small v-if="steamID">Potential ID: {{ steamID }}</small>
          </div>
          <div class="form-group" style="width: 75%; margin: auto;">
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
      steamID: "",
      warningText: ""
    };
  },
  methods: {
    submitted(e) {
      console.log("Clicked me", this.steamID);
      this.$emit("submitSteamID", this.steamID);
    },
    userChanged(e) {
      var userID = e.target.value;
      if (userID.includes("steamcommunity")) {
        var splits = userID.split("/");
        for (const [i, value] of splits.entries()) {
          console.log("%d: %s", i, value);
          if (value === "profiles" || value === "id") {
            var possibleID = splits[i + 1];
            if (possibleID) {
              this.warningText = "";
              this.steamID = splits[i + 1];
              console.log("ID Found: %s", this.steamID);
              return;
            }
          }
        }
        this.steamID = "";
        this.warningText = "I haven't found a user ID in this URL";
      } else if (userID !== "") {
        this.steamID = userID;
      } else {
        this.steamID = "";
        this.warningText = "";
      }
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

<template>
  <div class="container">
    <div class="columns is-vcentered is-mobile" style="height: 100vh">
      <div class="column is-three-fifths is-offset-one-fifth">
        <div class="field">
          <label class="label">Username</label>
          <div class="control">
            <input
              :class="usernameInputClass"
              type="text"
              placeholder="Username"
              v-model="username"
              autocomplete="true"
            />
          </div>
          <p :class="usernameHelpClass">{{ usernameHelpMessage }}</p>
        </div>

        <div class="field">
          <label class="label">Password</label>
          <div class="control">
            <input
              :class="passwordInputClass"
              type="password"
              v-model="password"
              placeholder="Password"
            />
          </div>
          <p :class="passwordHelpClass">{{ passwordHelpMessage }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "LoginForm",
  props: {
    //
  },
  data: function () {
    return {
      username: "",
      password: "",
    };
  },
  computed: {
    usernameIsValid(): {
      valid: boolean;
      showFeedback: boolean;
      message: string;
    } {
      const obj = {
        valid: false,
        showFeedback: this.username.length != 0,
        message: "",
      };
      if (this.username.length < 3) {
        obj.message = "Username needs to be atleast 3 charcters";
      } else {
        obj.valid = true;
        obj.message = "Nice Username!";
      }
      return obj;
    },
    passwordIsValid(): any {
      const obj = {
        valid: false,
        showFeedback: this.password.length != 0,
        message: "",
      };
      if (this.password.length <= 8) {
        obj.message = "Password needs to be atleast 8 charcters";
      } else {
        obj.valid = true;
        obj.message = "Nice Password!";
      }
      return obj;
    },
    usernameHelpClass(): any {
      return {
        help: true,
        "is-danger":
          !this.usernameIsValid.valid && this.usernameIsValid.showFeedback,
        "is-success":
          this.usernameIsValid.valid && this.usernameIsValid.showFeedback,
      };
    },
    usernameInputClass(): any {
      return {
        input: true,
        "is-danger":
          !this.usernameIsValid.valid && this.usernameIsValid.showFeedback,
        "is-success": this.usernameIsValid && this.usernameIsValid.showFeedback,
      };
    },
    usernameHelpMessage(): any {
      return this.usernameIsValid.showFeedback
        ? this.usernameIsValid.message
        : "";
    },
    passwordHelpClass(): any {
      return {
        help: true,
        "is-danger":
          !this.passwordIsValid.valid && this.passwordIsValid.showFeedback,
        "is-success":
          this.passwordIsValid.valid && this.passwordIsValid.showFeedback,
      };
    },
    passwordInputClass(): any {
      return {
        input: true,
        "is-danger":
          !this.passwordIsValid.valid && this.passwordIsValid.showFeedback,
        "is-success":
          this.passwordIsValid.valid && this.passwordIsValid.showFeedback,
      };
    },
    passwordHelpMessage(): string {
      return this.passwordIsValid.showFeedback
        ? this.passwordIsValid.message
        : "";
    },
  },
  methods: {
    login() {
      this.$store.commit("LOGIN", "kjafhkahdakshdka");
      this.$router.push({ path: "/" });
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>

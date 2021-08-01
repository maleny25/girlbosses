<template>
  <div id="signup-form">
    <h1>Create an Account</h1>
    <form @submit.prevent="handleSubmit">

      <label>Username: </label>
      <input
        ref="first"
        type="text"
        :class="{ 'has-error': submitting && invalidUsername }"
        v-model="user.username"
        @focus="clearStatus"
        @keypress="clearStatus"
      >
      <br>
      <label>Password: </label>
      <input
        type="password"
        :class="{ 'has-error': submitting && invalidPassword }"
        v-model="user.password"
        @focus="clearStatus"
      >
      <br>
      <label>Confirm Password: </label>
      <input
        type="password"
        :class="{ 'has-error': submitting && noMatch }"
        v-model="user.confirmPassword"
        @focus="clearStatus"
      >
      <br>
      <button>Create Account</button>
      <p
        v-if="error && submitting"
        class="error-message"
      >❗Please fill out all required fields</p>
      <p
        v-if="unmatch && submitting"
        class="error-message"
      >❗Passwords don't match</p>
      <p
        v-if="success"
        class="success-message"
      >✅ Account successfully created</p>
    </form>

  </div>
</template>

<script>
export default {
  name: 'signup-form',
  data() {
    return {
      error: false,
      submitting: false,
      success: false,
      unmatch: false,
      user: {
        username: '',
        password: '',
      }
    }
  },
  computed: {
    invalidUsername() {
      return this.user.username === ''
    },
    invalidPassword() {
      return this.user.password === ''
    },
    noMatch() {
      return this.user.confirmPassword !== this.user.password
    },
  },
  methods: {
    handleSubmit() {
      this.clearStatus()
      this.submitting = true
      if (this.noMatch && !this.invalidUsername) {
        this.unmatch = true
        return
      }
      if (this.invalidUsername || this.invalidPassword ) {
        this.error = true
        return
      }
      this.$emit('create:user', this.user)
      this.$refs.first.focus()
      this.user = {
        username: '',
        password: '',
      }
      this.success = true
      this.error = false
      this.submitting = false
      this.unmatch = false
    },
    clearStatus() {
      this.success = false
      this.error = false
      this.unmatch = false
    }
  }}
</script>

<style scoped>
h1 {
  margin-block-start: 0;
  margin-block-end: 0;
}
form {
  margin-bottom: 0;
  margin-top: 0;
}
[class*="-message"] {
  font-weight: 500;
}
.error-message {
  color: #d33c40;
}
.success-message {
  color: #32a95d;
}
button {
  background-color: lavender;
  border-color: lightgrey;
  border-radius: 40px;
  font-size: 15px;
  margin-top: 2px;
  height: 30px;
}
input {
  height: 20px;
  margin: 2px;
  width: 200px;
}

</style>
<template>
    <div id="login-form">
        <h1>Login</h1>
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
            <button>Login</button>
            <p
                v-if="error && submitting"
                class="error-message"
            >❗Username or password are incorrect</p>
            <p
                v-if="success"
                class="success-message"
            >✅ Successfully logged in</p>
        </form>
    </div>
</template>

<script>
export default {
  name: 'login-form',
  data() {
    return {
      error: false,
      submitting: false,
      success: false,
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
  },
  methods: {
    handleSubmit() {
      this.clearStatus()
      this.submitting = true
      if (this.invalidUsername || this.invalidPassword ) {
        this.error = true
        return
      }
      this.$emit('login:user', this.user)
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
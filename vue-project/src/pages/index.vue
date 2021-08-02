<template>
  <div id="app">
    <Navigation/>
    <h1>Welcome to G³</h1>
    <h4>
      G³ is a platform that helps minority students find events and 
      find other students to network with.
    </h4>
    <button v-on:click="openSignup">Open/Close Signup</button>
    <button v-on:click="openLogin">Open/Close Login</button>
    <div class="twoforms">
      <div class="forms">
        <!-- <button v-on:click="openSignup">Open/Close Signup</button> -->
        <div id="signupButton" style="display: none;"> 
          <user-form @create:user="createUser" />
        </div>
      </div>
      <div class="forms">
        <!-- <button v-on:click="openLogin">Open/Close Login</button> -->
        <div id="loginButton" style="display: none;"> 
          <login-form @login:user="loginUser" />
        </div>
      </div>
    </div>

    <!-- <user-list :users="users" /> -->  

  </div>

</template>

<script>
import UserForm from '@/components/UserForm.vue';
// import UserList from '@/components/UserList.vue';
import LoginForm from '@/components/LoginForm.vue';
import Navigation from '@/components/Navigation.vue'
const appData = {
  users: [],
  usermap: []
}
export default {
  name: 'app',
  components: {
    // UserList,
    Navigation,
    UserForm,
    LoginForm
  },
  data() {
    return appData;
  },
  mounted: function() {
    this.getUsers();
  },
  methods: {
    getUsers: getUsers,
    createUser: createUser,
    openSignup: openSignup,
    openLogin: openLogin,
    loginUser: loginUser
  }
}

async function getUsers() {
    try {
      const response = await fetch('api/users')
      const data = await response.json()
      appData.users = data.list
    } 
    catch (error) {
      console.error(error)
    }
}
async function createUser(user) {
  try {
    const response = await fetch('/api/users', {
      method: 'POST',
      body: JSON.stringify(user),
      headers: { "Content-type": "application/json; charset=UTF-8" }
    })
    console.log(response.status)
    // await getUsers()
  } 
  catch (error) {
    console.error(error)
  }
}

async function loginUser(user) {
  try {
    const response = await fetch('/api/users/login', {
      method: 'POST',
      body: JSON.stringify(user),
      headers: { "Content-type": "application/json; charset=UTF-8" }
    })
    console.log(response.status)
    // await getUsers()
  } 
  catch (error) {
    console.error(error)
  }
}

function openSignup() {
  var x = document.getElementById("signupButton");
  if (x.style.display === "none") {
    x.style.display = "block";
  } else {
    x.style.display = "none";
  }
}

function openLogin() {
  var x = document.getElementById("loginButton");
  if (x.style.display === "none") {
    x.style.display = "block";
  } else {
    x.style.display = "none";
  }
}

// async function loginUser() {
//     console.log("hello")
// }

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}
.twoforms {
  display: grid;
  grid-template-columns: 50% 50%; 
  grid-column-gap: 10px;
  margin: 5px;
}
.forms {
  padding: 10px;
}
button {
  background-color: lavender;
  border-color: lightgrey;
  border-radius: 40px;
  font-size: 15px;
  margin-top: 2px;
  height: 30px;
}

</style>
